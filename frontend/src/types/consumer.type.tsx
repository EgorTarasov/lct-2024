import { Priority } from "./priority.type";

export namespace Consumer {
  export interface Polygon {
    id: number;
    position: number[][];
    priority: Priority.Item;
  }

  export interface Item {
    id: number;
  }
}
