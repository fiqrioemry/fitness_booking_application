import { toast } from "sonner";
import { useState } from "react";
import { PlusCircle, X } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { Controller, useFormContext } from "react-hook-form";

const InputFileElement = ({
  name,
  label,
  maxImages = 5,
  maxSizeMB = 2,
  isSingle = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();
  const [isDragging, setIsDragging] = useState(false);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleFiles = (files) => {
          const validFiles = Array.from(files).filter((file) => {
            const isValidSize = file.size / (1024 * 1024) <= maxSizeMB;
            if (!isValidSize) {
              toast.warning(`${file.name} exceeds ${maxSizeMB}MB`);
            }
            return isValidSize;
          });

          if (validFiles.length === 0) return;

          if (isSingle) {
            field.onChange(validFiles[0]);
          } else {
            const updated = [
              ...(field.value || []).filter(
                (f) => f instanceof File || typeof f === "string"
              ),
              ...validFiles,
            ].slice(0, maxImages);
            field.onChange(updated);
          }
        };

        const handleRemoveImage = (img) => {
          if (isSingle) {
            field.onChange(null);
          } else {
            const updated = (field.value || []).filter((file) => file !== img);
            field.onChange(updated);
          }
        };

        const handleDrop = (e) => {
          e.preventDefault();
          setIsDragging(false);
          if (e.dataTransfer.files.length > 0) {
            handleFiles(e.dataTransfer.files);
            e.dataTransfer.clearData();
          }
        };

        const getImageURL = (item) => {
          if (item instanceof File) return URL.createObjectURL(item);
          if (typeof item === "string") return item;
          return "";
        };

        const renderSinglePreview = () => {
          if (!field.value) return null;
          const url = getImageURL(field.value);
          return (
            <div className="relative w-full h-full">
              <img
                src={url}
                alt="preview"
                className="object-cover w-full h-full rounded-md"
              />
              <Button
                type="button"
                size="icon"
                variant="destructive"
                onClick={() => handleRemoveImage(field.value)}
                className="absolute top-1 right-1 p-1"
              >
                <X className="w-4 h-4" />
              </Button>
            </div>
          );
        };

        const renderMultiplePreview = () => {
          return (field.value || []).map((img, idx) => (
            <div
              key={idx}
              className="relative w-32 h-32 border border-border rounded-md overflow-hidden"
            >
              <img
                src={getImageURL(img)}
                alt="preview"
                className="object-cover w-full h-full"
              />
              <Button
                size="icon"
                type="button"
                variant="destructive"
                onClick={() => handleRemoveImage(img)}
                className="absolute top-1 right-1 p-1"
              >
                <X className="w-2 h-2" />
              </Button>
            </div>
          ));
        };

        return (
          <div className="space-y-2">
            {label && (
              <label
                htmlFor={name}
                className="block text-sm font-medium text-foreground"
              >
                {label}
              </label>
            )}

            <div
              onDrop={handleDrop}
              onDragOver={(e) => {
                e.preventDefault();
                setIsDragging(true);
              }}
              onDragLeave={(e) => {
                e.preventDefault();
                setIsDragging(false);
              }}
              className={`${
                isSingle
                  ? "relative w-full h-64 flex items-center justify-center overflow-hidden"
                  : "flex flex-wrap gap-4 p-4"
              } border-2 rounded-md transition ${
                isDragging
                  ? "border-primary bg-primary/10"
                  : "border-border bg-muted/30"
              }`}
            >
              {isSingle ? (
                field.value ? (
                  renderSinglePreview()
                ) : (
                  <label
                    htmlFor={`${name}-upload`}
                    className="flex flex-col items-center justify-center w-full h-full cursor-pointer hover:bg-muted transition rounded-md"
                  >
                    <PlusCircle className="text-primary mb-2" />
                    <span className="text-sm ">Select Image</span>
                  </label>
                )
              ) : (
                <>
                  {renderMultiplePreview()}
                  {(!field.value || field.value.length < maxImages) && (
                    <label
                      htmlFor={`${name}-upload`}
                      className="flex flex-col items-center justify-center w-32 h-32 border-2 border-dashed border-primary rounded-md cursor-pointer hover:bg-muted transition"
                    >
                      <PlusCircle className="text-primary mb-2" />
                      <span className="text-sm">Select Images</span>
                    </label>
                  )}
                </>
              )}
              <input
                id={`${name}-upload`}
                type="file"
                accept="image/*"
                onChange={(e) => {
                  handleFiles(e.target.files);
                  e.target.value = "";
                }}
                multiple={!isSingle}
                hidden
              />
            </div>

            {fieldState.error && (
              <p className="text-destructive text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputFileElement };
