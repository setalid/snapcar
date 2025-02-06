import { Plus } from "lucide-react";
import { Container } from "~/components/layout/container";
import { Button } from "~/components/ui/button";
import { Separator } from "~/components/ui/separator";
import { Car } from "~/types/types";

const cars: Car[] = [
  {
    id: "car-1",
    category: {
      id: "",
      name: "Small car",
      priceFormula: ""
    },
    registrationNumber: "AB1234"
  },
  {
    id: "car-2",
    category: {
      id: "",
      name: "Combi",
      priceFormula: ""
    },
    registrationNumber: "BA1234"
  },
]

export default function CategoriesPage() {
  return (
    <main>
      <Container>
        <div className="flex justify-end">
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            Create Car (TBD)
          </Button>
        </div>
        <Separator className="mt-6" />
        <h1 className="text-xl mt-6">Cars</h1>
      </Container>
    </main>
  )
}
