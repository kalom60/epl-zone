import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { Data } from "../utils/columns";
import { DataTable } from "@/utils/table";

const TeamDetail = () => {
  const { name } = useParams();
  const [details, setDetails] = useState<Data[]>([]);

  const fetchDetails = async () => {
    const response = await fetch(`http://localhost:8080/player/team/${name}`);
    const data = await response.json();
    setDetails(data);
  };

  useEffect(() => {
    fetchDetails();
  }, []);

  return (
    <div className="container mt-14">
      <DataTable data={details} />
    </div>
  );
};

export default TeamDetail;
