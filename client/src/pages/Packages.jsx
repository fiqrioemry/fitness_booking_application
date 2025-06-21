import {
  Card,
  CardTitle,
  CardFooter,
  CardHeader,
  CardContent,
  CardDescription,
} from "@/components/ui/Card";
import { CheckCircle2 } from "lucide-react";
import { packagesTitle } from "@/lib/constant";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { usePackagesQuery } from "@/hooks/usePackage";
import { useQueryStore } from "@/store/useQueryStore";
import { Pagination } from "@/components/ui/Pagination";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";
import { PackagesSkeleton } from "@/components/loading/PackagesSkeleton";

const Packages = () => {
  useDocumentTitle(packagesTitle);
  const navigate = useNavigate();
  const { page, limit, setPage } = useQueryStore();
  const { data, isLoading, isError, refetch } = usePackagesQuery({
    page,
    limit,
  });
  if (isLoading) return <PackagesSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const packages = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section py-24 text-foreground">
      <div className="bg-primary text-primary-foreground rounded-xl shadow px-6 py-10 text-center space-y-2 mb-10">
        <h3 className="text-3xl font-bold">Choose a Package</h3>
        <p className="text-sm opacity-80">
          Find the right plan that matches your fitness goals and schedule.
        </p>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {packages.map((pkg) => {
          const hasDiscount = pkg.discount > 0;
          const discountedPrice = pkg.price * (1 - pkg.discount / 100);

          return (
            <Card
              key={pkg.id}
              className={`relative overflow-hidden card-hover ${
                !pkg.isActive ? "opacity-60" : ""
              }`}
            >
              {/* SOLD OUT Banner */}
              {!pkg.isActive && (
                <div className="absolute top-4 -left-12 w-40 rotate-[-45deg] bg-red-600 text-white text-center text-xs font-bold py-1 shadow-lg z-10">
                  SOLD OUT
                </div>
              )}

              {/* Discount Badge */}
              {hasDiscount && (
                <span className="absolute top-4 right-4 bg-green-600 text-white text-xs px-2 py-1 rounded-full z-10 shadow">
                  {pkg.discount}% OFF
                </span>
              )}

              {/* Image */}
              <img
                src={pkg.image}
                alt={pkg.name}
                className="w-full h-44 object-cover"
              />

              <CardHeader className="text-center">
                <CardTitle className="line-clamp-1 text-lg">
                  {pkg.name}
                </CardTitle>
                <CardDescription className="line-clamp-2">
                  {pkg.description}
                </CardDescription>
              </CardHeader>

              <CardContent className="px-5 pb-0">
                <div className="text-sm text-muted-foreground mb-1">
                  Credit: <span className="font-medium">{pkg.credit}</span>
                </div>

                <div className="text-lg font-bold">
                  {hasDiscount ? (
                    <>
                      <span className="text-red-600">
                        Rp {discountedPrice.toLocaleString("id-ID")}
                      </span>
                      <span className="ml-2 text-sm line-through text-muted-foreground">
                        Rp {pkg.price.toLocaleString("id-ID")}
                      </span>
                    </>
                  ) : (
                    <span>Rp {pkg.price.toLocaleString("id-ID")}</span>
                  )}
                </div>

                <ul className="mt-4 text-sm text-muted-foreground space-y-2 h-20 overflow-y-auto">
                  {pkg.additional.slice(0, 4).map((item, idx) => (
                    <li key={idx} className="flex items-start gap-2">
                      <CheckCircle2 className="w-4 h-4 text-primary mt-1" />
                      <span>{item}</span>
                    </li>
                  ))}
                </ul>
              </CardContent>

              <CardFooter className="px-5 pb-5">
                <Button
                  onClick={() => navigate(`/packages/${pkg.id}`)}
                  className="w-full"
                  disabled={!pkg.isActive}
                >
                  {pkg.isActive ? "Buy Now" : "Unavailable"}
                </Button>
              </CardFooter>
            </Card>
          );
        })}
        <div>
          {pagination && pagination.totalRows > 10 && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </div>
      </div>
    </section>
  );
};

export default Packages;
