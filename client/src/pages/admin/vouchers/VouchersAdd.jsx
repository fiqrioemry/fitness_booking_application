import { useNavigate } from "react-router-dom";
import { createVoucherSchema } from "@/lib/schema";
import { createVoucherState } from "@/lib/constant";
import { FormInput } from "@/components/form/FormInput";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SwitchElement } from "@/components/input/SwitchElement";
import { SelectElement } from "@/components/input/SelectElement";
import { InputDateElement } from "@/components/input/InputDateElement";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputNumberElement } from "@/components/input/InputNumberElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const discountTypeOptions = [
  { label: "Fixed", value: "fixed" },
  { label: "Percentage", value: "percentage" },
];

const VoucherAdd = () => {
  const navigate = useNavigate();
  const { createVoucher } = useVoucherMutation();

  const handleCreateVoucher = async (data) => {
    try {
      const payload = {
        ...data,
        expiredAt: data.expiredAt
          ? new Date(data.expiredAt).toISOString()
          : null,
      };
      await createVoucher.mutateAsync(payload);
      navigate("/admin/vouchers");
    } catch (error) {
      console.error("Failed to create voucher:", error);
    }
  };

  return (
    <section className="section">
      <SectionTitle
        title="Create New Voucher"
        description="Fill out the form below to add a new voucher."
      />

      <div className="bg-background border shadow-sm rounded-xl p-6">
        <FormInput
          text="Create Voucher"
          className="w-72"
          state={createVoucherState}
          schema={createVoucherSchema}
          action={handleCreateVoucher}
          isLoading={createVoucher.isPending}
        >
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <InputTextElement
              name="code"
              label="Voucher Code"
              placeholder="e.g. HEALTHY100K"
            />
            <SelectElement
              name="discountType"
              label="Discount Type"
              options={discountTypeOptions}
              placeholder="Select type"
            />
          </div>

          <InputTextareaElement
            name="description"
            label="Description"
            placeholder="e.g. Diskon 100 ribu untuk semua paket"
            maxLength={200}
          />

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            <InputNumberElement
              name="discount"
              label="Discount Amount"
              placeholder="e.g. 100000 or 50"
            />
            <InputNumberElement
              name="maxDiscount"
              label="Max Discount (if %)"
              placeholder="e.g. 30000"
              isOptional
            />
            <InputNumberElement
              name="quota"
              label="Quota"
              placeholder="e.g. 10"
            />
          </div>

          <InputDateElement
            name="expiredAt"
            mode="future"
            label="Expiration Date"
            placeholder="YYYY-MM-DD"
          />
          <SwitchElement name="isReusable" label="Allow multiple usage?" />
        </FormInput>
      </div>
    </section>
  );
};

export default VoucherAdd;
