export const NoClassSchedule = ({ type = "upcoming" }) => {
  const isUpcoming = type === "upcoming";

  return (
    <div className="text-center py-16 px-6 bg-muted/40 border border-dashed border-border rounded-xl space-y-4">
      <div className="flex justify-center">
        <img
          src={"/no-bookings.webp"}
          alt="no-classes-schedule"
          className="h-60 md:h-72 object-contain"
        />
      </div>

      <div className="mb-6 space-y-1">
        <h2 className="text-lg font-semibold text-foreground">
          {isUpcoming
            ? "You don’t have any upcoming classes"
            : "You don’t have any past classes"}
        </h2>
      </div>
    </div>
  );
};
