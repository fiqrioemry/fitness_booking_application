import { checkoutState } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { checkoutSchema } from "@/lib/schema";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { useCheckoutBookingMutation } from "@/hooks/useBooking";
import { InputTextElement } from "@/components/input/InputTextElement";

export const CheckoutClass = ({ bookings }) => {
  const { mutate: checkout, isPending } = useCheckoutBookingMutation();

  const handleCheckoutClass = (data) => {
    checkout({ id: bookings.id, data });
  };

  return (
    <FormAddDialog
      loading={isPending}
      state={checkoutState}
      schema={checkoutSchema}
      title="Checkout from class"
      action={handleCheckoutClass}
      buttonElement={
        <Button disabled={bookings.checkedOut}>
          <span>checkout</span>
        </Button>
      }
    >
      <InputTextElement
        maxLength={6}
        label="Checkout code"
        name="verificationCode"
        placeholder="Input a valid checkout code"
      />
    </FormAddDialog>
  );
};
