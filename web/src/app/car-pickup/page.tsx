import CarPickupForm from "~/components/car-pickup-form";
import { Container } from "~/components/layout/container";
import { categories } from "~/mock/categories";

export default function CarPickupPage() {
  return (
    <main>
      <Container>
        <h1 className="text-xl mt-6">Car pickup</h1>
        <div className="mt-4">
          <CarPickupForm categories={categories} />
        </div>
      </Container>
    </main>
  )
}
