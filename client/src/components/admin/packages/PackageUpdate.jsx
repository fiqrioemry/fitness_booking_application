import { packageSchema } from "@/lib/schema";
import { usePackageMutation } from "@/hooks/usePackage";
import { SwitchElement } from "@/components/input/SwitchElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { MultiSelectElement } from "@/components/input/MultiSelectElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const PackageUpdate = ({ pkg }) => {
  const { updatePackage } = usePackageMutation();

  return (
    <FormUpdateDialog
      state={{
        ...pkg,
        classIds: pkg.classes?.map((c) => c.id) || [],
      }}
      title="Update Package"
      schema={packageSchema}
      loading={updatePackage.isPending}
      action={updatePackage.mutateAsync}
    >
      <InputTextElement
        name="name"
        label="Package Name"
        placeholder="Enter package name"
      />
      <InputTextareaElement
        maxLength={200}
        name="description"
        label="Package Description"
        placeholder="Enter package description (min. 20 characters)"
      />
      <InputNumberElement
        name="price"
        label="Package Price"
        placeholder="Price in IDR"
      />
      <InputNumberElement
        name="credit"
        label="Credit"
        placeholder="Total credits in unit"
      />
      <MultiSelectElement name="classIds" label="Accessible Classes" />
      <InputNumberElement
        name="expired"
        label="Expiration Time"
        placeholder="Package expiration duration in days"
      />
      <InputNumberElement
        min={0}
        max={99}
        name="discount"
        label="Discount"
        placeholder="Discount in percent"
      />
      <InputTagsElement
        name="additional"
        label="Additional Information"
        placeholder="Enter info, press enter to add"
      />
      <InputFileElement name="image" label="Thumbnail" isSingle />
      <SwitchElement name="isActive" label="Set as active package" />
    </FormUpdateDialog>
  );
};

export { PackageUpdate };
