import { ArrowRight } from "lucide-react";
import Link from "next/link";
import { Container } from "~/components/layout/container";
import { Button } from "~/components/ui/button";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "~/components/ui/table";

export default function HomePage() {
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
            <Link href="/car-delivery">
              Car Delivery
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
                <TableHead>Meter reading (km)</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow>
                <TableCell>INV001</TableCell>
                <TableCell>Paid</TableCell>
                <TableCell>Credit Card</TableCell>
                <TableCell>$250.00</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </Container>
    </main>
  );
}
