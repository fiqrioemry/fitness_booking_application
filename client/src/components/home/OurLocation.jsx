export const OurLocation = () => {
  return (
    <section className="py-20 px-4 bg-muted border-t border-border">
      <div className="max-w-6xl mx-auto space-y-10">
        <div className="text-center space-y-2">
          <h2 className="text-4xl font-bold text-center mb-4 font-heading">
            Visit Our Studio
          </h2>
          <p className="text-sm text-muted-foreground max-w-xl mx-auto">
            Experience our studio space — drop in, feel the energy, and meet our
            instructors in person.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div className="w-full h-80 rounded-xl overflow-hidden shadow ring-1 ring-border">
            <iframe
              title="Studio Map"
              width="100%"
              height="100%"
              style={{ border: 0 }}
              loading="lazy"
              allowFullScreen
              src="https://www.google.com/maps?q=40.712776,-74.005974&z=15&output=embed"
            ></iframe>
          </div>
          <div className="p-6 space-y-4">
            <h4 className="text-lg font-semibold">Fitness Studio A</h4>
            <p className="text-sm text-muted-foreground">
              <span className="font-medium">Address:</span> 123 Fitness St, New
              York, NY
            </p>
            <p className="text-sm text-muted-foreground">
              <span className="font-medium">Hours:</span> 6AM – 10PM daily
            </p>
            <p className="text-sm text-muted-foreground">
              Join a class, get a tour, or just stop by to see what we’re all
              about!
            </p>
          </div>
        </div>
      </div>
    </section>
  );
};
