import { Button } from "@/components/ui/button";
import EPL from "../assets/epl.png";
import { ModeToggle } from "./mode-toggle";
import { NavLink } from "react-router-dom";

const Header = () => {
  return (
    <header className="flex flex-col relative z-20">
      <div className="max-w-[1400px] mx-auto w-full flex items-center justify-between p-4 py-6">
        <NavLink to="/">
          <img src={EPL} alt="epl logo" className="w-[50px] h-[50px] ml-6" />
        </NavLink>
        <button
          // on:click={() => ($openModal = true)}
          className=" md:hidden grid place-items-center"
        >
          <i className="fa-solid fa-bars"></i>
        </button>
        <nav className=" hidden md:flex items-center gap-4 lg:gap-6">
          <NavLink to="/">
            <Button variant="ghost">Home</Button>
          </NavLink>

          <NavLink to="/teams">
            <Button variant="ghost">Teams</Button>
          </NavLink>

          <NavLink to="/nations">
            <Button variant="ghost">Nations</Button>
          </NavLink>

          <NavLink to="/positions">
            <Button variant="ghost">Positions</Button>
          </NavLink>

          <a
            href="#faqs"
            className=" duration-200 hover:text-indigo-400 cursor-pointer"
          >
            <i className="fa fa-search px-2"></i>
          </a>
        </nav>

        {/* <div className="hidden md:block mr-6">
          <ModeToggle />
        </div>  */}
      </div>
    </header>
  );
};

export default Header;
