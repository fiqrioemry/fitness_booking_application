import { useRef } from "react";
import html2canvas from "html2canvas";
import jsPDF from "jspdf";
import { DollarSign } from "lucide-react";
import { useParams } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { Button } from "@/components/ui/Button";
import { formatRupiah, formatDate } from "@/lib/utils";
import { usePaymentDetailQuery } from "@/hooks/usePayment";

const TransactionDetail = () => {
  const invoiceRef = useRef();
  const { id } = useParams();
  const { data, isLoading } = usePaymentDetailQuery(id);

  const handleDownload = async () => {
    const element = invoiceRef.current;

    const canvas = await html2canvas(element, { scale: 2 });
    const imgData = canvas.toDataURL("image/png");

    const pdf = new jsPDF({
      orientation: "portrait",
      unit: "mm",
      format: "a4",
    });

    const pdfWidth = pdf.internal.pageSize.getWidth();
    const pdfHeight = (canvas.height * pdfWidth) / canvas.width;

    pdf.addImage(imgData, "PNG", 0, 0, pdfWidth, pdfHeight);
    pdf.save(`invoice-${id}.pdf`);
  };

  if (isLoading || !data) return <Loading />;

  return (
    <div className="w-full bg-background px-8 py-10 print:p-0 print:max-w-full">
      <div className="flex justify-end mb-4 print:hidden">
        <Button onClick={handleDownload}>Download Invoice PDF</Button>
      </div>

      <div
        ref={invoiceRef}
        className="relative bg-white text-background p-10 shadow-md print:shadow-none"
      >
        {data.status === "success" || data.status === "pending" ? (
          <div className="absolute flex items-center gap-4 border border-green-500 border-10 p-10 text-[72px] font-bold text-green-500 opacity-20 rotate-[-30deg] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 pointer-events-none select-none">
            <DollarSign className="h-20 w-20" /> <span>PAID</span>
          </div>
        ) : (
          <div className="absolute flex items-center gap-4 border border-red-500 border-10 p-10 text-[72px] font-bold text-red-500 opacity-20 rotate-[-30deg] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 pointer-events-none select-none">
            <DollarSign className="h-20 w-20" /> <span>UNPAID</span>
          </div>
        )}

        <h2 className="text-2xl font-bold text-center mb-6">INVOICE</h2>

        <div className="mb-6 text-sm space-y-1">
          <p>
            Invoice No.:{" "}
            <span className="font-medium">{data.invoiceNumber}</span>
          </p>
          <p>Transaction Date: {formatDate(data.paidAt)}</p>
          <p>Recipient: {data.fullname}</p>
        </div>

        <div className="border rounded mb-6">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b bg-muted">
                <th className="p-2 text-left">Package</th>
                <th className="p-2 text-right">Total</th>
              </tr>
            </thead>
            <tbody>
              <tr className="border-b">
                <td className="p-2">{data.packageName}</td>
                <td className="p-2 text-right">{formatRupiah(data.total)}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div className="text-sm mb-6">
          <div className="grid gap-2 text-right w-full md:w-[60%] ml-auto">
            <div className="flex justify-between items-center">
              <span className="text-muted-foreground">Tax</span>
              <span>{formatRupiah(data.tax)}</span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-muted-foreground">Discount</span>
              <span>- {formatRupiah(data.voucherDiscount)}</span>
            </div>
            <hr className="my-1 border-muted" />
            <div className="flex justify-between items-center font-bold text-lg">
              <span>Total Payment</span>
              <span>{formatRupiah(data.total)}</span>
            </div>
          </div>
        </div>

        <div className="text-sm">
          <p>Email: {data.email}</p>
          <p>Payment Method: {data.paymentMethod}</p>
        </div>
      </div>
    </div>
  );
};

export default TransactionDetail;
