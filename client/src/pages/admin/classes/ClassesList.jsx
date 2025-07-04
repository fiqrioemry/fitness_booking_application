import { CirclePlus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { statusOptions } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { useClassesQuery } from "@/hooks/useClass";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SearchNotFound } from "@/components/ui/SearchNotFound";
import { ClassCard } from "@/components/admin/classes/ClassCard";
import { FilterSelection } from "@/components/ui/FilterSelection";
import { useClassStore } from "@/store/useClassStore";

const ClassesList = () => {
  const navigate = useNavigate();
  const { q, setQ, page, limit, sort, status, setPage, setSort, setStatus } =
    useClassStore();

  const debouncedQ = useDebounce(q, 500);

  const { data, isLoading, isError, refetch } = useClassesQuery({
    q: debouncedQ,
    sort,
    status,
    page,
    limit,
  });

  const classes = data?.classes || [];

  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="Class Management"
        description="View, add, and manage training classes available for users."
      />

      <div className="flex flex-col md:flex-row justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder="search by title"
        />

        <div className="flex justify-end items-center gap-4">
          <FilterSelection
            value={status}
            onChange={setStatus}
            className="w-32 h-12"
            options={statusOptions}
          />
          <Button size="nav" onClick={() => navigate("/admin/classes/add")}>
            <CirclePlus className="w-4 h-4 mr-2" />
            Add Class
          </Button>
        </div>
      </div>

      <Card className="shadow-sm">
        <CardContent className="p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : classes.length === 0 ? (
            <SearchNotFound title="No Classes found" q={q} />
          ) : (
            <ClassCard classes={classes} sort={sort} setSort={setSort} />
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

export default ClassesList;
