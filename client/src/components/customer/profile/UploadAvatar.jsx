import { Camera } from "lucide-react";
import { avatarSchema } from "@/lib/schema";
import { Button } from "@/components/ui/Button";
import { useProfileMutation } from "@/hooks/useProfile";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

export const UploadAvatar = ({ profile }) => {
  const { updateAvatar } = useProfileMutation();

  return (
    <div className="relative flex flex-col justify-center items-center space-y-4">
      <img
        src={profile.avatar}
        alt={profile.fullname}
        className="w-32 h-32 rounded-full object-cover border"
      />
      <FormAddDialog
        icon={true}
        state={profile}
        title="Edit Avatar"
        schema={avatarSchema}
        loading={updateAvatar.isPending}
        action={updateAvatar.mutateAsync}
        buttonElement={
          <Button className="w-full" type="button">
            <Camera className="w-4 h-4" />
            <span>Change Avatar</span>
          </Button>
        }
      >
        <InputFileElement
          isSingle
          name="avatar"
          label="Upload Avatar"
          note="Recommend: ratio 1:1 (400x400px)"
        />
      </FormAddDialog>
    </div>
  );
};
