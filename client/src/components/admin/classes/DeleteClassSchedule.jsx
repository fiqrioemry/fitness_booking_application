import { FormDelete } from "@/components/form/FormDelete";
import { useScheduleMutation } from "@/hooks/useSchedules";

const DeleteClassSchedule = ({ schedule, onClose }) => {
  const { deleteSchedule } = useScheduleMutation();

  const handleDeleteSchedule = async () => {
    await deleteSchedule.mutateAsync(schedule.id);
    onClose();
  };

  return (
    <FormDelete
      title="Delete Schedule"
      onDelete={handleDeleteSchedule}
      loading={deleteSchedule.isPending}
      description="Are you sure want to delete this Schedule?"
    />
  );
};

export { DeleteClassSchedule };
