import { Trash2 } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { FormDelete } from "@/components/form/FormDelete";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const DeleteTemplate = ({ template }) => {
  const { deleteTemplate } = useScheduleTemplateMutation();

  const handleDeleteTemplate = () => {
    deleteTemplate.mutate(template.id);
  };

  return (
    <FormDelete
      title="Delete Recuring Schedule"
      onDelete={handleDeleteTemplate}
      loading={deleteTemplate.isPending}
      description="Are you sure want to delete this Recuring Schedule ?"
      buttonElement={
        <Button variant="destructive" className="w-full" type="button">
          <span>Delete</span>
          <Trash2 className="w-4 h-4" />
        </Button>
      }
    />
  );
};

export { DeleteTemplate };
