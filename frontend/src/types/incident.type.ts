import { IncidentDto } from "@/api/models/incident.model";
import { Consumer } from "./consumer.type";
import { HeatDistributor } from "./heat.type";
import { Issue } from "./issue.type";
import { MapFilters } from "./map-filters";
import { Priority } from "./priority.type";

export namespace Incident {
  export interface DependentObjects {
    tps: string;
    heatSource: HeatDistributor.Item | null;
    consumers: Consumer.Item[];
  }

  export interface BaseProps {
    number: string;
    priority: Priority;
    address: string;
    info: [title: string, value: string][];
    incidentTitle: string;
    incidentIssue: Issue;
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

  export const convertDto = (dto: IncidentDto.Item): Incident.Item => {
    const dependentObjects: Incident.DependentObjects = {
      tps: "", // Populate based on dto
      heatSource: null, // Convert dto to HeatDistributor.Item
      consumers: [] // Convert dto to array of Consumer.Item
    };

    const baseProps: Incident.BaseProps = {
      number: dto.id.toString(),
      priority: Priority.HIGH,
      address: dto.relatedObjects.heatingPoint?.heating_point_full_address.address || "",
      info: [
        // Add more info based on dto
      ],
      incidentTitle: dto.title,
      incidentIssue: dto.status as Issue,
      incidentStatus: dto.closedAt ? "closed" : "active",
      date: new Date(dto.createdAt),
      comments: [], // Add comments if available
      unom: dto.unom,
      dependentObjects
    };

    // Assuming we need to determine the type based on some DTO fields
    if (dto.relatedObjects.heatingPoint) {
      return {
        ...baseProps,
        type: "heat-source",
        data: dependentObjects.heatSource
      } as Incident.HeatItem;
    } else {
      return {
        ...baseProps,
        type: "consumer",
        data: dependentObjects.consumers[0] // Assuming first consumer for simplicity
      } as Incident.ConsumerItem;
    }
  };
}
