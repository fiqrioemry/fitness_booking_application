import { instructorSchema } from "@/lib/schema";
import { instructorState } from "@/lib/constant";
import { useInstructorMutation } from "@/hooks/useInstructor";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { SelectUsersElement } from "@/components/input/SelectUsersElement";

const AddInstructor = () => {
  const { createInstructor } = useInstructorMutation();

  return (
    <FormAddDialog
      className="w-72"
      state={instructorState}
      schema={instructorSchema}
      title="Appoint new instructor"
      action={createInstructor.mutateAsync}
      isLoading={createInstructor.isPending}
    >
      <SelectUsersElement
        data="users"
        name="userId"
        label="Instructor name"
        placeholder="Select instructor name"
      />
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
    </FormAddDialog>
  );
};

export { AddInstructor };
