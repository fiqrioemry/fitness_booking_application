import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/Dialog";
import { openClassSchema } from "@/lib/schema";
import { openClassState } from "@/lib/constant";
import { FormInput } from "@/components/form/FormInput";
import { useNavigate, useParams } from "react-router-dom";
import { useOpenScheduleMutation } from "@/hooks/useSchedules";
import { InputTextElement } from "@/components/input/InputTextElement";

export const StartClassSchedule = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { mutateAsync: openClass, isPending } = useOpenScheduleMutation();

  const handleOpenClass = async (data) => {
    await openClass({ id, data });
    navigate(-1);
  };

  return (
    <Dialog open={true} onOpenChange={() => navigate(-1)}>
      <DialogContent className="max-w-xl w-full">
        <DialogHeader>
          <DialogTitle>Start Class Schedule</DialogTitle>
        </DialogHeader>
        <FormInput
          text="Start Class"
          className="w-full"
          isLoading={isPending}
          state={openClassState}
          schema={openClassSchema}
          action={handleOpenClass}
        >
          <InputTextElement
            name="verificationCode"
            maxLength={6}
            label="Verification Code (For checkout)"
            placeholder="Input 6 randomg string + number for verification code"
          />
          <InputTextElement
            name="zoomLink"
            label="Zoom Link *(Optional for online classes)"
            placeholder="Place your zoom meeting link here"
          />
        </FormInput>
      </DialogContent>
    </Dialog>
  );
};
