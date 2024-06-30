import { Input } from "@/components/ui/input";
import React from "react";

type Props = {
  title: string;
  placeholder: string;
};

const SearchBox: React.FC<Props> = ({ title, placeholder }) => {
  return (
    <>
      <h1 className="text-4xl">{title}</h1>
      <Input type="text" className="mt-6 border-2" placeholder={placeholder} />
    </>
  );
};

export default SearchBox;
