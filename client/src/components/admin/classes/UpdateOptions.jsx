import { useMutationOptions } from "@/hooks/useSelectOptions";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { locationSchema, subcategorySchema, optionSchema } from "@/lib/schema";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

export const UpdateOptions = ({ option, activeTab }) => {
  const schema =
    activeTab === "location"
      ? locationSchema
      : activeTab === "subcategory"
      ? subcategorySchema
      : optionSchema;

  const { updateOptions } = useMutationOptions(activeTab);

  return (
    <FormUpdateDialog
      state={option}
      schema={schema}
      title={`Add New ${activeTab}`}
      buttonText={`New ${activeTab}`}
      loading={updateOptions.isPending}
      action={updateOptions.mutateAsync}
    >
      <InputTextElement
        name="name"
        label={`Name of ${activeTab}`}
        placeholder={`Enter ${activeTab} name`}
      />
      {activeTab === "location" && <LocationInputElement />}
      {activeTab === "subcategory" && <CategoryInputElement />}
    </FormUpdateDialog>
  );
};

const CategoryInputElement = () => {
  return (
    <div>
      <SelectOptionsElement
        name="categoryId"
        label="Category"
        placeholder="Select Category options"
      />
    </div>
  );
};

const LocationInputElement = () => {
  return (
    <>
      <div>
        <InputTextareaElement
          name="address"
          label="Address"
          placeholder="Enter full address"
        />
        <InputTextElement
          name="geoLocation"
          label="Geo Location"
          placeholder="Enter geolocation coordinates"
        />
      </div>
    </>
  );
};
