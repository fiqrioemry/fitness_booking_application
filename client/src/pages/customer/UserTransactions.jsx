import { Loading } from "@/components/ui/Loading";
import { paymentStatusOptions } from "@/lib/constant";
import { Pagination } from "@/components/ui/Pagination";
import { useMyPaymentsQuery } from "@/hooks/usePayment";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { FilterSelection } from "@/components/ui/FilterSelection";
import { useTransactionStore } from "@/store/useTransactionStore";
import { MyTransactionCard } from "@/components/customer/transactions/MyTransactionCard";
import { NoTransactionRecord } from "@/components/customer/transactions/NoTransactionRecord";

const UserTransactions = () => {
  const { page, limit, sort, setPage, status, setSort, setStatus } =
    useTransactionStore();

  const { data, isLoading, isError, refetch } = useMyPaymentsQuery({
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
            <NoTransactionRecord />
          ) : (
            <MyTransactionCard
              sort={sort}
              setSort={setSort}
              transactions={transactions}
            />
          )}
          {pagination && pagination.totalRows > 10 && (
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

export default UserTransactions;
