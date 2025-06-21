import {
  Table,
  TableCell,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { formatDate } from "@/lib/utils";
import { Badge } from "@/components/ui/Badge";
import { ChevronDown, ChevronUp, Eye } from "lucide-react";
import { useLocation, useNavigate } from "react-router-dom";

export const UsersCard = ({ users, sort, setSort }) => {
  const navigate = useNavigate();
  const location = useLocation();

  const openModal = (id) => {
    navigate(`/admin/users/${id}`, {
      state: { backgroundLocation: location },
    });
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
      <div className="hidden md:block max-w-8xl w-full">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Avatar</TableHead>
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
              <TableHead>Role</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("joined")}
              >
                Joined Since {renderSortIcon("joined")}
              </TableHead>
              <TableHead>Details</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="h-12">
            {users.map((user) => (
              <TableRow
                key={user.id}
                className="border-t border-border hover:bg-muted transition"
              >
                <TableCell className="flex justify-center">
                  <img
                    src={user.avatar}
                    alt={user.fullname}
                    className="w-10 h-10 rounded-full object-cover border"
                  />
                </TableCell>
                <TableCell className="font-medium">{user.fullname}</TableCell>
                <TableCell>{user.email}</TableCell>
                <TableCell>
                  <Badge
                    variant={
                      user.role === "admin"
                        ? "destructive"
                        : user.role === "instructor"
                        ? "secondary"
                        : "default"
                    }
                  >
                    {user.role}
                  </Badge>
                </TableCell>
                <TableCell>{formatDate(user.joinedAt)}</TableCell>
                <TableCell>
                  <div className="flex justify-center">
                    <Eye
                      onClick={() => openModal(user.id)}
                      className="w-4 h-4"
                    />
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <div className="md:hidden w-full space-y-4 p-4">
        {users.map((user) => (
          <div
            key={user.id}
            className="border rounded-lg p-4 shadow-sm space-y-2"
          >
            <div className="flex gap-4 mb-6">
              <img
                src={user.avatar}
                alt={user.fullname}
                className="w-12 h-12 rounded-full object-cover border"
              />
              <div className="flex-1 text-start">
                <h3 className="text-base font-semibold">{user.fullname}</h3>
                <p className="text-sm text-muted-foreground">{user.email}</p>
              </div>

              <div>
                <Badge
                  variant={
                    user.role === "admin"
                      ? "destructive"
                      : user.role === "instructor"
                      ? "secondary"
                      : "default"
                  }
                >
                  {user.role}
                </Badge>
              </div>
            </div>

            <div className="flex items-center justify-between ">
              <div className="text-xs text-start">
                <p className="text-muted-foreground">Joined: {user.joinedAt}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
