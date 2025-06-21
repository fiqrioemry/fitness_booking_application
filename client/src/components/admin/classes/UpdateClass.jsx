import { classSchema } from "@/lib/schema";
import { useClassMutation } from "@/hooks/useClass";
import { SwitchElement } from "@/components/input/SwitchElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

const UpdateClass = ({ classes }) => {
  const { updateClass } = useClassMutation();
  const { isPending, mutateAsync } = updateClass;

  return (
    <FormUpdateDialog
      state={classes}
      title="Update Class"
      loading={isPending}
      schema={classSchema}
      action={mutateAsync}
    >
      <InputTextElement
        name="title"
        label="Class Title"
        placeholder="Enter class title"
      />
      <div className="grid grid-cols-2 gap-4">
        <InputFileElement name="image" label="Class Thumbnail" isSingle />
        <div>
          <SelectOptionsElement
            label="Location"
            data="location"
            name="locationId"
            placeholder="Select class location"
          />
          <InputNumberElement
            name="duration"
            label="Duration"
            placeholder="Class duration in minutes"
          />

          <InputTextareaElement
            maxLength={500}
            name="description"
            label="Class Description"
            placeholder="Enter class description (min. 20 characters)"
          />
        </div>
      </div>
      <div className="grid grid-cols-2 gap-4">
        <SelectOptionsElement
          name="categoryId"
          data="category"
          label="Category"
          placeholder="Select class category"
        />
        <SelectOptionsElement
          name="levelId"
          data="level"
          label="Level"
          placeholder="Select difficulty level"
        />
        <SelectOptionsElement
          name="subcategoryId"
          data="subcategory"
          label="Subcategory"
          placeholder="Select class subcategory"
        />
        <SelectOptionsElement
          data="type"
          name="typeId"
          label="Class Type"
          placeholder="Select class type"
        />
      </div>
      <InputTagsElement
        name="additional"
        label="Additional Information"
        placeholder="Enter info, press enter to add"
      />
      <SwitchElement name="isActive" label="Set as active class" />
    </FormUpdateDialog>
  );
};

export { UpdateClass };
