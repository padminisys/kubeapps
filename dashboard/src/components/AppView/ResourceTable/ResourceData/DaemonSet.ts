import { get } from "lodash";
import { IResource } from "shared/types";

export const DaemonSetColumns = [
  {
    accessor: "name",
    Header: "NAME",
    getter: (target: IResource) => get(target, "metadata.name"),
  },
  {
    accessor: "desired",
    Header: "DESIRED",
    getter: (target: IResource) => get(target, "status.currentNumberScheduled"),
  },
  {
    accessor: "available",
    Header: "AVAILABLE",
    getter: (target: IResource) => get(target, "status.numberReady"),
  },
];
