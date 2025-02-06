export type Category = {
  id: string
  name: string
  priceFormula: string
}

// TODO: Remove?
export type Car = {
  id: string
  category: Category
  registrationNumber: string
}
