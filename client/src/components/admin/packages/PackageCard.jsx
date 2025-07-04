import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { Badge } from "@/components/ui/Badge";
import { PackageUpdate } from "./PackageUpdate";
import { PackageDelete } from "./PackageDelete";
import { ChevronDown, ChevronUp } from "lucide-react";
import { formatRupiah, truncateText } from "@/lib/utils";

export const PackageCard = ({ packages, sort, setSort }) => {
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
              <TableHead>Thumbnail</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("name")}
              >
                Package Name
                {renderSortIcon("name")}
              </TableHead>
              <TableHead>Description</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("price")}
              >
                Price
                {renderSortIcon("price")}
              </TableHead>
              <TableHead>Credit</TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="h-12">
            {packages.map((pkg) => (
              <TableRow key={pkg.id}>
                <TableCell className="flex justify-center">
                  <img
                    src={pkg.image}
                    alt={pkg.name}
                    className="w-14 h-14 rounded-md object-cover border"
                  />
                </TableCell>
                <TableCell>{pkg.name}</TableCell>
                <TableCell title={pkg.description}>
                  {truncateText(pkg.description, 25)}
                </TableCell>
                <TableCell>{formatRupiah(pkg.price)}</TableCell>
                <TableCell>{pkg.credit} sessions</TableCell>
                <TableCell>
                  <Badge variant={pkg.isActive ? "default" : "secondary"}>
                    {pkg.isActive ? "Active" : "Inactive"}
                  </Badge>
                </TableCell>
                <TableCell className="space-x-2">
                  <PackageUpdate pkg={pkg} />
                  <PackageDelete pkg={pkg} />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <div className="md:hidden w-full space-y-4 p-4">
        {packages.map((pkg) => (
          <div
            key={pkg.id}
            className="border border-border rounded-lg p-4 shadow-sm bg-background"
          >
            <div className="flex items-center gap-4 mb-3">
              <img
                src={pkg.image}
                alt={pkg.name}
                className="w-16 h-16 rounded-md object-cover border"
              />
              <div className="flex-1">
                <h3 className="text-base text-start font-semibold text-foreground">
                  {pkg.name}
                </h3>
                <p className="text-xs text-start text-muted-foreground">
                  {pkg.description}
                </p>
              </div>
            </div>

            <div className="text-sm text-start text-muted-foreground space-y-1 mb-3">
              <p>
                <span className="font-medium text-foreground">Price :</span>{" "}
                {formatRupiah(pkg.price)}
              </p>
              <p>
                <span className="font-medium text-foreground">Credit :</span>{" "}
                {pkg.credit} sessions
              </p>
              <p>
                <span className="font-medium text-foreground">Status :</span>{" "}
                <span
                  className={`ml-1 px-2 py-0.5 rounded-full text-xs font-medium ${
                    pkg.isActive
                      ? "bg-green-100 text-green-700"
                      : "bg-red-100 text-red-700"
                  }`}
                >
                  {pkg.isActive ? "Active" : "Inactive"}
                </span>
              </p>
            </div>

            <div className="flex justify-end gap-2">
              <PackageUpdate pkg={pkg} />
              <PackageDelete pkg={pkg} />
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
