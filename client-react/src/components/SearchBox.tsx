import { Input } from "@/components/ui/input";
import React, { ChangeEvent } from "react";

type Props = {
  title: string;
  placeholder: string;
  onSearch: (query: string) => void;
};

const SearchBox: React.FC<Props> = ({ title, placeholder, onSearch }) => {
  const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
    onSearch(e.target.value);
  };

  return (
    <>
      <h1 className="text-4xl font-semibold">{title}</h1>
      <Input
        type="text"
        className="mt-6 border-2 p-5"
        placeholder={placeholder}
        onChange={handleInputChange}
      />
    </>
  );
};

export default SearchBox;
