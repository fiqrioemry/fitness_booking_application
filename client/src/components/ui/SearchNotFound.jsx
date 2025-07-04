export const SearchNotFound = ({ title, q }) => {
  return (
    <div className="py-12 text-center text-gray-500 text-sm">
      {title} {q && ` for "${q}"`}
    </div>
  );
};
