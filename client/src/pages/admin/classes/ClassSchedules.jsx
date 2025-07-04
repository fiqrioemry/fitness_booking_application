import { useState, useEffect } from "react";
import { useSchedulesQuery } from "@/hooks/useSchedules";
import { SectionSkeleton } from "@/components/loading/SectionSkeleton";
import { AddClassSchedule } from "@/components/admin/classes/AddClassSchedule";
import { ClassScheduleDetail } from "@/components/admin/classes/ClassScheduleDetail";
import { ClassScheduleCalendar } from "@/components/admin/classes/ClassScheduleCalendar";

const ClassSchedule = () => {
  const [events, setEvents] = useState([]);
  const [selectedDate, setSelectedDate] = useState(null);
  const [selectedEvent, setSelectedEvent] = useState(null);
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false);
  const [isDetailDialogOpen, setIsDetailDialogOpen] = useState(false);
  const { data: schedules = [], isLoading } = useSchedulesQuery();

  const handleEmptySlotClick = (dateTime) => {
    setSelectedDate(dateTime);
    setIsAddDialogOpen(true);
  };

  const handleSelectEvent = (event) => {
    setSelectedEvent(event);
    setIsDetailDialogOpen(true);
  };

  useEffect(() => {
    if (!schedules || schedules.length === 0) return;

    const mapped = schedules?.map((item) => {
      const start = new Date(item.date);
      start.setHours(item.startHour);
      start.setMinutes(item.startMinute);

      const end = new Date(start);
      end.setMinutes(end.getMinutes() + 60);

      return {
        id: item.id,
        title: `${item.className} - ${item.instructorName} (${item.bookedCount}/${item.capacity})`,
        start,
        end,
        allDay: false,
        resource: {
          ...item,
        },
      };
    });

    setEvents(mapped);
  }, [schedules]);

  if (isLoading) return <SectionSkeleton />;

  return (
    <section>
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold py-4">Class Schedules Event</h2>
      </div>
      <AddClassSchedule
        open={isAddDialogOpen}
        setOpen={setIsAddDialogOpen}
        defaultDateTime={selectedDate}
      />

      <ClassScheduleDetail
        open={isDetailDialogOpen}
        onClose={() => setIsDetailDialogOpen(false)}
        event={selectedEvent}
      />
      <div className="w-full">
        <ClassScheduleCalendar
          events={events}
          schedules={schedules}
          onSelectEvent={handleSelectEvent}
          onEmptySlotClick={handleEmptySlotClick}
        />
      </div>
    </section>
  );
};

export default ClassSchedule;
