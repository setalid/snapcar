import CarReturnForm from "~/components/car-return-form";
import { Container } from "~/components/layout/container";

export default function CarPickupPage() {
  return (
    <main>
      <Container>
        <h1 className="text-xl mt-6">Car return</h1>
        <div className="mt-4">
          <CarReturnForm />
        </div>
      </Container>
    </main>
  )
}
