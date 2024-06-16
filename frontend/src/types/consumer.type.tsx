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
      priority: Priority;
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
    if (dto.mkdConsumers && dto.mkdConsumers.length > 1) {
      const item = dto.mkdConsumers[0];

      console.log(item);

      return {
        id: item.externalID,
        address: item.fullAddress,
        name: item.district,
        priority: Priority.HIGH,
        issue: Issue.EMERGENCY,
        info: {
          type: item.typesOfHousingStock.toString()
        },
        incidentCount: item.События?.length ?? 0,
        unom: item.unom.toString(),
        consumerType: item.wallMaterial.toString()
      };
    } else {
      const item = dto.stateHeatConsumers[0];

      console.log(item);

      return {
        id: item.unom,
        address: item.admDistrict,
        name: item.municupalDistrict,
        priority: Priority.LOW,
        issue: Issue.EMERGENCY,
        info: {
          type: item.type
        },
        incidentCount: item.События?.length ?? 0,
        unom: item.unom.toString(),
        consumerType: item.material
      };
    }
  };
}
