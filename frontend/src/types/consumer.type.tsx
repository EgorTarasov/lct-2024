import { Issue } from "./issue.type";
import { Priority } from "./priority.type";
import * as GJ from "geojson";

export namespace Consumer {
  export interface Polygon {
    id: string;
    position: number[][];
    data: {
      number: string;
      source: string;
      balanceHolder: string;
      commissioningDate: number;
    };
  }

  export interface Features {
    feature: GJ.Feature<
      GJ.Polygon,
      {
        priority: Priority;
        number: string;
        source: string;
        balanceHolder: string;
        commissioningDate: number;
      }
    >;
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
