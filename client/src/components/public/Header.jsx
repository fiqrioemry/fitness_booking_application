import { Menu } from "lucide-react";
import { NavLink } from "react-router-dom";
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/Button";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { UserDropdown } from "@/components/header/UserDropdown";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/Sheet";

const Header = () => {
  const { user } = useAuthStore();
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 20);
    };

    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const linkClass =
    "hover:text-primary font-medium transition-colors duration-200";
  const activeClass = "text-primary font-semibold";

  return (
    <header
      className={`fixed top-0 w-full z-20 transition-all duration-300 ${
        scrolled ? "bg-background shadow py-2" : "bg-transparent py-4"
      }`}
    >
      <div className="max-w-7xl mx-auto px-4 flex justify-between items-center">
        <WebLogo />

        <div className="hidden md:flex gap-6 items-center">
          <NavLink
            to="/"
            end
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Home
          </NavLink>
          <NavLink
            to="/classes"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Class
          </NavLink>
          <NavLink
            to="/packages"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Package
          </NavLink>
          <NavLink
            to="/schedules"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            Schedule
          </NavLink>
          <NavLink
            to="/about"
            className={({ isActive }) =>
              `${linkClass} ${isActive ? activeClass : ""}`
            }
          >
            About Us
          </NavLink>
          {user ? (
            <UserDropdown />
          ) : (
            <NavLink to="/signin">
              <Button className="px-5 py-2 rounded-full">Get Started</Button>
            </NavLink>
          )}
        </div>
        <div className="flex md:hidden">
          <Sheet>
            <SheetTrigger asChild>
              <Menu className="w-7 h-7" />
            </SheetTrigger>
            <SheetContent side="right" className="p-6">
              <nav className="flex flex-col gap-4">
                <NavLink to="/" end className={linkClass}>
                  Home
                </NavLink>
                <NavLink to="/classes" className={linkClass}>
                  Class
                </NavLink>
                <NavLink to="/packages" className={linkClass}>
                  Package
                </NavLink>
                <NavLink to="/schedules" className={linkClass}>
                  Schedule
                </NavLink>
                <NavLink to="/about" className={linkClass}>
                  About Us
                </NavLink>
                {user ? (
                  <UserDropdown />
                ) : (
                  <NavLink to="/signin">
                    <Button className="w-full mt-2">Get Started</Button>
                  </NavLink>
                )}
              </nav>
            </SheetContent>
          </Sheet>
        </div>
      </div>
    </header>
  );
};

export default Header;
