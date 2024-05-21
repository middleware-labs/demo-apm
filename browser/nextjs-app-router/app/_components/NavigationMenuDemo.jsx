"use client";

import * as React from "react";
import Link from "next/link";

import { cn } from "@/lib/utils";
import { Icons } from "@/components/ui/icons";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";

const components = [
  {
    title: "OpenTelemetry",
    href: "https://opentelemetry.io/",
    description:
      "OpenTelemetry is a set of APIs, libraries, agents, and instrumentation to provide observability for cloud-native software.",
  },
  {
    title: "Tailwind CSS",
    href: "https://tailwindcss.com/",
    description:
      "Tailwind CSS is a utility-first CSS framework for rapidly building custom designs.",
  },
  {
    title: "Next.js",
    href: "https://nextjs.org/",
    description:
      "Next.js is a React framework that enables functionality such as server-side rendering and generating static websites.",
  },
  {
    title: "ESLint",
    href: "https://eslint.org/",
    description:
      "ESLint is a static code analysis tool for identifying problematic patterns found in JavaScript code.",
  },
  {
    title: "Kubernetes",
    href: "https://kubernetes.io/",
    description:
      "Kubernetes is an open-source container-orchestration system for automating computer application deployment, scaling, and management.",
  },
  {
    title: "Middleware",
    href: "https://middleware.io/",
    description:
      "Middleware is a real-time observability platform that provides insights into the performance and behavior of your applications.",
  },
];

export function NavigationMenuDemo() {
  return (
    <NavigationMenu>
      <NavigationMenuList>
        <NavigationMenuItem>
          <Link href="/" legacyBehavior passHref>
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              Home
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link href="/form" legacyBehavior passHref>
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              Form Example
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <NavigationMenuTrigger>Dropdown Example</NavigationMenuTrigger>
          <NavigationMenuContent>
            <ul className="grid w-[400px] gap-3 p-4 md:w-[500px] md:grid-cols-2 lg:w-[600px] ">
              {components.map((component) => (
                <ListItem
                  key={component.title}
                  title={component.title}
                  href={component.href}
                >
                  {component.description}
                </ListItem>
              ))}
            </ul>
          </NavigationMenuContent>
        </NavigationMenuItem>
        <NavigationMenuItem>
          <Link
            href="https://docs.middleware.io/rum/rum-nextjs"
            legacyBehavior
            passHref
          >
            <NavigationMenuLink className={navigationMenuTriggerStyle()}>
              Documentation
            </NavigationMenuLink>
          </Link>
        </NavigationMenuItem>
      </NavigationMenuList>
    </NavigationMenu>
  );
}

const ListItem = React.forwardRef(
  ({ className, title, children, ...props }, ref) => {
    return (
      <li>
        <NavigationMenuLink asChild>
          <a
            ref={ref}
            className={cn(
              "block select-none space-y-1 rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground",
              className,
            )}
            {...props}
          >
            <div className="text-sm font-medium leading-none">{title}</div>
            <p className="line-clamp-2 text-sm leading-snug text-muted-foreground">
              {children}
            </p>
          </a>
        </NavigationMenuLink>
      </li>
    );
  },
);
ListItem.displayName = "ListItem";
