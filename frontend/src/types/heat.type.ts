import { Issue } from "./issue.type";
import { Priority } from "./priority.type";

export namespace HeatSource {
  export interface Item {
    id: number;
    number: string;
    address: string;
    issue: Issue;
    priority: Priority.Item;
    consumerCount: number;
    issues: Issue[];
    incidentCount: number;
  }
}
