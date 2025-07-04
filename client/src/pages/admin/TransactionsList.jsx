import { Loading } from "@/components/ui/Loading";
import { useDebounce } from "@/hooks/useDebounce";
import { paymentStatusOptions } from "@/lib/constant";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { useAllPaymentsQuery } from "@/hooks/usePayment";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SearchNotFound } from "@/components/ui/SearchNotFound";
import { FilterSelection } from "@/components/ui/FilterSelection";
import { useTransactionStore } from "@/store/useTransactionStore";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const TransactionsList = () => {
  const { q, setQ, page, limit, sort, setPage, status, setSort, setStatus } =
    useTransactionStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useAllPaymentsQuery({
    q: debouncedQ,
    status,
    sort,
    page,
    limit,
  });

  const transactions = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section">
      <SectionTitle
        title="Transaction List"
        description="Manage all user transactions and monitor payment activities."
      />

      <div className="flex items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by name or email"}
        />

        <FilterSelection
          value={status}
          onChange={setStatus}
          options={paymentStatusOptions}
        />
      </div>

      <Card>
        <CardContent className="p-0">
          {isLoading ? (
            <Loading />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : transactions.length === 0 ? (
            <SearchNotFound title="No transactions found" q={q} />
          ) : (
            <TransactionCard
              sort={sort}
              setSort={setSort}
              transactions={transactions}
            />
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

export default TransactionsList;
