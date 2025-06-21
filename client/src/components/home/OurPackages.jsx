import { Link } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { usePackagesQuery } from "@/hooks/usePackage";
import { OurPackageSkeleton } from "@/components/loading/OurPackageSkeleton";

export const OurPackages = () => {
  const { data, isLoading } = usePackagesQuery({ limit: 3 });

  const packages = data?.data || [];

  return (
    <section className="py-20 px-4 max-w-7xl mx-auto">
      <h2 className="text-4xl font-bold text-center mb-4 font-heading">
        Our Packages
      </h2>
      <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
        From relaxing yoga sessions to high-intensity workouts, choose the class
        that suits your fitness goals.
      </p>
      {isLoading ? (
        <OurPackageSkeleton />
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto">
          {packages.map((pkg) => (
            <div
              key={pkg.id}
              className="bg-card text-foreground border border-border rounded-xl shadow hover:shadow-xl transition overflow-hidden"
            >
              <img
                src={pkg.image}
                alt={pkg.name}
                className="w-full h-56 object-cover"
              />
              <div className="p-6 space-y-3">
                <h4 className="text-lg font-bold">{pkg.name}</h4>
                <p className="text-sm text-muted-foreground line-clamp-2">
                  {pkg.description}
                </p>
                <div className="flex justify-end     items-center pt-2">
                  <Link to="/packages">
                    <Button size="sm">Get it now</Button>
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </section>
  );
};
