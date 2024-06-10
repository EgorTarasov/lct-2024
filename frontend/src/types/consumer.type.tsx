import { Priority } from "./priority.type";

export namespace Consumer {
  export interface Item {
    id: number;
    position: number[][];
    priority: Priority.Item;
  }
}
