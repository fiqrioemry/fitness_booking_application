import { notificationSchema } from "@/lib/schema";
import { FormInput } from "@/components/form/FormInput";
import { notificationState, typeCode } from "@/lib/constant";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SelectElement } from "@/components/input/SelectElement";
import { useSendPromoNotification } from "@/hooks/useNotification";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const Notifications = () => {
  const { mutate: sendPromo, isPending } = useSendPromoNotification();

  return (
    <section className="section">
      <SectionTitle
        title="Send New Message"
        description="Send a notification through all your users with just one click."
      />

      <div className="bg-background  rounded-xl shadow-sm border p-6">
        <FormInput
          shouldReset
          className="w-full md:w-72"
          state={notificationState}
          schema={notificationSchema}
          text={"Send Promo Message"}
          isLoading={isPending}
          action={sendPromo}
        >
          <div className="space-y-4">
            <InputTextElement
              name="title"
              label="Notification Title"
              placeholder="Enter the title for notifications"
            />
            <InputTextareaElement
              name="message"
              label="Notification Message"
              placeholder="Write a message for notifications"
              maxLength={200}
            />
            <SelectElement
              name="typeCode"
              options={typeCode}
              label="Notification Type"
              placeholder="Select Notification Type"
            />
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default Notifications;
