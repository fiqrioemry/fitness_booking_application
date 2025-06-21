import { operationHours } from "@/lib/constant";
import { operationMinutes } from "@/lib/constant";
import { updateScheduleSchema } from "@/lib/schema";
import { useScheduleMutation } from "@/hooks/useSchedules";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateSheet } from "@/components/form/FormUpdateSheet";
import { ColorPickerElement } from "@/components/input/ColorPickerElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateClassSchedule = ({ schedule, onClose }) => {
  const { updateSchedule } = useScheduleMutation();
  const { isPending, mutateAsync } = updateSchedule;

  const handleUpdateSchedule = async (data) => {
    await mutateAsync({ id: schedule.id, ...data });
    onClose();
  };

  return (
    <FormUpdateSheet
      state={schedule}
      loading={isPending}
      title="Update Schedule"
      schema={updateScheduleSchema}
      action={handleUpdateSchedule}
    >
      <SelectCalendarElement name="date" label="Event Date" />
      <ColorPickerElement name="color" label="Cardboard Color" />
      <SelectOptionsElement
        data="class"
        name="classId"
        label="Class"
        placeholder="Select option for class"
      />
      <SelectOptionsElement
        data="instructor"
        label="Instructor"
        name="instructorId"
        placeholder="select option for instructor"
      />
      <div className="grid grid-cols-2 gap-4">
        <SelectElement
          name="startHour"
          isNumeric={true}
          label="Start Hour"
          options={operationHours}
          placeholder="Select Hour"
        />

        <SelectElement
          isNumeric={true}
          name="startMinute"
          label="Start Minute"
          options={operationMinutes}
          placeholder="Select Minute"
        />
      </div>
      <InputNumberElement name="capacity" label="Capacity" min={1} />
    </FormUpdateSheet>
  );
};

export { UpdateClassSchedule };
