import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { Badge } from "@/components/ui/Badge";
import { UpdateClass } from "./UpdateClass";
import { DeleteClass } from "./DeleteClass";
import { ChevronDown, ChevronUp } from "lucide-react";
import { UploadClassGallery } from "./UploadClassGallery";

export const ClassCard = ({ classes, sort, setSort }) => {
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
                onClick={() => setSort("title")}
                className="cursor-pointer"
              >
                Title {renderSortIcon("title")}
              </TableHead>
              <TableHead>Duration</TableHead>
              <TableHead>Status</TableHead>
              <TableHead
                onClick={() => setSort("created")}
                className="cursor-pointer"
              >
                Created At {renderSortIcon("created")}
              </TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {classes.map((item) => (
              <TableRow key={item.id}>
                <TableCell className="flex justify-center">
                  <img
                    src={item.image}
                    alt={item.title}
                    className="w-14 h-14 rounded-md object-cover border"
                  />
                </TableCell>
                <TableCell>{item.title}</TableCell>
                <TableCell>{item.duration} minutes</TableCell>
                <TableCell>
                  <Badge variant={item.isActive ? "default" : "secondary"}>
                    {item.isActive ? "Active" : "Inactive"}
                  </Badge>
                </TableCell>
                <TableCell>{item.createdAt}</TableCell>
                <TableCell className="space-x-2">
                  <UpdateClass classes={item} />
                  <DeleteClass classes={item} />
                  <UploadClassGallery classes={item} />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <div className="md:hidden w-full space-y-4 p-4">
        {classes.map((item) => (
          <div
            key={item.id}
            className="border border-border rounded-lg p-4 shadow-sm bg-background"
          >
            <div className="flex items-center gap-4 mb-3">
              <img
                src={item.image}
                alt={item.title}
                className="w-16 h-16 rounded-md object-cover border"
              />
              <div className="flex-1">
                <h3 className="text-base font-semibold">{item.title}</h3>
                <p className="text-sm text-muted-foreground">
                  {item.duration} minutes
                </p>
              </div>
            </div>

            <div className="text-sm text-muted-foreground space-y-1 mb-3">
              <p>
                <span className="font-medium text-foreground">Status:</span>{" "}
                <span
                  className={`ml-1 px-2 py-0.5 rounded-full text-xs font-medium ${
                    item.isActive
                      ? "bg-green-100 text-green-700"
                      : "bg-red-100 text-red-700"
                  }`}
                >
                  {item.isActive ? "Active" : "Inactive"}
                </span>
              </p>
              <p>
                <span className="font-medium text-foreground">Created:</span>{" "}
                {item.createdAt}
              </p>
            </div>

            <div className="flex justify-end gap-2">
              <UpdateClass classes={item} />
              <DeleteClass classes={item} />
              <UploadClassGallery classes={item} />
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
