import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { Printer } from "lucide-react";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { useNavigate } from "react-router-dom";
import { formatDate, formatRupiah } from "@/lib/utils";
import { ChevronDown, ChevronUp } from "lucide-react";

export const TransactionCard = ({ transactions, sort, setSort }) => {
  const navigate = useNavigate();

  const handlePrint = (id) => {
    navigate(`/admin/transactions/${id}`);
  };

  const renderSortIcon = (field) => {
    if (sort === `${field}_asc`)
      return <ChevronUp className="w-4 h-4 inline" />;
    if (sort === `${field}_desc`)
      return <ChevronDown className="w-4 h-4 inline" />;
    return null;
  };

  return (
    <>
      <div className="hidden md:block w-full">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("name")}
              >
                Fullname {renderSortIcon("name")}
              </TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("email")}
              >
                Email
                {renderSortIcon("email")}
              </TableHead>
              <TableHead>Invoice Number</TableHead>
              <TableHead>Method</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Total</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("paid_at")}
              >
                Paid At
                {renderSortIcon("paid_at")}
              </TableHead>
              <TableHead>Print</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="h-12">
            {transactions.map((tx) => (
              <TableRow key={tx.id}>
                <TableCell>
                  <div>
                    {tx.fullname.length > 20
                      ? tx.fullname.slice(0, 20) + "..."
                      : tx.fullname}
                  </div>
                </TableCell>

                <TableCell>{tx.email || ""}</TableCell>

                <TableCell>{tx.invoiceNumber || ""}</TableCell>

                <TableCell>{tx.paymentMethod || ""}</TableCell>

                <TableCell>
                  {tx.status === "success" ? (
                    <Badge> success</Badge>
                  ) : "failed" ? (
                    <Badge variant="destructive">failed</Badge>
                  ) : (
                    <Badge variant="secondary">failed</Badge>
                  )}
                </TableCell>

                <TableCell>{formatRupiah(tx.total)}</TableCell>
                <TableCell>{tx.paidAt ? formatDate(tx.paidAt) : "-"}</TableCell>
                <TableCell>
                  <Button
                    size="icon"
                    variant="ghost"
                    title="Print Invoice"
                    onClick={() => handlePrint(tx.id)}
                  >
                    <Printer className="w-5 h-5" />
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      {/* Mobile view */}
      <div className="md:hidden space-y-4 p-4  w-full">
        {transactions.map((tx) => (
          <div
            key={tx.id}
            className="border rounded-lg p-4 shadow-sm space-y-2"
          >
            <div>
              <h3 className="text-base font-semibold">{tx.fullname}</h3>
              <p className="text-sm text-muted-foreground">{tx.userEmail}</p>
            </div>
            <div className="text-sm space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <p>
                  <strong>{tx.invoiceNumber}</strong>
                </p>
                <p>{formatRupiah(tx.total)}</p>
                <p>{tx.method?.toUpperCase() || "-"}</p>

                <Badge
                  variant={
                    tx.status === "success"
                      ? "default"
                      : tx.status === "failed"
                      ? "destructive"
                      : "secondary"
                  }
                >
                  {tx.status}
                </Badge>
              </div>
              <div className="space-x-4">
                <span className="text-muted-foreground">Paid At</span>
                <span className="text-right whitespace-nowrap">
                  {tx.paidAt ? formatDate(tx.paidAt) : "-"}
                </span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
