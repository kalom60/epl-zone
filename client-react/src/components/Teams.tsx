import { useEffect, useState } from "react";
import SearchBox from "./SearchBox";
import { motion, AnimatePresence } from "framer-motion";
import { NavLink } from "react-router-dom";

type State = {
  id: string;
  team: string;
  logo: string;
};

const Teams = () => {
  const [teams, setTeams] = useState<State[]>([]);
  const [searchQuery, setSearchQuery] = useState<string>("");

  const filteredTeams = teams.filter((team) =>
    team.team.toLowerCase().includes(searchQuery.toLowerCase()),
  );

  const getTeams = async () => {
    const respose = await fetch("http://localhost:8080/teams");
    const data = await respose.json();
    setTeams(data);
  };

  useEffect(() => {
    getTeams();
  }, []);

  return (
    <div className="container mt-14">
      <div className="md:mx-48">
        <SearchBox
          title="Teams"
          placeholder="Search for teams"
          onSearch={(query) => setSearchQuery(query)}
        />
      </div>

      <div className="grid sm:grid-cols-2 md:grid-cols-4 gap-4 place-items-center mt-20">
        {filteredTeams.length > 0 &&
          filteredTeams.map((team: State) => {
            const arr: string[] = team.logo.split("mini.");
            const logo: string = arr.join("");
            const teamName: string[] = team.team.split("-");
            const club: string = teamName.slice(0, -1).join(" ");
            const clubUrl: string = teamName.slice(0, -1).join("-");

            return (
              <AnimatePresence mode="wait">
                <motion.div
                  className="my-10"
                  key={team.id}
                  initial={{ y: 100, opacity: 0 }}
                  animate={{ y: 0, opacity: 1 }}
                  exit={{ y: -10, opacity: 0 }}
                  transition={{ duration: 0.2 }}
                  whileHover={{ scale: 1.2 }}
                  whileTap={{ scale: 0.8 }}
                >
                  <NavLink to={`/team/${clubUrl}`}>
                    <img src={logo} className="mx-auto w-48 h-48" />
                    <h1 className="text-center mt-3 text-1xl font-medium">
                      {club}
                    </h1>
                  </NavLink>
                </motion.div>
              </AnimatePresence>
            );
          })}
      </div>
    </div>
  );
};

export default Teams;
