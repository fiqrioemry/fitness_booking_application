import { createVoucherSchema } from "@/lib/schema";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { SwitchElement } from "@/components/input/SwitchElement";
import { SelectElement } from "@/components/input/SelectElement";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const discountTypeOptions = [
  { label: "Fixed", value: "fixed" },
  { label: "Percentage", value: "percentage" },
];

const VoucherUpdate = ({ voucher }) => {
  const { updateVoucher } = useVoucherMutation();

  const handleUpdateVoucher = ({ id, data }) => {
    const payload = {
      ...data,
      expiredAt: data.expiredAt ? new Date(data.expiredAt).toISOString() : null,
    };
    updateVoucher.mutateAsync({ id, data: payload });
  };

  return (
    <FormUpdateDialog
      state={voucher}
      title="Update Voucher"
      schema={createVoucherSchema}
      action={handleUpdateVoucher}
      loading={updateVoucher.isPending}
    >
      <InputTextElement
        name="code"
        label="Voucher Code (cannot changed)"
        disabled
      />
      <InputTextareaElement
        maxLength={200}
        name="description"
        label="Description"
        placeholder="e.g. Diskon 50% for all classes"
      />
      <SelectElement
        name="discountType"
        label="Discount Type"
        options={discountTypeOptions}
        placeholder="Select discount type"
      />
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <InputNumberElement
          name="discount"
          label="Discount"
          placeholder="e.g. 50000 or 50"
        />
        <InputNumberElement
          name="maxDiscount"
          label="Max Discount (if %)"
          placeholder="e.g. 30000 (Rupiah)"
        />
        <InputNumberElement name="quota" label="Quota" placeholder="e.g. 10" />
      </div>
      <InputDateElement
        mode="future"
        name="expiredAt"
        label="Expiration Date"
        placeholder="YYYY-MM-DD"
      />
      <SwitchElement name="isReusable" label="Allow multiple usage?" />
    </FormUpdateDialog>
  );
};

export { VoucherUpdate };
