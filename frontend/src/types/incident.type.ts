import { Consumer } from "./consumer.type";
import { HeatDistributor } from "./heat.type";
import { Issue } from "./issue.type";
import { Priority } from "./priority.type";

export namespace Incident {
  export interface DependentObjects {
    tps: string;
    heatSource: HeatDistributor.Item;
    consumers: Consumer.Item[];
  }

  export interface BaseProps {
    number: string;
    priority: Priority;
    address: string;
    info: [title: string, value: string][];
    incidentTitle: string;
    incidentType: Issue;
    incidentStatus: "active" | "closed";
    date: Date;
    comments: string[];
    unom: number;
    dependentObjects: DependentObjects;
  }

  export interface HeatItem extends BaseProps {
    type: "heat-source";
    data: HeatDistributor.Item;
  }

  export interface ConsumerItem extends BaseProps {
    type: "consumer";
    data: Consumer.Item;
  }

  export type Item = HeatItem | ConsumerItem;
}
