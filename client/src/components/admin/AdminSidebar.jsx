import {
  Plus,
  List,
  Bell,
  Clock,
  Users,
  Package,
  Dumbbell,
  UserPlus,
  BookUser,
  Calendar,
  BarChart2,
  LayoutList,
  ShoppingCart,
  TicketPercent,
  LogOut,
  Repeat,
  PlusSquareIcon,
  MessageCircle,
  MessageSquare,
  MailOpen,
} from "lucide-react";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "@/components/ui/Accordion";

import {
  Sidebar,
  SidebarMenu,
  SidebarHeader,
  SidebarFooter,
  SidebarContent,
} from "@/components/ui/Sidebar";
import { cn } from "@/lib/utils";
import { useAuthStore } from "@/store/useAuthStore";
import { Link, useLocation } from "react-router-dom";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/Avatar";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/DropdownMenu";

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
const directMenus = [
  {
    to: "/admin/users",
    icon: Users,
    title: "Users",
  },
  {
    to: "/admin/transactions",
    icon: ShoppingCart,
    title: "Transactions",
  },
  {
    to: "/admin/messages",
    icon: MailOpen,
    title: "Message",
  },
  {
    to: "/admin/instructors",
    icon: UserPlus,
    title: "Instructors",
  },
];

const accordionMenus = [
  {
    value: "vouchers",
    icon: TicketPercent,
    title: "Vouchers",
    children: [
      { to: "/admin/vouchers", title: "Vouchers List", icon: List },
      {
        to: "/admin/vouchers/add",
        title: "New Vouchers",
        icon: PlusSquareIcon,
      },
    ],
  },
  {
    value: "classes",
    icon: Dumbbell,
    title: "Classes",
    children: [
      { to: "/admin/classes", title: "Class List", icon: List },
      { to: "/admin/classes/add", title: "Add Class", icon: Plus },
      { to: "/admin/classes/options", title: "Class Options", icon: Calendar },
      { to: "/admin/classes/schedules", title: "Class Schedules", icon: Clock },
      {
        to: "/admin/classes/recuring",
        title: "Recuring Schedule",
        icon: Repeat,
      },
    ],
  },
  {
    value: "packages",
    icon: Package,
    title: "Packages",
    children: [
      { to: "/admin/packages", title: "Package List", icon: List },
      { to: "/admin/packages/add", title: "Add Package", icon: Plus },
    ],
  },
];

const AdminSidebar = () => {
  const location = useLocation();
  const currentPath = location.pathname;
  const { user, logout } = useAuthStore();

  return (
    <Sidebar>
      <SidebarContent className="px-4 space-y-4 text-sm text-gray-700">
        <SidebarHeader className="mb-4 py-2">
          <img src="/logo.png" />
        </SidebarHeader>

        <SidebarMenu className="space-y-1">
          <NavItem
            to="/admin"
            title="Dashboard"
            icon={BarChart2}
            active={currentPath === "/admin"}
          />

          {directMenus.map((item) => (
            <NavItem
              key={item.to}
              to={item.to}
              icon={item.icon}
              title={item.title}
              active={currentPath === item.to}
            />
          ))}

          <Accordion type="multiple" className="space-y-1">
            {accordionMenus.map((menu) => (
              <AccordionItem key={menu.value} value={menu.value}>
                <AccordionTrigger
                  className={cn(
                    "w-full px-4 py-2 text-sm rounded-md transition flex items-center gap-2",
                    "text-muted-foreground hover:bg-muted [&[data-state=open]]:bg-muted"
                  )}
                >
                  <menu.icon className="w-4 h-4" />
                  {menu.title}
                </AccordionTrigger>

                <AccordionContent className="pl-6 space-y-1 mt-1">
                  {menu.children.map((child) => (
                    <NavItem
                      key={child.to}
                      to={child.to}
                      title={child.title}
                      icon={child.icon}
                      active={currentPath === child.to}
                    />
                  ))}
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
        </SidebarMenu>
      </SidebarContent>
      <SidebarFooter className="p-4 text-xs text-muted-foreground">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div className="flex items-center gap-3 cursor-pointer hover:bg-muted px-3 py-2 rounded-md transition">
              <Avatar className="w-9 h-9">
                <AvatarImage src={user?.avatar} alt={user?.fullname} />
                <AvatarFallback>{user?.fullname?.[0] || "A"}</AvatarFallback>
              </Avatar>
              <div className="flex flex-col text-left overflow-hidden">
                <span className="text-sm font-medium text-foreground truncate">
                  {user?.fullname || "Admin"}
                </span>
                <span className="text-xs text-muted-foreground truncate">
                  {user?.email || "admin@gmail.com"}
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

export default AdminSidebar;
