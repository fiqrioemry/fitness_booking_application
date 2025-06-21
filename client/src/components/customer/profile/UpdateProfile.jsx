import { PencilIcon } from "lucide-react";
import { profileSchema } from "@/lib/schema";
import { genderOptions } from "@/lib/constant";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { SelectElement } from "@/components/input/SelectElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputTextElement } from "@/components/input/InputTextElement";

export const UpdateProfile = ({ profile, edit = "" }) => {
  const { updateProfile } = useProfileMutation();

  return (
    <FormAddDialog
      state={profile}
      title="Edit Profile"
      schema={profileSchema}
      loading={updateProfile.isPending}
      action={updateProfile.mutateAsync}
      buttonElement={<PencilIcon className="w-4 h-4 " />}
    >
      {edit === "fullname" && (
        <InputTextElement name="fullname" label="Full Name" />
      )}
      {edit === "birthday" && (
        <InputDateElement name="birthday" label="Date of Birth" />
      )}
      {edit === "gender" && (
        <SelectElement
          name="gender"
          label="Gender"
          options={genderOptions}
          placeholder="Select gender"
          rules={{ required: false }}
        />
      )}
      {edit === "bio" && (
        <InputTextareaElement
          name="bio"
          label="Your bio"
          maxLength={200}
          rules={{ required: false }}
          placeholder="Enter your bio"
        />
      )}
      {edit === "phone" && (
        <InputTextElement
          isNumeric
          name="phone"
          label="Phone Number"
          placeholder="08xxxx"
          maxLength={13}
          rules={{ required: false }}
        />
      )}
    </FormAddDialog>
  );
};
