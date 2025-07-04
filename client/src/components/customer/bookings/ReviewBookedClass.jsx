import { reviewState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { createReviewSchema } from "@/lib/schema";
import { useCreateReviewMutation } from "@/hooks/useReview";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputRatingElement } from "@/components/input/InputRatingElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

export const ReviewBookedClass = ({ id }) => {
  const { mutate: createReview, isPending } = useCreateReviewMutation();

  const handleReviewClass = (data) => {
    createReview({ id, data });
  };

  return (
    <FormAddDialog
      loading={isPending}
      state={reviewState}
      title="Create a comment"
      schema={createReviewSchema}
      action={handleReviewClass}
      buttonElement={
        <Button variant="outline">
          <span>Review Class</span>
        </Button>
      }
    >
      <InputRatingElement name="rating" label="Your Rating" />
      <InputTextareaElement
        name="comment"
        label="Comment"
        placeholder="Write your comment here"
      />
    </FormAddDialog>
  );
};
