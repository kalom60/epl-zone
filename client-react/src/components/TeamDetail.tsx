import { useParams } from "react-router-dom";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { useEffect, useState } from "react";

type State = {
  id: string;
  player: string;
  nation: string;
  position: string;
  age: number;
  matchesPlayed: number;
  starts: number;
  minutesPlayed: number;
  goals: number;
  assists: number;
  penalitiesScored: number;
  yellowCards: number;
  redCards: number;
  expectedGoals: number;
  expectedAssists: number;
  teamName: string;
};

const TeamDetail = () => {
  const { name } = useParams();
  const [details, setDetails] = useState<State[]>([]);

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
      <Table>
        <TableCaption>
          A list of {name?.split("-").join(" ")} players
        </TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Player</TableHead>
            <TableHead>Nation</TableHead>
            <TableHead>Position</TableHead>
            <TableHead>Age</TableHead>
            <TableHead>Matches Played</TableHead>
            <TableHead>Starts</TableHead>
            <TableHead>Minutes Played</TableHead>
            <TableHead>Goals</TableHead>
            <TableHead>Assists</TableHead>
            <TableHead>Penalities Scored</TableHead>
            <TableHead>Yellow Cards</TableHead>
            <TableHead>Red Cards</TableHead>
            <TableHead>Expected Goals</TableHead>
            <TableHead>Expected Assists</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {details.map((detail: State) => (
            <TableRow key={detail.id}>
              <TableCell className="font-medium">{detail.player}</TableCell>
              <TableCell>{detail.nation}</TableCell>
              <TableCell>{detail.position}</TableCell>
              <TableCell>{detail.age}</TableCell>
              <TableCell>{detail.matchesPlayed}</TableCell>
              <TableCell>{detail.starts}</TableCell>
              <TableCell>{detail.minutesPlayed}</TableCell>
              <TableCell>{detail.goals}</TableCell>
              <TableCell>{detail.assists}</TableCell>
              <TableCell>{detail.penalitiesScored}</TableCell>
              <TableCell>{detail.yellowCards}</TableCell>
              <TableCell>{detail.redCards}</TableCell>
              <TableCell>{detail.expectedGoals}</TableCell>
              <TableCell>{detail.expectedAssists}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};

export default TeamDetail;
