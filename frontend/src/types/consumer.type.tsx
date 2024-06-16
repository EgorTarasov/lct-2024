import { ConsumersDto } from "@/api/models/consumers.model";
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
    info: Record<string, string>;
    incidentCount: number;
    unom: string;
    consumerType: string;
  }

  export const convertDto = (dto: ConsumersDto.Item): Item => {
    const item = dto.mkdConsumers[0];
    return {
      id: item.externalID,
      address: item.fullAddress,
      name: item.district,
      priority: Priority.HIGH,
      issue: Issue.EMERGENCY,
      info: {
        type: item.typesOfHousingStock.toString()
      },
      incidentCount: item.События.length,
      unom: item.unom.toString(),
      consumerType: item.wallMaterial.toString()
    };
  };
}
