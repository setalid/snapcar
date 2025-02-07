'use client'

import { useForm } from "react-hook-form"
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod"
import { Button } from "./ui/button";
import { FormField, FormItem, FormLabel, FormControl, FormDescription, FormMessage, Form } from "./ui/form";
import { Input } from "./ui/input";
import { useRouter } from "next/navigation";


const CarReturnSchema = z.object({
  bookingNumber: z.string().min(1, 'Booking number is required'),
  returnDateTime: z.preprocess(input => `${input}:00Z`,
    z.string().datetime({ local: true })),
  returnMeterReading: z.number({ coerce: true }).min(0)
});

export default function CarReturnForm() {
  const router = useRouter();

  const form = useForm<z.infer<typeof CarReturnSchema>>({
    resolver: zodResolver(CarReturnSchema),
    defaultValues: {
      bookingNumber: "",
      returnDateTime: "",
      returnMeterReading: 0,
    },
  })

  function onSubmit(values: z.infer<typeof CarReturnSchema>) {
    console.log(values)
    const body = {
      return_date_time: values.returnDateTime,
      return_meter_reading: values.returnMeterReading,
    };

    fetch("http://localhost:8080/rental/return/" + values.bookingNumber, {
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
          name="returnDateTime"
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
          name="returnMeterReading"
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
