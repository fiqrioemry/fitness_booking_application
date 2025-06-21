export const SectionTitle = ({ title, description }) => {
  return (
    <div className="space-y-1 text-center">
      <h2 className="text-2xl font-bold">{title}</h2>
      <p className="text-muted-foreground text-sm">{description}</p>
    </div>
  );
};
