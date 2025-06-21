// import * as auth from "@/services/auth";
// import { useAuthStore } from "@/store/useAuthStore";
// import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
// import { toast } from "sonner";
// import { useAuthStepStore } from "../store/useAuthStore";

// // GET: /auth/me
// export const useAuthQuery = () => {
//   const setUser = useAuthStore((s) => s.setUser);

//   return useQuery({
//     queryKey: ["auth", "me"],
//     queryFn: auth.getMe,
//     select: (data) => {
//       setUser(data);
//       return data;
//     },
//     onError: () => setUser(null),
//     retry: 1,
//     staleTime: 1000 * 60 * 30, // 30 minutes
//   });
// };

// export const useLoginMutation = () => {
//   const queryClient = useQueryClient();
//   const setRememberMe = useAuthStore((s) => s.setRememberMe);

//   return useMutation({
//     mutationFn: auth.login, // expects formData { email, password, rememberMe }
//     onSuccess: (_, variables) => {
//       toast.success("Login successful");

//       if (variables.rememberMe) {
//         setRememberMe(true);
//       } else {
//         setRememberMe(false);
//       }
//       queryClient.invalidateQueries({ queryKey: ["auth", "me"] });
//     },
//     onError: (error) => {
//       toast.error(error?.response?.data?.message || "Login failed");
//     },
//   });
// };

// // POST: /auth/send-otp
// export const useSendOTPMutation = () => {
//   const { setStep } = useAuthStepStore();

//   return useMutation({
//     mutationFn: auth.sendOTP,
//     onSuccess: (_, variables) => {
//       toast.success("OTP sent to email.");
//       setStep(2);
//     },
//     onError: (err) => {
//       toast.error(err?.response?.data?.message || "Failed to send OTP.");
//     },
//   });
// };

// // POST: /auth/verify-otp
// export const useVerifyOTPMutation = () => {
//   const { setStep } = useAuthStepStore();
//   return useMutation({
//     mutationFn: auth.verifyOTP,
//     onSuccess: () => {
//       toast.success("OTP verified.");
//       setStep(3);
//     },
//     onError: (err) => {
//       toast.error(err?.response?.data?.message || "OTP verification failed.");
//     },
//     u,
//   });
// };

// export const useLogout = () => {
//   const clearUser = useAuthStore((s) => s.clearUser);
//   const queryClient = useQueryClient();

//   return useMutation({
//     mutationFn: auth.logout,
//     onSuccess: () => {
//       clearUser();
//       queryClient.clear();
//       toast.success("Logged out successfully.");
//     },
//     onError: (error) => {
//       console.error(error?.message);
//     },
//   });
// };

// // POST: /auth/register
// export const useRegisterMutation = (email) => {
//   const { resetStep } = useAuthStepStore();
//   const queryClient = useQueryClient();

//   return useMutation({
//     mutationFn: async (data) => {
//       return await auth.register({ ...data, email });
//     },
//     onSuccess: () => {
//       toast.success("Registration successful!");
//       queryClient.invalidateQueries({ queryKey: ["auth", "me"] });
//       resetStep();
//     },
//     onError: (err) => {
//       toast.error(err?.response?.data?.message || "Registration failed.");
//     },
//   });
// };
