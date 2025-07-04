/* eslint-disable react/prop-types */
import { FormDelete } from "@/components/form/FormDelete";
import { useInstructorMutation } from "@/hooks/useInstructor";

const DeleteInstructor = ({ instructor }) => {
  const { deleteInstructor } = useInstructorMutation();

  const handleDeleteInstructor = () => {
    deleteInstructor.mutate(instructor.id);
  };

  return (
    <FormDelete
      title="Delete Instrutors"
      onDelete={handleDeleteInstructor}
      loading={deleteInstructor.isPending}
      description="Are you sure want to remove this Instructors ?"
    />
  );
};

export { DeleteInstructor };
