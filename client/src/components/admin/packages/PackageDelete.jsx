/* eslint-disable react/prop-types */
import { usePackageMutation } from "@/hooks/usePackage";
import { FormDelete } from "@/components/form/FormDelete";

const PackageDelete = ({ pkg }) => {
  const { deletePackage } = usePackageMutation();

  const handleDeletePackage = () => {
    deletePackage.mutate(pkg.id);
  };
  return (
    <FormDelete
      title="Delete Package"
      loading={deletePackage.isPending}
      onDelete={handleDeletePackage}
      description="Are you sure want to delete this package ?"
    />
  );
};

export { PackageDelete };
