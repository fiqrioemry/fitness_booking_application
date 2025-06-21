import { Link } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { useClassesQuery } from "@/hooks/useClass";
import { OurClassSkeleton } from "@/components/loading/OurClassSkeleton";

export const OurClasses = () => {
  const { data: response, isLoading } = useClassesQuery({ limit: 3 });

  const classes = response?.classes || [];

  return (
    <section className="py-20 px-4 bg-muted">
      <h2 className="text-4xl font-bold text-center mb-4 font-heading">
        Our Classes
      </h2>
      <p className="text-center text-muted-foreground mb-12 max-w-xl mx-auto text-base">
        From relaxing yoga sessions to high-intensity workouts, choose the class
        that suits your fitness goals.
      </p>
      {isLoading ? (
        <OurClassSkeleton />
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto">
          {classes.map((classItem) => (
            <div
              key={classItem.id}
              className="bg-card text-foreground border border-border rounded-xl shadow hover:shadow-xl transition overflow-hidden"
            >
              <img
                src={classItem.image}
                alt={classItem.title}
                className="w-full h-56 object-cover"
              />
              <div className="p-6 space-y-3">
                <h4 className="text-lg font-bold">{classItem.title}</h4>
                <p className="text-sm text-muted-foreground line-clamp-2">
                  {classItem.description}
                </p>
                <div className="flex justify-between items-center pt-2">
                  <span className="text-sm font-medium text-primary">
                    {classItem.duration} min
                  </span>
                  <Link to="/schedules">
                    <Button size="sm">Book Now</Button>
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </section>
  );
};
