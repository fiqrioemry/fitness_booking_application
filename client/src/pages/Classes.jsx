import {
  Card,
  CardHeader,
  CardTitle,
  CardFooter,
  CardDescription,
} from "@/components/ui/Card";
import { classesTitle } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useClassesQuery } from "@/hooks/useClass";
import { useQueryStore } from "@/store/useQueryStore";
import { Pagination } from "@/components/ui/Pagination";
import { useSearchParams, Link } from "react-router-dom";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";
import { ClassesSkeleton } from "@/components/loading/ClassesSkeleton";
import { SearchFilterSelection } from "@/components/input/SearchFilterSelection";

const Classes = () => {
  useDocumentTitle(classesTitle);

  const { q, page, limit, setPage } = useQueryStore();
  const [searchParams, setSearchParams] = useSearchParams();

  const filters = {
    typeId: searchParams.get("typeId") || "",
    levelId: searchParams.get("levelId") || "",
    categoryId: searchParams.get("categoryId") || "",
    locationId: searchParams.get("locationId") || "",
    subcategoryId: searchParams.get("subcategoryId") || "",
  };

  const {
    data: response,
    isLoading,
    isError,
    refetch,
  } = useClassesQuery({
    q,
    page,
    limit,
    status: "active",
    ...filters,
  });

  if (isLoading) return <ClassesSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const classes = response?.classes || [];

  const pagination = response?.pagination || [];

  return (
    <section className="section py-24 text-foreground">
      {/* Heading */}
      <div className="bg-primary text-primary-foreground rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <h3 className="text-3xl font-bold">Explore Fitness Classes</h3>
        <p className="text-sm opacity-80">
          Discover personalized sessions tailored for your needs, from beginner
          to advanced levels.
        </p>
      </div>

      {/* Filter Bar */}
      <div className="sticky top-4 z-10 bg-card text-foreground border border-border shadow-sm rounded-xl p-4 mb-8">
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
          <SearchFilterSelection
            paramKey="locationId"
            label="Location"
            data="location"
          />
          <SearchFilterSelection paramKey="typeId" label="Type" data="type" />
          <SearchFilterSelection
            paramKey="categoryId"
            label="Category"
            data="category"
          />
          <SearchFilterSelection
            paramKey="subcategoryId"
            label="Subcategory"
            data="subcategory"
          />
          <SearchFilterSelection
            paramKey="levelId"
            label="Level"
            data="level"
          />
        </div>
      </div>

      {/* Result */}
      {classes.length === 0 ? (
        <div className="flex flex-col items-center justify-center text-center py-4 col-span-full">
          <img
            src="/no-classes.webp"
            alt="no classes image found"
            className="w-72 mb-6 opacity-50"
          />
          <h3 className="text-xl font-semibold">No classes found</h3>
          <p className="text-muted-foreground mb-4 max-w-md text-sm">
            We couldn't find any classes based on your current filters. Try
            adjusting your selections or reset all filters.
          </p>
          <Button variant="default" onClick={() => setSearchParams({})}>
            Reset Filters
          </Button>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {classes.map((cls) => (
            <Link to={`/classes/${cls.id}`} key={cls.id}>
              <Card className="group h-full flex flex-col transition-transform hover:-translate-y-1 duration-300">
                <div className="relative h-48 w-full overflow-hidden">
                  <img
                    src={cls.image}
                    alt={cls.title}
                    className={`w-full h-full object-cover ${
                      !cls.isActive
                        ? "grayscale opacity-60"
                        : "group-hover:scale-105 transition-all"
                    }`}
                  />
                  {!cls.isActive && (
                    <div className="absolute inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-10">
                      <span className="bg-destructive text-destructive-foreground text-xs font-bold uppercase px-3 py-1 rounded-full">
                        Registration Closed
                      </span>
                    </div>
                  )}
                </div>

                <CardHeader className="p-4">
                  <CardTitle className="text-lg line-clamp-2">
                    {cls.title}
                  </CardTitle>
                  <CardDescription className="mb-1 line-clamp-2">
                    {cls.description}
                  </CardDescription>
                </CardHeader>

                <CardFooter className="text-xs text-muted-foreground pt-0 px-4 pb-4 mt-auto">
                  Duration: {cls.duration} mins
                </CardFooter>
              </Card>
            </Link>
          ))}
        </div>
      )}
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
    </section>
  );
};

export default Classes;
