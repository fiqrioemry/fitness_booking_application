import { toast } from "sonner";
import { loadStripe } from "@stripe/stripe-js";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "@/store/useAuthStore";
import { useCreatePaymentMutation } from "./usePayment";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY);

export function useStripePayment(packageId, voucherCode) {
  const { user } = useAuthStore();
  const navigate = useNavigate();
  const { mutate: createPayment, isPending } = useCreatePaymentMutation();

  const handleBuyNow = () => {
    if (!user) return navigate("/signin");

    const payload = {
      packageId,
      voucherCode,
    };

    createPayment(payload, {
      onSuccess: async (res) => {
        if (!res.sessionId) {
          toast.error("No session ID received.");
          return;
        }

        const stripe = await stripePromise;
        if (!stripe) {
          toast.error("Stripe SDK not loaded.");
          return;
        }

        const result = await stripe.redirectToCheckout({
          sessionId: res.sessionId,
        });

        if (result.error) {
          toast.error(result.error.message || "Failed to redirect to Stripe.");
        }
      },
      onError: () => {
        toast.error("Failed to create Stripe session.");
      },
    });
  };

  return {
    handleBuyNow,
    isPending,
  };
}
