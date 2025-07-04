import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  CartesianGrid,
} from "recharts";

const RevenueChart = ({ data, range }) => {
  const formattedData = data?.map((item) => {
    const date = new Date(item.date);
    let label = item.date;

    if (range === "daily") {
      label = date.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "short",
      });
    } else if (range === "monthly") {
      label = date.toLocaleDateString("id-ID", {
        month: "short",
        year: "numeric",
      });
    } else if (range === "yearly") {
      label = date.getFullYear().toString();
    }

    return { ...item, label };
  });

  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={formattedData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="label" />
        <YAxis />
        <Tooltip
          formatter={(value) => `Rp${value.toLocaleString("id-ID")}`}
          labelFormatter={(label) => `Tanggal: ${label}`}
        />
        <Line
          type="monotone"
          dataKey="total"
          stroke="#2563eb"
          strokeWidth={2}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export { RevenueChart };
