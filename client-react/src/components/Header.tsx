import { Button } from "@/components/ui/button";
import EPL from "../assets/epl.png";
// import { ModeToggle } from "./mode-toggle";
import { NavLink } from "react-router-dom";
import { Search } from "lucide-react";

const epllogo: string =
  "https://www.premierleague.com/resources/rebrand/v7.149.0/i/elements/pl-main-logo.png";

const Header = () => {
  return (
    <div>
      <div className="hidden md:block">
        <NavLink
          to="/"
          className="absolute top-0 left-0 w-48 flex justify-center z-20 bg-transparent"
        >
          <img
            src={epllogo ? epllogo : EPL}
            alt="epl logo"
            className="object-contain z-[inherit] opacity-100 relative w-36 h-44"
          />
        </NavLink>
      </div>
      <header className="relative bg-[#37003c] top-0 left-0 w-full mt-12">
        <div className="max-w-[1400px] mx-auto w-full flex flex-row-reverse items-center justify-between p-4 py-6">
          <button
            // on:click={() => ($openModal = true)}
            className=" md:hidden grid place-items-center"
          >
            <i className="fa-solid fa-bars"></i>
          </button>
          <nav className=" hidden md:flex items-center gap-4 lg:gap-6 text-white">
            <NavLink
              to="/"
              style={({ isActive }) =>
                isActive
                  ? {
                      color: "#000",
                      background: "#fff",
                      borderRadius: "calc(var(--radius) - 4px)",
                    }
                  : {}
              }
            >
              <Button variant="ghost" className="font-bold">
                Home
              </Button>
            </NavLink>

            <NavLink
              to="/teams"
              style={({ isActive }) =>
                isActive
                  ? {
                      color: "#000",
                      background: "#fff",
                      borderRadius: "calc(var(--radius) - 4px)",
                    }
                  : {}
              }
            >
              <Button variant="ghost" className="font-bold">
                Teams
              </Button>
            </NavLink>

            <NavLink
              to="/nations"
              style={({ isActive }) =>
                isActive
                  ? {
                      color: "#000",
                      background: "#fff",
                      borderRadius: "calc(var(--radius) - 4px)",
                    }
                  : {}
              }
            >
              <Button variant="ghost" className="font-bold">
                Nations
              </Button>
            </NavLink>

            <NavLink
              to="/positions"
              style={({ isActive }) =>
                isActive
                  ? {
                      color: "#000",
                      background: "#fff",
                      borderRadius: "calc(var(--radius) - 4px)",
                    }
                  : {}
              }
            >
              <Button variant="ghost" className="font-bold">
                Positions
              </Button>
            </NavLink>

            <a
              href="#faqs"
              className=" duration-200 hover:text-indigo-400 cursor-pointer"
            >
              <Search />
            </a>
          </nav>

          {/* <div className="hidden md:block mr-6">
          <ModeToggle />
        </div>  */}
        </div>
      </header>
    </div>
  );
};

export default Header;
