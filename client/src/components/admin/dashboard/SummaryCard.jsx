const SummaryCard = ({ title, value }) => (
  <div className="bg-background rounded-xl p-4 shadow text-center">
    <h3 className="text-sm font-medium text-muted-foreground">{title}</h3>
    <p className="text-2xl font-bold text-primary mt-2">{value}</p>
  </div>
);

export { SummaryCard };
