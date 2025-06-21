import {
  Card,
  CardHeader,
  CardTitle,
  CardContent,
  CardDescription,
} from "@/components/ui/Card";
import { Link } from "react-router-dom";
import { aboutTitle } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useInstructorsQuery } from "@/hooks/useInstructor";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";
import { HeartPulse, CalendarClock, Users } from "lucide-react";
import { AboutSkeleton } from "@/components/loading/AboutSkeleton";

const About = () => {
  useDocumentTitle(aboutTitle);
  const {
    data: instructors = [],
    isLoading,
    isError,
    refetch,
  } = useInstructorsQuery();

  if (isLoading) return <AboutSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section py-24 space-y-20 text-foreground">
      {/* Hero Section */}
      <div className="bg-primary text-primary-foreground rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <h3 className="text-3xl font-bold">Empowering Your Wellness Journey</h3>
        <p className="text-sm opacity-80 max-w-xl mx-auto">
          At Sweat Up, we believe that fitness is not just a goal — it's a
          lifestyle. Join our community and book classes tailored to your body
          and mind.
        </p>
      </div>

      {/* Our Mission */}
      <div className="text-center space-y-4">
        <h2 className="text-2xl font-bold">Our Mission</h2>
        <p className="text-muted-foreground max-w-xl mx-auto text-sm">
          To provide accessible, flexible, and enjoyable wellness classes that
          empower everyone to stay active, reduce stress, and live healthier.
        </p>
      </div>

      {/* Why Choose Us */}
      <div className="space-y-6">
        <h3 className="text-xl font-semibold text-center">
          Why Choose FitBook Studio?
        </h3>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 text-center">
          <Card>
            <CardHeader className="flex flex-col items-center space-y-2">
              <HeartPulse className="text-primary w-8 h-8" />
              <CardTitle className="text-base">Holistic Wellness</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                We focus on your physical, mental, and emotional well-being
                through diverse classes.
              </CardDescription>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="flex flex-col items-center space-y-2">
              <CalendarClock className="text-primary w-8 h-8" />
              <CardTitle className="text-base">Flexible Scheduling</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Book anytime, anywhere. Our platform is designed to fit your
                lifestyle.
              </CardDescription>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="flex flex-col items-center space-y-2">
              <Users className="text-primary w-8 h-8" />
              <CardTitle className="text-base">Expert Instructors</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                All classes are guided by certified instructors who care about
                your progress.
              </CardDescription>
            </CardContent>
          </Card>
        </div>
      </div>

      {/* Instructor Profiles */}
      <div className="space-y-6">
        <h3 className="text-xl font-semibold text-center">
          Meet Our Instructors
        </h3>
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-6">
          {instructors.map((ins) => (
            <Card
              key={ins.id}
              className="p-4 flex flex-col items-center text-center"
            >
              <img
                src={ins.avatar}
                alt={ins.fullname}
                className="w-20 h-20 rounded-full border object-cover mb-3"
              />
              <CardTitle className="text-sm">{ins.fullname}</CardTitle>
              <CardDescription className="text-xs">
                {ins.specialties}
              </CardDescription>
            </Card>
          ))}
        </div>
      </div>

      {/* CTA Section */}
      <div className="text-center bg-muted rounded-xl p-10">
        <h3 className="text-xl font-semibold mb-2">Ready to feel your best?</h3>
        <p className="text-sm text-muted-foreground mb-4">
          Start your fitness journey with us today.
        </p>
        <Link to="/classes">
          <Button size="lg">Explore Classes</Button>
        </Link>
      </div>
    </section>
  );
};

export default About;
