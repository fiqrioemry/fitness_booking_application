import { Link } from "react-router-dom";
import { Button } from "@/components/ui/Button";

const NotFound = () => {
  return (
    <section className="min-h-screen flex items-center justify-center bg-background px-4 py-12">
      <div className="text-center max-w-xl w-full">
        <img
          src="/page-not-found.webp"
          alt="page-not-found-illustration"
          className="mx-auto max-w-xs md:max-w-md mb-8 rounded-md"
        />
        <p className="text-gray-500 text-base md:text-lg mb-6">
          Oops! The page you're looking for doesn't exist or has been moved.
        </p>

        <Link to="/">
          <Button>← Back to Homepage</Button>
        </Link>
      </div>
    </section>
  );
};

export default NotFound;
