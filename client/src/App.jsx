// public pages
import Home from "./pages/Home";
import About from "./pages/About";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import Classes from "./pages/Classes";
import Packages from "./pages/Packages";
import NotFound from "./pages/NotFound";
import Schedules from "./pages/Schedules";
import ClassDetail from "./pages/ClassDetail";
import PackageDetail from "./pages/PackageDetail";
import ScheduleDetail from "./pages/ScheduleDetail";

// admin pages
import UsersList from "./pages/admin/UsersList";
import Dashboard from "./pages/admin/Dashboard";
import ClassAdd from "./pages/admin/classes/ClassAdd";
import Notifications from "./pages/admin/Notifications";
import PackageAdd from "./pages/admin/packages/PackageAdd";
import ClassesList from "./pages/admin/classes/ClassesList";
import InstructorsList from "./pages/admin/InstructorsList";
import VouchersAdd from "./pages/admin/vouchers/VouchersAdd";
import ClassOptions from "./pages/admin/classes/ClassOptions";
import TransactionsList from "./pages/admin/TransactionsList";
import VouchersList from "./pages/admin/vouchers/VouchersList";
import PackagesList from "./pages/admin/packages/PackagesList";
import ClassRecuring from "./pages/admin/classes/ClassRecuring";
import ClassSchedules from "./pages/admin/classes/ClassSchedules";
import { UserDetailDialog } from "./components/admin/users/UserDetailDialog";

// customer pages
import UserProfile from "./pages/customer/UserProfile";
import UserBookings from "./pages/customer/UserBookings";
import UserPackages from "./pages/customer/UserPackages";
import UserSettings from "./pages/customer/UserSettings";
import UserTransactions from "./pages/customer/UserTransactions";
import UserNotifications from "./pages/customer/UserNotifications";
import { BookedScheduleDetail } from "./components/customer/bookings/BookedScheduleDetail";

// instructor pages
import InstructorSchedules from "./pages/instructor/InstructorSchedules";
import { StartClassSchedule } from "./components/instructor/schedules/StartClassSchedule";
import { ClassAttendanceDetail } from "./components/instructor/schedules/ClassAttendanceDetail";

// route config & support
import { Toaster } from "sonner";
import { Loading } from "@/components/ui/Loading";
import { ScrollToTop } from "./hooks/useScrollToTop";
import { Navigate, Route, Routes, useLocation } from "react-router-dom";
import {
  AuthRoute,
  AdminRoute,
  PublicRoute,
  NonAuthRoute,
  InstructorRoute,
} from "./middleware";

// layout pages
import { useEffect } from "react";
import { useAuthStore } from "@/store/useAuthStore";
import AdminLayout from "./components/admin/AdminLayout";
import PublicLayout from "./components/public/PublicLayout";
import CustomerLayout from "./components/customer/CustomerLayout";
import InstructorLayout from "./components/instructor/InstructorLayout";
import { MyTransactionDetail } from "./components/customer/transactions/MyTransactionDetail";
import TransactionDetail from "./pages/admin/TransactionDetail";

