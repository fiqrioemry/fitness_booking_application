import { Input } from "@/components/ui/Input";

export const SearchInput = ({ q, setPage, setQ, placeholder }) => {
  return (
    <Input
      value={q}
      className="md:w-1/2"
      onChange={(e) => {
        setPage(1);
        setQ(e.target.value);
      }}
      placeholder={placeholder}
    />
  );
};
