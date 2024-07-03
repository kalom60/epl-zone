import SearchBox from "./SearchBox";
import Goalie2 from "../assets/goalie2.jpg";
import Def3 from "../assets/def3.avif";
import ImageCard from "./ImageCard";
import { NavLink } from "react-router-dom";

const Positions = () => {
  return (
    <div className="container mt-8">
      <div className="md:mx-48">
        <SearchBox title="Positions" placeholder="Search for positions" />
      </div>

      <div className="w-full relative mt-20">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <NavLink to="/positions/GK">
            <ImageCard title="Goalie" image={Goalie2} />
          </NavLink>

          <NavLink to="/positions/DF">
            <ImageCard title="Defender" image={Def3} />
          </NavLink>
        </div>
      </div>
    </div>
  );
};

export default Positions;
