type Props = {
  title: string;
  image: string;
};

const ImageCard: React.FC<Props> = ({ title, image }) => {
  return (
    <div className="overflow-hidden  aspect-video bg-red-400 cursor-pointer rounded-xl relative group">
      <div className="rounded-xl z-50 opacity-0 group-hover:opacity-100 transition duration-300 ease-in-out cursor-pointer absolute from-black/80 to-transparent bg-gradient-to-t inset-x-0 -bottom-2 pt-30 text-white flex items-end">
        <div>
          <div className="p-4 space-y-3 text-xl group-hover:opacity-100 group-hover:translate-y-0 translate-y-4 pb-10 transform transition duration-300 ease-in-out">
            <div className="font-bold">{title}</div>
          </div>
        </div>
      </div>
      <img
        alt={title}
        className="w-full group-hover:scale-110 transition duration-300 ease-in-out"
        src={image}
      />
    </div>
  );
};

export default ImageCard;
