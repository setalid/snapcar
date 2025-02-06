"use client";

import Link from "next/link";
import { Container } from "~/components/layout/container";
import { Home, Car, Boxes } from "lucide-react";
import { NavigationMenu, NavigationMenuItem, NavigationMenuLink, NavigationMenuList, navigationMenuTriggerStyle } from "./ui/navigation-menu";
import { usePathname } from "next/navigation";

export function Header() {
  const pathname = usePathname();
  return (
    <div className="border-b">
      <Container>
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              <Link href="/" legacyBehavior passHref>
                <NavigationMenuLink active={pathname == "/"} className={navigationMenuTriggerStyle()}>
                  <Home className="h-4 w-4" />
                  <div className="ml-1">Rentals</div>
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link href="/categories" legacyBehavior passHref>
                <NavigationMenuLink active={pathname == "/categories"} className={navigationMenuTriggerStyle()}>
                  <Boxes className="h-4 w-4" />
                  <div className="ml-1">Categories</div>
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
            <NavigationMenuItem>
              <Link href="/cars" legacyBehavior passHref>
                <NavigationMenuLink active={pathname == "/cars"} className={navigationMenuTriggerStyle()}>
                  <Car className="h-4 w-4" />
                  <div className="ml-1">Cars</div>
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
      </Container>
    </div>
  );
}
