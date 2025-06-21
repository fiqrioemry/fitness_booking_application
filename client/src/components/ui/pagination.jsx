import { ChevronLeft, ChevronRight } from "lucide-react";
// file naming typo
const Pagination = ({ page, limit, total, onPageChange }) => {
  const totalPages = Math.ceil(total / limit);
  const start = (page - 1) * limit + 1;
  const end = Math.min(start + limit - 1, total);

  return (
    <div className="flex items-center justify-between py-4 px-2 text-sm">
      <div className="text-muted-foreground">
        Showing <span className="font-medium text-primary">{start}</span>–
        <span className="font-medium text-primary">{end}</span> of{" "}
        <span className="font-medium text-primary">{total}</span>
      </div>
      <div className="flex items-center gap-2">
        <button
          onClick={() => onPageChange(page - 1)}
          disabled={page === 1}
          className="border rounded-md p-1 disabled:opacity-50"
        >
          <ChevronLeft className="w-4 h-4" />
        </button>
        <span className="font-medium text-sm w-8 text-center">{page}</span>
        <button
          onClick={() => onPageChange(page + 1)}
          disabled={page >= totalPages}
          className="border rounded-md p-1 disabled:opacity-50"
        >
          <ChevronRight className="w-4 h-4" />
        </button>
      </div>
    </div>
  );
};

export { Pagination };