function App() {
  const location = useLocation();
  const state = location.state;
  const backgroundLocation = state?.backgroundLocation;
  const { checkingAuth, authMe } = useAuthStore();

  useEffect(() => {
    authMe();
  }, []);

  if (checkingAuth) return <Loading />;

  return (
    <>
      <Toaster position="top-center" />
      <ScrollToTop />
      <Routes location={backgroundLocation || location}>
        {/* Auth Routes (No Layout) */}
        <Route
          path="/signin"
          element={
            <NonAuthRoute>
              <SignIn />
            </NonAuthRoute>
          }
        />
        <Route
          path="/signup"
          element={
            <NonAuthRoute>
              <SignUp />
            </NonAuthRoute>
          }
        />

        {/* Public Routes */}
        <Route
          path="/"
          element={
            <PublicRoute>
              <PublicLayout />
            </PublicRoute>
          }
        >
          <Route index element={<Home />} />
          <Route path="about" element={<About />} />
          <Route path="classes" element={<Classes />} />
          <Route path="packages" element={<Packages />} />
          <Route path="schedules" element={<Schedules />} />
          <Route path="classes/:id" element={<ClassDetail />} />
          <Route path="packages/:id" element={<PackageDetail />} />
          <Route
            path="schedules/:id"
            element={
              <AuthRoute>
                <ScheduleDetail />
              </AuthRoute>
            }
          />
          {/* 404 for public routes */}
          <Route path="*" element={<NotFound />} />
        </Route>

        {/* Customer Routes */}
        <Route
          path="/profile"
          element={
            <AuthRoute>
              <CustomerLayout />
            </AuthRoute>
          }
        >
          <Route index element={<UserProfile />} />
          <Route path="packages" element={<UserPackages />} />
          <Route path="bookings" element={<UserBookings />} />
          <Route path="settings" element={<UserSettings />} />
          <Route path="transactions" element={<UserTransactions />} />
          <Route path="notifications" element={<UserNotifications />} />
          {/* 404 for profile routes */}
          <Route path="*" element={<NotFound />} />
        </Route>

        {/* Admin Routes */}
        <Route
          path="/admin"
          element={
            <AdminRoute>
              <AdminLayout />
            </AdminRoute>
          }
        >
          <Route index element={<Navigate to="dashboard" replace />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="users" element={<UsersList />} />
          <Route path="messages" element={<Notifications />} />
          <Route path="classes" element={<ClassesList />} />
          <Route path="classes/add" element={<ClassAdd />} />
          <Route path="classes/options" element={<ClassOptions />} />
          <Route path="classes/recuring" element={<ClassRecuring />} />
          <Route path="classes/schedules" element={<ClassSchedules />} />
          <Route path="vouchers" element={<VouchersList />} />
          <Route path="vouchers/add" element={<VouchersAdd />} />
          <Route path="packages" element={<PackagesList />} />
          <Route path="packages/add" element={<PackageAdd />} />
          <Route path="instructors" element={<InstructorsList />} />
          <Route path="transactions" element={<TransactionsList />} />
          <Route path="transactions/:id" element={<TransactionDetail />} />
          {/* 404 for admin routes */}
          <Route path="*" element={<NotFound />} />
        </Route>

        {/* Instructor Routes */}
        <Route
          path="/instructor"
          element={
            <InstructorRoute>
              <InstructorLayout />
            </InstructorRoute>
          }
        >
          <Route index element={<Navigate to="schedules" replace />} />
          <Route path="schedules" element={<InstructorSchedules />} />
          {/* 404 for instructor routes */}
          <Route path="*" element={<NotFound />} />
        </Route>

        {/* Global 404 - Catch all unmatched routes */}
        <Route path="*" element={<NotFound />} />
      </Routes>

      {/* Background Dialog Routes */}
      {backgroundLocation && (
        <Routes>
          {/* Admin Dialogs */}
          <Route
            path="/admin/users/:id"
            element={
              <AdminRoute>
                <UserDetailDialog />
              </AdminRoute>
            }
          />

          {/* Instructor Dialogs */}
          <Route
            path="/instructor/schedules/:id/attendance"
            element={
              <InstructorRoute>
                <ClassAttendanceDetail />
              </InstructorRoute>
            }
          />
          <Route
            path="/instructor/schedules/:id/open"
            element={
              <InstructorRoute>
                <StartClassSchedule />
              </InstructorRoute>
            }
          />

          {/* Customer Dialogs */}
          <Route
            path="/profile/bookings/:id"
            element={
              <AuthRoute>
                <BookedScheduleDetail />
              </AuthRoute>
            }
          />
          <Route
            path="/profile/transactions/:id"
            element={
              <AuthRoute>
                <MyTransactionDetail />
              </AuthRoute>
            }
          />
        </Routes>
      )}
    </>
  );
}

export default App;
