import { Issue } from "./issue.type";
import { Priority } from "./priority.type";

export namespace Consumer {
  export interface Polygon {
    id: number;
    position: number[][];
    priority: Priority;
  }

  export type Info = "type";

  export interface Item {
    id: number;
    address: string;
    name: string;
    priority: Priority;
    issue: Issue;
    info: Record<Info, string>;
    incidentCount: 2;
    unom: string;
  }
}
