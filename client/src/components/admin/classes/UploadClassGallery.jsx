import { CirclePlus } from "lucide-react";
import { Button } from "@/components//ui/Button";
import { uploadGallerySchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputFileElement } from "@/components/input/InputFileElement";

const UploadClassGallery = ({ classes }) => {
  const { uploadGallery } = useClassMutation();

  return (
    <FormAddDialog
      title="Upload Gallery"
      schema={uploadGallerySchema}
      loading={uploadGallery.isPending}
      state={{ images: classes.galleries }}
      buttonElement={
        <Button size="icon" variant="secondary">
          <CirclePlus />
        </Button>
      }
      action={(data) =>
        uploadGallery.mutateAsync({ id: classes.id, images: data.images })
      }
    >
      <InputFileElement name="images" label="Gallery (Required)" />
    </FormAddDialog>
  );
};

export { UploadClassGallery };
