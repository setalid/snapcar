import CarPickupForm from "~/components/car-pickup-form";
import { Container } from "~/components/layout/container";

export default async function CarPickupPage() {
  const res = await fetch("http://localhost:8080/category/all");
  const data = await res.json();
  const { categories } = data

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

