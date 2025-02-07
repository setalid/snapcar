import { Container } from "~/components/layout/container";
import { Plus } from "lucide-react";
import { Button } from "~/components/ui/button";
import { Separator } from "~/components/ui/separator";
import { Category } from "~/types/types";

export default async function CategoriesPage() {
  const res = await fetch("http://localhost:8080/category/all");
  const data = await res.json();
  const { categories } = data

  return (
    <main>
      <Container>
        <h1 className="text-xl mt-6">Categories</h1>
        <div className="space-y-4 mt-4">
          {categories.map((category: Category) => (
            <div key={category.name} className="p-4 border rounded-lg">
              <h3 className="text-lg font-semibold">{category.name}</h3>
              <p className="text-sm text-muted-foreground">
                {category.priceFormula}
              </p>
            </div>
          ))}
        </div>
      </Container>
    </main>
  )
}
