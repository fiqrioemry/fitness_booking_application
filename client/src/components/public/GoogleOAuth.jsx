import { FcGoogle } from "react-icons/fc";
import { Button } from "@/components/ui/Button";

export const GoogleOAuth = ({ buttonText }) => {
  const handleGoogleAuth = () => {
    window.location.href = `${import.meta.env.VITE_API_SERVICES}/auth/google`;
  };

  return (
    <div>
      <Button variant="outline" onClick={handleGoogleAuth} className="w-full">
        <FcGoogle size={24} />
        {buttonText}
      </Button>
    </div>
  );
};
