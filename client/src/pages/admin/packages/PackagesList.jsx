import { Link } from "react-router-dom";
import { CirclePlus } from "lucide-react";
import { statusOptions } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useDebounce } from "@/hooks/useDebounce";
import { Loading } from "@/components/ui/Loading";
import { usePackagesQuery } from "@/hooks/usePackage";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { usePackageStore } from "@/store/usePackageStore";
import { SearchNotFound } from "@/components/ui/SearchNotFound";
import { SectionTitle } from "@/components/header/SectionTitle";
import { FilterSelection } from "@/components/ui/FilterSelection";
import { PackageCard } from "@/components/admin/packages/PackageCard";

const PackagesList = () => {
  const { q, setQ, page, limit, sort, status, setPage, setSort, setStatus } =
    usePackageStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = usePackagesQuery({
    q: debouncedQ,
    status,
    sort,
    page,
    limit,
  });

  const packages = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="Package Management"
        description="View, add, and manage training packages available for purchase by
          users."
      />

      <div className="flex flex-col md:flex-row justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by name or descriptions"}
        />

        <div className="flex justify-end items-center  gap-4">
          <FilterSelection
            value={status}
            className="w-32 h-12"
            onChange={setStatus}
            options={statusOptions}
          />
          <Link to="/admin/packages/add">
            <Button>
              <CirclePlus className="w-4 h-4 mr-2" />
              Add Package
            </Button>
          </Link>
        </div>
      </div>

      <Card>
        <CardContent className="p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : packages.length === 0 ? (
            <SearchNotFound title="No Packages found" q={q} />
          ) : (
            <PackageCard packages={packages} sort={sort} setSort={setSort} />
          )}

          {pagination && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default PackagesList;
