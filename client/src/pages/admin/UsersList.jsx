import { roleOptions } from "@/lib/constant";
import { useUsersQuery } from "@/hooks/useUsers";
import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { useQueryStore } from "@/store/useQueryStore";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { UsersCard } from "@/components/admin/users/UsersCard";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SearchNotFound } from "@/components/ui/SearchNotFound";
import { FilterSelection } from "@/components/ui/FilterSelection";

const UsersList = () => {
  const { q, setQ, page, limit, sort, role, setPage, setSort, setRole } =
    useQueryStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useUsersQuery({
    q: debouncedQ,
    sort,
    role,
    page,
    limit,
  });

  const users = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="User List"
        description="Manage all registered users efficiently."
      />

      <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-4 mt-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by name or email"}
        />
        <FilterSelection
          value={role}
          onChange={setRole}
          options={roleOptions}
        />
      </div>

      <Card>
        <CardContent className="p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : users?.length === 0 ? (
            <SearchNotFound title="User not found" q={q} />
          ) : (
            <UsersCard users={users} sort={sort} setSort={setSort} />
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

export default UsersList;
