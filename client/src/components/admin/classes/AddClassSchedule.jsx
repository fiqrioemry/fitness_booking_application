import { scheduleSchema } from "@/lib/schema";
import { format } from "date-fns";
import { FormSheet } from "@/components/form/FormSheet";
import { useScheduleMutation } from "@/hooks/useSchedules";
import { useFormContext, useWatch } from "react-hook-form";
import { SelectElement } from "@/components/input/SelectElement";
import { SwitchElement } from "@/components/input/SwitchElement";
import { DaySelectorElement } from "@/components/input/DaySelectorElement";
import { ColorPickerElement } from "@/components/input/ColorPickerElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const AddClassSchedule = ({ open, setOpen, defaultDateTime }) => {
  const { createSchedule } = useScheduleMutation();
  const { isPending, mutateAsync } = createSchedule;

  const initialState = {
    classId: "",
    capacity: 0,
    endDate: "",
    dayOfWeeks: [],
    instructorId: "",
    color: "#4ade80",
    isRecurring: false,
    startHour: defaultDateTime?.getHours(),
    startMinute: defaultDateTime?.getMinutes(),
    date: defaultDateTime
      ? format(defaultDateTime, "yyyy-MM-dd").split("T")[0]
      : "",
    format,
  };

  return (
    <FormSheet
      open={open}
      setOpen={setOpen}
      loading={isPending}
      state={initialState}
      action={mutateAsync}
      schema={scheduleSchema}
      title="Create Class Schedule"
    >
      <RecurringSection />
      <ColorPickerElement name="color" label="Cardboard Color" />
      <SelectOptionsElement
        data="class"
        label="Class"
        name="classId"
        placeholder="Select class"
      />
      <SelectOptionsElement
        data="instructor"
        label="Instructor"
        name="instructorId"
        placeholder="Select instructor"
      />
      <div className="grid grid-cols-2 gap-4">
        <SelectElement
          name="startHour"
          label="Start Hour"
          isNumeric={true}
          placeholder="Hour"
          options={[...Array(9)].map((_, i) => 8 + i)}
        />
        <SelectElement
          isNumeric={true}
          name="startMinute"
          label="Start Minute"
          placeholder="Minute"
          options={[0, 15, 30, 45]}
        />
      </div>

      <InputNumberElement name="capacity" label="Capacity" />
    </FormSheet>
  );
};

const RecurringSection = () => {
  const { control } = useFormContext();
  const isRecurring = useWatch({ control, name: "isRecurring" });

  return (
    <>
      <SwitchElement name="isRecurring" label="Repeat weekly?" />
      {!isRecurring && <SelectCalendarElement name="date" label="Event Date" />}
      {isRecurring && (
        <>
          <DaySelectorElement name="dayOfWeeks" label="Days" />
          <SelectCalendarElement name="endDate" label="End Date" />
        </>
      )}
    </>
  );
};

export { AddClassSchedule };
