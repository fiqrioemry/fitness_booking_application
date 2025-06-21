import { Button } from "@/components/ui/Button";

export const NoPackage = () => {
  return (
    <div className="text-center py-16 px-6 bg-muted/50 border border-dashed border-border rounded-xl space-y-4">
      <div className="flex justify-center">
        <img
          src="/no-packages.webp"
          alt="No Packages"
          className="h-60 md:h-72 object-contain"
        />
      </div>

      <h2 className="text-lg font-semibold text-foreground">
        You don’t have any active packages yet.
      </h2>

      <p className="text-sm text-muted-foreground">
        Purchase a package now to start joining your favorite classes!
      </p>

      <Button asChild>
        <a href="/packages">View Packages</a>
      </Button>
    </div>
  );
};
