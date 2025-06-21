import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/useAuthStore";

const RoleRoute = ({ children, allow, redirect = "/" }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();

  const isAllowed = allow.includes(user?.role);

  useEffect(() => {
    if (!isAllowed) navigate(redirect);
  }, [isAllowed, navigate, redirect]);

  return isAllowed ? children : null;
};

const PublicRoute = ({ children }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();

  useEffect(() => {
    if (user?.role === "admin") navigate("/admin/dashboard");
    else if (user?.role === "instructor") navigate("/instructor");
  }, [user, navigate]);

  return user?.role === "admin" || user?.role === "instructor"
    ? null
    : children;
};

const NonAuthRoute = ({ children }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();

  useEffect(() => {
    if (user) navigate("/");
  }, [user, navigate]);

  return user ? null : children;
};

export const AuthRoute = ({ children }) => (
  <RoleRoute allow={["customer"]}>{children}</RoleRoute>
);
export const AdminRoute = ({ children }) => (
  <RoleRoute allow={["admin"]}>{children}</RoleRoute>
);
export const InstructorRoute = ({ children }) => (
  <RoleRoute allow={["instructor"]}>{children}</RoleRoute>
);
export { PublicRoute, NonAuthRoute };
