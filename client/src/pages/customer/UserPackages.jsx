import { Pagination } from "@/components/ui/Pagination";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserPackagesQuery } from "@/hooks/useUserPackage";
import { SectionTitle } from "@/components/header/SectionTitle";
import { NoPackage } from "@/components/customer/packages/NoPackage";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { PackageCard } from "@/components/customer/packages/PackageCard";

const UserPackages = () => {
  const { data, isError, refetch, isLoading } = useUserPackagesQuery({
    page: 1,
    limit: 10,
    status: "created_at_asc",
  });

  if (isLoading) return <SectionSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const pkgs = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Packages"
        description="List of your purchased packages & remaining available sessions."
      />
      {pkgs.length === 0 ? <NoPackage /> : <PackageCard pkgs={pkgs} />}

      {pagination && pagination.totalRows > 10 && (
        <Pagination
          page={pagination.page}
          onPageChange={setPage}
          limit={pagination.limit}
          total={pagination.totalRows}
        />
      )}
    </section>
  );
};

export default UserPackages;
