import { Pagination } from "@/components/ui/Pagination";
import { useBookingStore } from "@/store/useBookingStore";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { useInstructorSchedulesQuery } from "@/hooks/useSchedules";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/Tabs";
import { NoClassSchedule } from "@/components/instructor/schedules/NoClassSchedule";
import { PastClassSchedules } from "@/components/instructor/schedules/PastClassSchedules";
import { UpcomingClassSchedules } from "@/components/instructor/schedules/UpcomingClassSchedules";

const InstructorSchedule = () => {
  const { page, limit, sort, status, setStatus, setPage } = useBookingStore();

  const { data, isError, refetch, isLoading } = useInstructorSchedulesQuery({
    status,
    sort,
    page,
    limit,
  });

  const bookings = data?.data || [];
  const pagination = data?.pagination || null;

  const isPast = status === "past";
  const isUpcoming = status === "upcoming";

  console.log(data);

  return (
    <section className="section p-8 space-y-6">
      <SectionTitle
        title="My Class Schedules"
        description="View and manage your upcoming and past class schedules as instructors."
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
                <NoClassSchedule type="upcoming" />
              ) : (
                <div className="grid gap-6">
                  {isUpcoming &&
                    bookings.map((schedule) => (
                      <UpcomingClassSchedules
                        key={schedule.id}
                        schedule={schedule}
                      />
                    ))}
                </div>
              )}
            </TabsContent>

            <TabsContent value="past">
              {isPast && bookings.length === 0 ? (
                <NoClassSchedule type="past" />
              ) : (
                <div className="grid gap-6">
                  {isPast &&
                    bookings.map((schedule) => (
                      <PastClassSchedules
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

      {pagination && pagination.totalRows > 5 && (
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

export default InstructorSchedule;
