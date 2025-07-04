import {
  SidebarInset,
  SidebarTrigger,
  SidebarProvider,
} from "@/components/ui/Sidebar";
import { MenuIcon } from "lucide-react";
import AdminSidebar from "./AdminSidebar";
import { Outlet } from "react-router-dom";
import { Separator } from "@/components/ui/Separator";

const AdminLayout = () => {
  return (
    <SidebarProvider>
      <AdminSidebar />
      <SidebarInset>
        <header className="sticky top-0 z-50 flex h-16 items-center gap-2 border-b border-border bg-background px-4">
          <SidebarTrigger className="-ml-1">
            <MenuIcon className="w-5 h-5" />
          </SidebarTrigger>
          <Separator orientation="vertical" className="mr-2 h-4" />
          <h1 className="text-lg font-semibold text-muted-foreground">
            Admin Dashboard
          </h1>
        </header>

        <main className="flex flex-1 flex-col bg-muted">
          <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
};

export default AdminLayout;
