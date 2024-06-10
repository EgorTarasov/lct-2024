import { Consumer } from "./consumer.type";
import { Priority } from "./priority.type";

export namespace Notification {
  export interface Item {
    id: number;
    priority: Priority.Item;
    date: Date;
    title: string;
    description: string;
  }
}
