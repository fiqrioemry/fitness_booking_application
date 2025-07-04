import { Pencil } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { updateTemplateSchema } from "@/lib/schema";
import { SelectElement } from "@/components/input/SelectElement";
import { operationHours, operationMinutes } from "@/lib/constant";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";
import { FormUpdateSheet } from "@/components/form/FormUpdateSheet";
import { DaySelectorElement } from "@/components/input/DaySelectorElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";
import { SelectCalendarElement } from "@/components/input/SelectCalendarElement";

const UpdateTemplate = ({ template }) => {
  const { updateTemplate } = useScheduleTemplateMutation();

  return (
    <FormUpdateSheet
      state={template}
      schema={updateTemplateSchema}
      title="Update Recurence Schedule"
      loading={updateTemplate.isPending}
      action={updateTemplate.mutateAsync}
      buttonElement={
        <Button variant="secondary" className="w-full" type="button">
          <span>Update</span>
          <Pencil className="w-4 h-4" />
        </Button>
      }
    >
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
      <SelectCalendarElement name="endDate" label="End Date" />
      <DaySelectorElement name="dayOfWeeks" label="Recurring Days" />
      <div className="grid grid-cols-3 gap-4">
        <SelectElement
          name="startHour"
          label="Start Hour"
          isNumeric={true}
          placeholder="Hour"
          options={operationHours}
        />
        <SelectElement
          isNumeric={true}
          name="startMinute"
          label="Start Minute"
          placeholder="Minute"
          options={operationMinutes}
        />
        <InputNumberElement name="capacity" label="Capacity" />
      </div>
    </FormUpdateSheet>
  );
};

export { UpdateTemplate };
