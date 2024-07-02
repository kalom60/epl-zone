import SearchBox from "./SearchBox";
import Goalie2 from "../assets/goalie2.jpg";
import Def3 from "../assets/def3.avif";
import ImageCard from "./ImageCard";

const Positions = () => {
  return (
    <div className="container mt-8">
      <div className="md:mx-48">
        <SearchBox title="Positions" placeholder="Search for positions" />
      </div>

      <div className="w-full relative mt-20">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <ImageCard title="Goalie" image={Goalie2} />
          <ImageCard title="Defender" image={Def3} />
        </div>
      </div>
    </div>
  );
};

export default Positions;
