import { localizer } from "@/lib/utils";
import ClassScheduleHead from "./ClassScheduleHead";
import { Calendar, Views } from "react-big-calendar";
import "react-big-calendar/lib/css/react-big-calendar.css";

const ClassScheduleCalendar = ({ onEmptySlotClick, onSelectEvent, events }) => {
  return (
    <section>
      {/* Calendar */}
      <div className="mt-6 rounded-xl bg-muted border border-border shadow-sm overflow-hidden">
        <Calendar
          components={{
            toolbar: ClassScheduleHead,
          }}
          selectable
          localizer={localizer}
          events={events}
          startAccessor="start"
          endAccessor="end"
          defaultView={Views.WEEK}
          views={["week"]}
          min={new Date(1970, 1, 1, 8, 0)}
          max={new Date(1970, 1, 1, 17, 0)}
          showMultiDayTimes={true}
          allDayAccessor={false}
          showAllDayEvents={false}
          style={{ height: "calc(100vh - 180px)" }}
          onSelectSlot={(slotInfo) => {
            const dateTime = new Date(slotInfo.start);
            onEmptySlotClick(dateTime);
          }}
          onSelectEvent={onSelectEvent}
          eventPropGetter={(event) => {
            const bgColor = event?.resource?.color || "hsl(var(--secondary))";
            return {
              style: {
                backgroundColor: bgColor,
                color: "white",
                borderRadius: "0.5rem",
                padding: "1px 1px",
                fontSize: "11px",
                border: "none",
              },
            };
          }}
        />
      </div>
    </section>
  );
};

export { ClassScheduleCalendar };
