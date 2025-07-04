import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import { WebLogo } from "@/components/ui/WebLogo";
import { useAuthStore } from "@/store/useAuthStore";
import { FormInput } from "@/components/form/FormInput";
import { GoogleOAuth } from "@/components/public/GoogleOAuth";
import { InputTextElement } from "@/components/input/InputTextElement";
import { sendOTPState, verifyOTPState, registerState } from "@/lib/constant";
import { registerSchema, sendOTPSchema, verifyOTPSchema } from "@/lib/schema";

const SignUp = () => {
  const [countdown, setCountdown] = useState(60);
  const [sentEmail, setSentEmail] = useState("");
  const [canResend, setCanResend] = useState(false);
  const { register, step, loading, resetStep, sendOTP } = useAuthStore();

  const getSchemaControl = () => {
    switch (step) {
      case 1:
        return sendOTPSchema;
      case 2:
        return verifyOTPSchema;
      case 3:
        return registerSchema;
      default:
        return [];
    }
  };

  const handleResendOTP = () => {
    if (!canResend || !sentEmail) return;
    sendOTP({ email: sentEmail });
    setCountdown(60);
    setCanResend(false);
  };

  useEffect(() => {
    resetStep();
  }, []);

  useEffect(() => {
    let timer;
    if (step === 2 && !canResend) {
      timer = setInterval(() => {
        setCountdown((prev) => {
          if (prev <= 1) {
            clearInterval(timer);
            setCanResend(true);
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
    }
    return () => clearInterval(timer);
  }, [step, canResend]);

  return (
    <section className="min-h-screen flex items-center justify-center bg-background px-4">
      <div className="w-full max-w-4xl grid grid-cols-1 md:grid-cols-2 bg-card border border-border rounded-xl shadow-lg overflow-hidden">
        <div className="hidden md:block relative h-[550px]">
          <img
            src="/register.png"
            alt="Wallpaper"
            className="absolute inset-0 w-full h-full object-cover"
          />
        </div>

        <div className="px-6 py-10">
          <div className="mb-6 flex justify-center text-center">
            <WebLogo />
          </div>

          <div className="flex justify-center mb-6 gap-2">
            {[1, 2, 3].map((s) => (
              <div
                key={s}
                className={`w-8 h-8 rounded-full flex items-center justify-center text-sm font-semibold ${
                  step === s
                    ? "bg-primary text-primary-foreground"
                    : "bg-muted text-muted-foreground"
                }`}
              >
                {s}
              </div>
            ))}
          </div>

          <FormInput
            text={
              step === 1
                ? "Register with Email"
                : step === 2
                ? "Submit OTP"
                : "Register"
            }
            action={(data) => {
              if (step === 1) setSentEmail(data.email);
              register(data);
            }}
            className="w-full py-3"
            isLoading={loading}
            state={
              step === 1
                ? sendOTPState
                : step === 2
                ? verifyOTPState
                : registerState
            }
            schema={getSchemaControl()}
          >
            {step === 1 && (
              <InputTextElement name="email" placeholder="Enter your email" />
            )}

            {step === 2 && (
              <>
                <div className="mb-2 text-sm text-muted-foreground">
                  We sent an OTP to:{" "}
                  <span className="font-medium text-foreground">
                    {sentEmail}
                  </span>
                </div>
                <InputTextElement
                  name="otp"
                  isNumeric
                  maxLength={6}
                  label="OTP Code"
                  placeholder="Enter the OTP code"
                />
                <div className="text-sm text-right mt-2">
                  {canResend ? (
                    <button
                      type="button"
                      onClick={handleResendOTP}
                      className="text-primary hover:underline"
                    >
                      Resend OTP
                    </button>
                  ) : (
                    <span className="text-muted-foreground">
                      Resend in {countdown}s
                    </span>
                  )}
                </div>
              </>
            )}

            {step === 3 && (
              <>
                <InputTextElement name="email" label="Email" disabled />
                <InputTextElement
                  name="password"
                  type="password"
                  label="Password"
                  placeholder="********"
                />
                <InputTextElement
                  name="fullname"
                  label="Fullname"
                  placeholder="Enter your fullname"
                />
              </>
            )}
          </FormInput>

          {step === 1 && (
            <>
              <div className="text-center py-2 text-sm text-muted-foreground">
                Or
              </div>

              <GoogleOAuth buttonText="Sign up with google" />
              <p className="text-sm text-center mt-6 text-muted-foreground">
                Already have an account?{" "}
                <Link
                  to="/signin"
                  className="text-primary font-medium hover:underline"
                >
                  Login
                </Link>
              </p>
            </>
          )}
        </div>
      </div>
    </section>
  );
};

export default SignUp;
