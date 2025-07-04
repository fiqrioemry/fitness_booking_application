import { Pagination } from "@/components/ui/Pagination";
import { useMyBookingsQuery } from "@/hooks/useBooking";
import { useBookingStore } from "@/store/useBookingStore";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/Tabs";
import { NoBookedSchedule } from "@/components/customer/bookings/NoBookedSchedule";
import { PastBookedSchedules } from "@/components/customer/bookings/PastBookedSchedules";
import { UpcomingBookedSchedules } from "@/components/customer/bookings/UpcomingBookedSchedules";

const UserBookings = () => {
  const { page, limit, sort, status, setStatus, setPage } = useBookingStore();

  const { data, isError, refetch, isLoading } = useMyBookingsQuery({
    status,
    sort,
    page,
    limit,
  });

  const bookings = data?.data || [];
  const pagination = data?.pagination || null;

  const isPast = status === "past";
  const isUpcoming = status === "upcoming";

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Bookings"
        description="View and manage your upcoming and past class bookings."
      />

      <Tabs defaultValue="upcoming" className="w-full">
        <TabsList className="mb-4">
          <TabsTrigger onClick={() => setStatus("upcoming")} value="upcoming">
            Upcoming
          </TabsTrigger>
          <TabsTrigger onClick={() => setStatus("past")} value="past">
            Past
          </TabsTrigger>
        </TabsList>

        {isLoading ? (
          <SectionSkeleton />
        ) : isError ? (
          <ErrorDialog onRetry={refetch} />
        ) : (
          <>
            <TabsContent value="upcoming">
              {isUpcoming && bookings.length === 0 ? (
                <NoBookedSchedule type="upcoming" />
              ) : (
                <div className="grid gap-6">
                  {isUpcoming &&
                    bookings.map((schedule) => (
                      <UpcomingBookedSchedules
                        key={schedule.id}
                        schedule={schedule}
                      />
                    ))}
                </div>
              )}
            </TabsContent>

            <TabsContent value="past">
              {isPast && bookings.length === 0 ? (
                <NoBookedSchedule type="past" />
              ) : (
                <div className="grid gap-6">
                  {isPast &&
                    bookings.map((schedule) => (
                      <PastBookedSchedules
                        key={schedule.id}
                        schedule={schedule}
                      />
                    ))}
                </div>
              )}
            </TabsContent>
          </>
        )}
      </Tabs>

      {pagination && (
        <Pagination
          page={pagination.page}
          onPageChange={setPage}
          limit={pagination.limit}
          total={pagination.totalRows}
        />
      )}
    </section>
  );
};

export default UserBookings;
