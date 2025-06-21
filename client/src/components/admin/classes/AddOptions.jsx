import { useMutationOptions } from "@/hooks/useSelectOptions";
import { FormAddDialog } from "@/components/form/FormAddDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { locationState, subcategoryState, optionState } from "@/lib/constant";
import { locationSchema, subcategorySchema, optionSchema } from "@/lib/schema";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

const AddOptions = ({ activeTab }) => {
  const { createOptions } = useMutationOptions(activeTab);

  // schema
  const schema =
    activeTab === "location"
      ? locationSchema
      : activeTab === "subcategory"
      ? subcategorySchema
      : optionSchema;

  // state
  const state =
    activeTab === "location"
      ? locationState
      : activeTab === "subcategory"
      ? subcategoryState
      : optionState;

  return (
    <FormAddDialog
      state={state}
      schema={schema}
      title={`Add New ${activeTab}`}
      buttonText={`New ${activeTab}`}
      loading={createOptions.isPending}
      action={createOptions.mutateAsync}
    >
      <InputTextElement
        name="name"
        label={`Name of ${activeTab}`}
        placeholder={`Enter ${activeTab} name`}
      />
      {activeTab === "location" && <LocationInputElement />}
      {activeTab === "subcategory" && <CategoryInputElement />}
    </FormAddDialog>
  );
};

export { AddOptions };

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
