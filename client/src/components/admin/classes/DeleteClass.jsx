import { useClassMutation } from "@/hooks/useClass";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteClass = ({ classes }) => {
  const { deleteClass } = useClassMutation();

  const handleDeleteClass = () => {
    deleteClass.mutate(classes.id);
  };

  return (
    <FormDelete
      title="Delete Class"
      onDelete={handleDeleteClass}
      loading={deleteClass.isPending}
      description="Are you sure want to delete this classes ?"
    />
  );
};

export { DeleteClass };
