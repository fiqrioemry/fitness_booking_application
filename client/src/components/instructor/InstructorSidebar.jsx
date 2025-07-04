import { User, LogOut, Package } from "lucide-react";
import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/Sidebar";
import { cn } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuContent,
} from "@/components/ui/DropdownMenu";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/Avatar";

const NavItem = ({ to, icon: Icon, title, active }) => (
  <Link
    to={to}
    className={cn(
      "flex items-center gap-3 px-4 py-2 text-sm rounded-md transition",
      active
        ? "bg-primary text-primary-foreground font-semibold"
        : "text-muted-foreground hover:bg-muted"
    )}
  >
    <Icon className="w-4 h-4" />
    {title}
  </Link>
);

const InstructorSidebar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const { user, logout } = useAuthStore();

  return (
    <Sidebar>
      <SidebarContent className="p-4 space-y-4 text-sm">
        <SidebarHeader className="mb-4">
          <WebLogo />
        </SidebarHeader>

        <SidebarMenu className="space-y-1">
          <NavItem
            to="/instructor/schedules"
            title="My Schedule"
            icon={Package}
            active={currentPath === "/instructor/schedules"}
          />
        </SidebarMenu>
      </SidebarContent>

      <SidebarFooter className="p-4 text-xs text-muted-foreground">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3 cursor-pointer hover:bg-muted px-3 py-2 rounded-md transition">
              <Avatar className="w-9 h-9">
                <AvatarImage src={user?.avatar} alt={user?.fullname} />
                <AvatarFallback>{user?.fullname?.[0]}</AvatarFallback>
              </Avatar>
              <div className="flex flex-col text-left overflow-hidden">
                <span className="text-sm font-medium text-foreground truncate">
                  {user?.fullname}
                </span>
                <span className="text-xs text-muted-foreground truncate">
                  {user?.email}
                </span>
              </div>
            </div>
          </DropdownMenuTrigger>

          <DropdownMenuContent side="top" align="start" className="w-60">
            <DropdownMenuItem onClick={logout}>
              <LogOut className="w-4 h-4 mr-2" />
              Logout
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarFooter>
    </Sidebar>
  );
};

export default InstructorSidebar;
