import { ArrowRight } from "lucide-react";
import Link from "next/link";
import { Container } from "~/components/layout/container";
import { Button } from "~/components/ui/button";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";
import { Rental } from "~/types/types";

type Data = {
  rentals: Rental[]
}

export default async function HomePage() {
  const res = await fetch("http://localhost:8080/rental/all");
  const data = await res.json();
  const { rentals } = data as Data

  return (
    <main>
      <Container>
        <div className="flex gap-4">
          <Button asChild className="flex-1 h-24 text-lg">
            <Link href="/car-pickup">
              Car Pickup
              <ArrowRight />
            </Link>
          </Button>
          <Button asChild className="flex-1 h-24 text-lg">
            <Link href="/car-return">
              Car Return
              <ArrowRight />
            </Link>
          </Button>
        </div>
        <div className="mt-8">
          <h1 className="text-xl">Car rentals</h1>
          <Table className="mt-4">
            <TableCaption>A list of car rentals</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead>Booking #</TableHead>
                <TableHead>Registration #</TableHead>
                <TableHead>Car category</TableHead>
                <TableHead>Pickup date</TableHead>
                <TableHead>Pickup reading (km)</TableHead>
                <TableHead>Return date</TableHead>
                <TableHead>Return reading (km)</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {rentals?.map((rental: Rental) => (
                <TableRow key={rental.bookingNumber}>
                  <TableCell>{rental.bookingNumber}</TableCell>
                  <TableCell>{rental.registrationNumber}</TableCell>
                  <TableCell>{rental.carCategoryName}</TableCell>
                  <TableCell>{rental.pickupDate}</TableCell>
                  <TableCell>{rental.pickupMeterReading}</TableCell>
                  <TableCell>{rental.returnDate}</TableCell>
                  <TableCell>{rental.returnMeterReading}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </Container>
    </main>
  );
}
