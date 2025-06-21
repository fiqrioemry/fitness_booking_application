import { Link } from "react-router-dom";
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/Button";

export const HeroSection = () => {
  const heroImages = [
    "/hero-image1.webp",
    "/hero-image2.webp",
    "/hero-image3.webp",
    "/hero-image4.webp",
  ];

  const [currentImage, setCurrentImage] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentImage((prev) => (prev + 1) % heroImages.length);
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <section
      className="relative w-full h-[800px] bg-cover bg-center transition-all duration-1000 ease-in-out"
      style={{ backgroundImage: `url('${heroImages[currentImage]}')` }}
    >
      <div className="absolute inset-0 bg-black/50 flex items-center justify-center">
        <div className="text-center px-4 max-w-3xl">
          <h1 className="text-4xl md:text-6xl font-bold text-primary/90 mb-4 font-heading">
            Train Better. Live Stronger.
          </h1>
          <p className="text-lg text-muted-foreground mb-6">
            Discover, schedule, and attend classes that energize your lifestyle
            — all in one platform.
          </p>
          <div className="flex flex-col md:flex-row gap-4 justify-center">
            <Link to="/classes">
              <Button className="w-60 h-12" size="lg">
                Book a Class
              </Button>
            </Link>
            <Link to="/packages">
              <Button className="w-60 h-12" variant="outline" size="lg">
                See Our Packages
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
};
