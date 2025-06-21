import { homeTitle } from "@/lib/constant";
import { OurClasses } from "@/components/home/OurClasses";
import { HeroSection } from "@/components/home/HeroSection";
import { OurLocation } from "@/components/home/OurLocation";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";
import { OurPackages } from "@/components/home/OurPackages";

export default function Home() {
  useDocumentTitle(homeTitle);
  return (
    <main className="bg-background text-foreground">
      {/* hero section */}
      <HeroSection />

      {/* our classes */}
      <OurClasses />

      {/* our packages*/}
      <OurPackages />

      {/* Visit Studio */}
      <OurLocation />
    </main>
  );
}
