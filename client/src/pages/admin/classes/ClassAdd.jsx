import { classSchema } from "@/lib/schema";
import { classState } from "@/lib/constant";
import { useNavigate } from "react-router-dom";
import { useClassMutation } from "@/hooks/useClass";
import { FormInput } from "@/components/form/FormInput";
import { SectionTitle } from "@/components/header/SectionTitle";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTagsElement } from "@/components/input/InputTagsElement";
import { InputFileElement } from "@/components/input/InputFileElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";
import { SelectOptionsElement } from "@/components/input/SelectOptionsElement";

const ClassAdd = () => {
  const navigate = useNavigate();
  const { createClass } = useClassMutation();

  const handleCreateClass = async (data) => {
    try {
      await createClass.mutateAsync(data);
      navigate("/admin/classes");
    } catch (error) {
      console.error("Failed to create classes:", error);
    }
  };

  return (
    <section className="section">
      <SectionTitle
        title="Add New Class"
        description="Fill out the form below to add a new class to the system."
      />

      <div className="bg-background border shadow-sm rounded-xl p-6">
        <FormInput
          shouldReset
          className="w-72"
          state={classState}
          text="Add New Class"
          schema={classSchema}
          action={handleCreateClass}
          isLoading={createClass.isPending}
        >
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <InputTextElement
                name="title"
                label="Class Title"
                placeholder="Enter class name"
              />
              <InputTextareaElement
                name="description"
                label="Description"
                placeholder="Write description for your class"
                maxLength={200}
              />
            </div>
            <InputFileElement
              isSingle
              name="image"
              label="Class Thumbnail"
              note="Recommended: 1:1 ratio (400x400px)"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <SelectOptionsElement
              name="locationId"
              data="location"
              label="Location"
              placeholder="Select class location"
            />
            <InputNumberElement
              name="duration"
              label="Duration (minutes)"
              placeholder="e.g. 60"
            />
          </div>

          <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
            <SelectOptionsElement
              name="categoryId"
              data="category"
              label="Category"
              placeholder="Select category"
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
              placeholder="Select subcategory"
            />
            <SelectOptionsElement
              name="typeId"
              data="type"
              label="Class Type"
              placeholder="Select type"
            />
          </div>

          <div className="space-y-4">
            <InputTagsElement
              name="additional"
              label="Additional Information"
              placeholder="Enter info, press enter to add"
            />
            <InputFileElement
              name="images"
              label="Gallery (Optional)"
              note="You can upload multiple images"
            />
          </div>
        </FormInput>
      </div>
    </section>
  );
};

export default ClassAdd;
