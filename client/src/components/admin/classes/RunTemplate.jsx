import { FormToggle } from "@/components/form/FormToggle";
import { useScheduleTemplateMutation } from "@/hooks/useSchedules";

const RunTemplate = ({ template }) => {
  const { runTemplate } = useScheduleTemplateMutation();

  const handleRunTemplate = () => {
    runTemplate.mutate(template.id);
  };

  return (
    <FormToggle
      type="start"
      loading={runTemplate.isPending}
      title="Run Recuring Schedule"
      onToggle={handleRunTemplate}
      description="Are you sure want to Run this recuring schedule ?"
    />
  );
};

export { RunTemplate };
