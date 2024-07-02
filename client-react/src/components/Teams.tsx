import { useEffect, useState } from "react";
import SearchBox from "./SearchBox";
import { motion, AnimatePresence } from "framer-motion";
import { NavLink } from "react-router-dom";
import { Button } from "./ui/button";

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
              <AnimatePresence mode="wait" key={team.id}>
                <motion.div
                  className="my-10"
                  initial={{ y: 100, opacity: 0 }}
                  animate={{ y: 0, opacity: 1 }}
                  exit={{ y: -10, opacity: 0 }}
                  transition={{ duration: 0.2 }}
                  whileHover={{ scale: 1.2 }}
                  whileTap={{ scale: 0.8 }}
                >
                  <NavLink to={`/teams/${clubUrl}`}>
                    <img src={logo} className="mx-auto w-48 h-48" />
                    <h1 className="text-center mt-6 text-1xl font-medium">
                      <Button className="relative inline-block font-medium group py-1.5 px-2.5 ">
                        <span className="absolute inset-0 w-full h-full transition duration-400 ease-out transform translate-x-1 translate-y-1 bg-[#37003c] group-hover:-translate-x-0 group-hover:-translate-y-0"></span>
                        <span className="absolute inset-0 w-full h-full bg-white border border-[#37003c] group-hover:bg-indigo-50"></span>
                        <span className="relative text-[#37003c]">{club}</span>
                      </Button>
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
