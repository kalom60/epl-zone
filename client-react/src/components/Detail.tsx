import { useParams } from "react-router-dom";
import { useCallback, useEffect, useState } from "react";
import { Data } from "../utils/columns";
import { DataTable } from "@/utils/table";

const Detail = () => {
  const { name } = useParams();
  const path: string = window.location.pathname.split("/")[1];
  const [details, setDetails] = useState<Data[]>([]);

  const fetchDetails = useCallback(
    async (p: string) => {
      const response = await fetch(`http://localhost:8080/player/${p}/${name}`);
      const data = await response.json();
      setDetails(data);
    },
    [name],
  );

  useEffect(() => {
    if (path === "teams") {
      fetchDetails("team");
    } else if (path === "positions") {
      fetchDetails("position");
    }
  }, [path, fetchDetails]);

  return (
    <div className="container mt-14">
      <DataTable data={details} />
    </div>
  );
};

export default Detail;
