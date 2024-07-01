import { useEffect, useState } from "react";
import SearchBox from "./SearchBox";

type State = {
  id: string;
  team: string;
  logo: string;
};

const Teams = () => {
  const [teams, setTeams] = useState<State[]>([]);

  const getTeams = async () => {
    const respose = await fetch("http://localhost:8080/teams");
    const data = await respose.json();
    setTeams(data);
  };

  useEffect(() => {
    getTeams();
  }, []);

  return (
    <div className="container mt-8">
      <div className="md:mx-48">
        <SearchBox title="Teams" placeholder="Search for teams" />
      </div>

      <div className="grid grid-cols-4 gap-4 place-items-center mt-20">
        {teams.length > 0 &&
          teams.map((team: State) => {
            const arr: string[] = team.logo.split("mini.");
            const logo = arr.join("");

            return (
              <div className="my-4">
                <img src={logo} />
                <h1>{team.team}</h1>
              </div>
            );
          })}
      </div>
    </div>
  );
};

export default Teams;
