import {
  SidebarInset,
  SidebarTrigger,
  SidebarProvider,
} from "@/components/ui/Sidebar";
import { MenuIcon } from "lucide-react";
import { Outlet } from "react-router-dom";
import InstructorSidebar from "./InstructorSidebar";
import { Separator } from "@/components/ui/Separator";

const InstructorLayout = () => {
  return (
    <SidebarProvider>
      <InstructorSidebar />
      <SidebarInset>
        <header className="flex sticky top-0 bg-background h-16 shrink-0 items-center gap-2 border-b px-4 z-50">
          <SidebarTrigger className="-ml-1">
            <MenuIcon className="text-muted-foreground" />
          </SidebarTrigger>
          <Separator orientation="vertical" className="mr-2 h-4" />
          <h1 className="text-lg font-semibold text-muted-foreground hidden sm:block">
            Dashboard
          </h1>
        </header>

        <div className="flex flex-1 flex-col bg-muted">
          <Outlet />
        </div>
      </SidebarInset>
    </SidebarProvider>
  );
};

export default InstructorLayout;
