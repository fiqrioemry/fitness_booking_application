import { useState } from "react";
import { formatRupiah } from "@/lib/utils";
import { revenueRangeOptions } from "@/lib/constant";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAllPaymentsQuery } from "@/hooks/usePayment";
import { FilterSelection } from "@/components/ui/FilterSelection";
import { SummaryCard } from "@/components/admin/dashboard/SummaryCard";
import { RevenueChart } from "@/components/admin/dashboard/RevenueChart";
import { DashboardSkeleton } from "@/components/loading/DashboardSkeleton";
import { useDashboardSummary, useRevenueStats } from "@/hooks/useDashboard";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const Dashboard = () => {
  const [range, setRange] = useState("daily");
  const { data: revenue } = useRevenueStats({ range });
  const { data: response } = useAllPaymentsQuery({ limit: 5 });

  const { data: summary, isLoading, isError, refetch } = useDashboardSummary();

  if (isLoading) return <DashboardSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const transactions = response?.data || [];

  return (
    <div className="p-6 space-y-6">
      <div>
        <h3 className="mb-4">Statistic Summary</h3>
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <SummaryCard title="Total Users" value={summary?.totalUsers} />
          <SummaryCard title="Active Classes" value={summary?.totalClasses} />
          <SummaryCard title="Total Bookings" value={summary?.totalBookings} />
          <SummaryCard
            title="Total Revenue"
            value={formatRupiah(summary?.totalRevenue)}
          />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Transaction Volume</h3>
        <div className="bg-background rounded-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <h2>Revenue</h2>

            <FilterSelection
              value={range}
              onChange={setRange}
              options={revenueRangeOptions}
            />
          </div>

          <RevenueChart data={revenue?.revenueSeries} range={revenue?.range} />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Recent transaction</h3>
        <Card className="border">
          <CardContent className="p-0">
            <TransactionCard transactions={transactions} />
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Dashboard;
