'use client'

import { useForm } from "react-hook-form"
import { z } from "zod";
import { format } from "date-fns"
import { zodResolver } from "@hookform/resolvers/zod"
import { Button } from "./ui/button";
import { FormField, FormItem, FormLabel, FormControl, FormDescription, FormMessage, Form } from "./ui/form";
import { Input } from "./ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { Category } from "~/types/types";
import { Popover, PopoverContent, PopoverTrigger } from "./ui/popover";
import { cn } from "~/lib/utils";
import { Calendar } from "./ui/calendar";
import { CalendarIcon } from "lucide-react";


const CarPickupSchema = z.object({
  bookingNumber: z.string().min(1, 'Booking number is required'),
  registrationNumber: z.string().min(1, 'Registration number is required'),
  customerSSN: z
    .string()
    .min(1, 'Social security number is required')
    .regex(/^\d{11}$/, 'Invalid SSN format (expected DDMMYYXXXXX)'),
  carCategory: z.string().min(1, 'Car category is required'),
  pickupDateTime: z.date(),
  currentMeterReading: z.number().min(0, 'Meter reading must be a positive number'), // Current meter reading (km)
});

type CarPickupFormProps = {
  categories: Category[]
}

export default function CarPickupForm(p: CarPickupFormProps) {
  const form = useForm<z.infer<typeof CarPickupSchema>>({
    resolver: zodResolver(CarPickupSchema),
    defaultValues: {},
  })

  function onSubmit(values: z.infer<typeof CarPickupSchema>) {
    console.log(values)
  }

  return (
    <Form  {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="bookingNumber"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Booking number</FormLabel>
              <FormControl>
                <Input placeholder="booking-1" {...field} />
              </FormControl>
              <FormDescription>
                Create or generate a booking number for the rental agreement
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="registrationNumber"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Registration number</FormLabel>
              <FormControl>
                <Input placeholder="AB12345" {...field} />
              </FormControl>
              <FormDescription>
                The registration number of the car
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="customerSSN"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Customer social security number (SSN)</FormLabel>
              <FormControl>
                <Input type="number" placeholder="DDMMYYXXXXX" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="carCategory"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Car category</FormLabel>
              <Select onValueChange={field.onChange} defaultValue={field.value}>
                <FormControl>
                  <SelectTrigger>
                    <SelectValue placeholder="Select a car category" />
                  </SelectTrigger>
                </FormControl>
                <SelectContent>
                  {p.categories.map((c: Category) => (
                    <SelectItem key={c.id} value={c.id}>{c.name}</SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <FormDescription>
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="pickupDateTime"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Current meter reading</FormLabel>
              <Popover>
                <PopoverTrigger asChild>
                  <FormControl>
                    <Button
                      variant={"outline"}
                      className={cn(
                        "w-[240px] pl-3 text-left font-normal",
                        !field.value && "text-muted-foreground"
                      )}
                    >
                      {field.value ? (
                        format(field.value, "PPP")
                      ) : (
                        <span>Pick a date</span>
                      )}
                      <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                    </Button>
                  </FormControl>
                </PopoverTrigger>
                <PopoverContent className="w-auto p-0" align="start">
                  <Calendar
                    mode="single"
                    selected={field.value}
                    onSelect={field.onChange}
                    disabled={(date) =>
                      date > new Date() || date < new Date("1900-01-01")
                    }
                    initialFocus
                  />
                </PopoverContent>
              </Popover>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="currentMeterReading"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Current meter reading</FormLabel>
              <FormControl>
                <Input type="number" placeholder="85000" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  )
}
