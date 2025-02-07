'use client'

import { useForm } from "react-hook-form"
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod"
import { Button } from "./ui/button";
import { FormField, FormItem, FormLabel, FormControl, FormDescription, FormMessage, Form } from "./ui/form";
import { Input } from "./ui/input";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { Category } from "~/types/types";
import { useRouter } from "next/navigation";


const CarPickupSchema = z.object({
  bookingNumber: z.string().min(1, 'Booking number is required'),
  registrationNumber: z.string().min(1, 'Registration number is required'),
  customerSSN: z
    .string()
    .min(1, 'Social security number is required')
    .regex(/^\d{11}$/, 'Invalid SSN format (expected DDMMYYXXXXX)'),
  carCategory: z.string().min(1, 'Car category is required'),
  pickupDateTime: z.preprocess(input => `${input}:00Z`,
    z.string().datetime({ local: true })),
  currentMeterReading: z.number({ coerce: true }).min(0)
});

type CarPickupFormProps = {
  categories: Category[]
}

export default function CarPickupForm(p: CarPickupFormProps) {
  const router = useRouter();

  const form = useForm<z.infer<typeof CarPickupSchema>>({
    resolver: zodResolver(CarPickupSchema),
    defaultValues: {
      bookingNumber: "",
      registrationNumber: "",
      customerSSN: "",
      carCategory: "",
      pickupDateTime: "",
      currentMeterReading: 0,
    },
  })

  function onSubmit(values: z.infer<typeof CarPickupSchema>) {
    console.log(values)
    const body = {
      booking_number: values.bookingNumber,
      registration_number: values.registrationNumber,
      customer_ssn: values.customerSSN,
      car_category_name: values.carCategory,
      pickup_date_time: values.pickupDateTime,
      pickup_meter_reading: values.currentMeterReading,
    };

    fetch("http://localhost:8080/rental/pickup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        console.log("Success:", data);
        router.replace("/")
      })
      .catch((error) => {
        console.error("Error:", error);
      });
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
                <Input placeholder="DDMMYYXXXXX" {...field} />
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
                    <SelectItem key={c.name} value={c.name}>{c.name}</SelectItem>
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
              <FormLabel>Booking number</FormLabel>
              <FormControl>
                <Input type="datetime-local" {...field} />
              </FormControl>
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
