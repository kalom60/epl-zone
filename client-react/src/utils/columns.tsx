import { ColumnDef } from "@tanstack/react-table";
import { ArrowUpDown } from "lucide-react";

import { Button } from "@/components/ui/button";

export type Data = {
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

const columnConfig = [
  {
    key: "player",
    header: "Player",
    className: "capitalize",
    sortable: true,
  },
  {
    key: "nation",
    header: "Nation",
    render: (value: string) => value && value.split(" ")[1]
  },
  {
    key: "position",
    header: "Position",
    render: (value: string) =>
      value && value.length > 3
        ? `${value.substring(0, 2)},${value.substring(2)}`
        : value,
  },
  {
    key: "age",
    header: "Age",
  },
  {
    key: "matchesPlayed",
    header: "Matches Played",
  },
  {
    key: "starts",
    header: "Starts",
  },
  {
    key: "minutesPlayed",
    header: "Minutes Played",
  },
  {
    key: "goals",
    header: "Goals",
  },
  {
    key: "assists",
    header: "Assists",
  },
  {
    key: "penalitiesScored",
    header: "Penalities Scored",
  },
  {
    key: "yellowCards",
    header: "Yellow Cards",
  },
  {
    key: "redCards",
    header: "Red Cards",
  },
  {
    key: "expectedGoals",
    header: "Expected Goals",
  },
  {
    key: "expectedAssists",
    header: "Expected Assists",
  },
  {
    key: "teamName",
    header: "Team",
    render: (value: string) => value.replace(/-/g, " "),
  },
];

export const columns: ColumnDef<Data>[] = columnConfig.map((col) => ({
  accessorKey: col.key,
  header: col.sortable
    ? ({ column }) => (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          {col.header}
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      )
    : col.header,
  cell: ({ row }) => (
    <div className={col.className}>
      {col.render ? col.render(row.getValue(col.key)) : row.getValue(col.key)}
    </div>
  ),
}));
