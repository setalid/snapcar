import { Container } from "~/components/layout/container";
import { Plus } from "lucide-react";
import { Button } from "~/components/ui/button";
import { Separator } from "~/components/ui/separator";
import { Category } from "~/types/types";
import { categories } from "~/mock/categories";

export default function CategoriesPage() {
  return (
    <main>
      <Container>
        <div className="flex justify-end">
          <Button>
            <Plus className="mr-2 h-4 w-4" />
            Create Category
          </Button>
        </div>
        <Separator className="mt-6" />
        <h1 className="text-xl mt-6">Categories</h1>
        <div className="space-y-4 mt-4">
          {categories.map((category: Category) => (
            <div key={category.id} className="p-4 border rounded-lg">
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
