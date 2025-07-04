/* eslint-disable react/prop-types */
"use client";

import * as React from "react";
import * as TabsPrimitive from "@radix-ui/react-tabs";
import { cva } from "class-variance-authority";

import { cn } from "@/lib/utils";

// Root Tabs tetap sama
const Tabs = TabsPrimitive.Root;

// Variants dan Sizes untuk TabsList
const tabsListVariants = cva(
  "inline-flex items-center text-muted-foreground transition-all",
  {
    variants: {
      variant: {
        default: "bg-muted px-2 py-10",
        underline: "w-full border-b border-b-2",
      },
      size: {
        sm: "rounded-none h-10",
        md: "rounded-md h-10",
        lg: "rounded-lg h-12",
      },
    },
    defaultVariants: {
      variant: "underline",
      size: "sm",
    },
  }
);

const TabsList = React.forwardRef(
  ({ className, variant, size, ...props }, ref) => (
    <TabsPrimitive.List
      ref={ref}
      className={cn(tabsListVariants({ variant, size, className }))}
      {...props}
    />
  )
);
TabsList.displayName = TabsPrimitive.List.displayName;

// Variants dan Sizes untuk TabsTrigger
const tabsTriggerVariants = cva(
  "inline-flex items-center justify-center whitespace-nowrap text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
  {
    variants: {
      variant: {
        default:
          "data-[state=active]:bg-background data-[state=active]:text-foreground data-[state=active]:shadow",
        underline:
          "data-[state=active]:border-b-2 data-[state=active]:border-primary",
        none: "data-[state=active]:text-primary data-[state=active]:border-primary",
        ghost:
          "hover:bg-accent hover:text-accent-foreground border-b-2 data-[state=active]:border-primary data-[state=active]:text-primary",
      },
      size: {
        sm: "px-9 py-1 h-10 text-sm rounded-none",
        md: "px-9 py-1 h-10 text-sm rounded-md",
        lg: "px-12 py-1 h-12 text-base rounded-lg",
      },
    },
    defaultVariants: {
      variant: "underline",
      size: "sm",
    },
  }
);

const TabsTrigger = React.forwardRef(
  ({ className, variant, size, ...props }, ref) => (
    <TabsPrimitive.Trigger
      ref={ref}
      className={cn(tabsTriggerVariants({ variant, size, className }))}
      {...props}
    />
  )
);
TabsTrigger.displayName = TabsPrimitive.Trigger.displayName;

// TabsContent tetap sama
const TabsContent = React.forwardRef(({ className, ...props }, ref) => (
  <TabsPrimitive.Content
    ref={ref}
    className={cn(
      "mt-2 ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
      className
    )}
    {...props}
  />
));
TabsContent.displayName = TabsPrimitive.Content.displayName;

export { Tabs, TabsList, TabsTrigger, TabsContent };
