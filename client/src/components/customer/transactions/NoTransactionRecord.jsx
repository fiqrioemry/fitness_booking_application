export const NoTransactionRecord = () => {
  return (
    <div className="p-10 text-center space-y-4">
      <div className="flex justify-center">
        <img
          src="/no-transactions.webp"
          alt="no-transactions"
          className="h-72"
        />
      </div>

      <h3 className="text-lg font-semibold text-foreground">
        No Transactions Found
      </h3>

      <p className="text-subtitle max-w-md mx-auto">
        You haven't made any transactions yet. Start exploring our packages and
        enjoy various training sessions today!
      </p>

      <a href="/packages" className="btn btn-primary">
        Browse Packages
      </a>
    </div>
  );
};
