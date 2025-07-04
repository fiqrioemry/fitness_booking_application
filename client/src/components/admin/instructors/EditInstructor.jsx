import { instructorSchema } from "@/lib/schema";
import { useInstructorMutation } from "@/hooks/useInstructor";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";

const EditInstructor = ({ instructor }) => {
  const { updateInstructor } = useInstructorMutation();

  return (
    <FormUpdateDialog
      state={instructor}
      schema={instructorSchema}
      title="Update Instructors"
      loading={updateInstructor.isPending}
      action={updateInstructor.mutateAsync}
    >
      <InputTextElement
        name="specialties"
        label="Specialties"
        placeholder="Add instructor specialties"
      />
      <InputTextElement
        name="certifications"
        label="Certification"
        placeholder="Add instructor certification"
      />
      <InputNumberElement
        name="experience"
        label="Experience"
        placeholder="Enter instructor Experience"
      />
    </FormUpdateDialog>
  );
};

export { EditInstructor };
