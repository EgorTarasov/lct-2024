import { HeatDistributorDto } from "@/api/models/heat-distribution.model";
import { Issue } from "./issue.type";
import { Priority } from "./priority.type";
import { MapDto } from "@/api/models/map.model";

export namespace HeatDistributor {
  export type Info =
    | "balanceHolder"
    | "commissioningDate"
    | "consumerAddress"
    | "district"
    | "locationType"
    | "number"
    | "source"
    | "type";

  export interface Item {
    id: number;
    number: string;
    address: string;
    issue: Issue | null;
    priority: Priority;
    consumerCount: number;
    issues: Issue[];
    incidentCount: number;
    unom: string;
  }

  export const convertDto = (v: MapDto.Property): HeatDistributor.Item => ({
    id: v.consumer_full_address.unom,
    number: v.heating_point_number,
    address: v.heating_point_full_address.address,
    issue: Issue.EMERGENCY,
    priority: Priority.HIGH,
    consumerCount: 1,
    issues: [Issue.EMERGENCY, Issue.REPAIR],
    incidentCount: 0,
    unom: v.heating_point_full_address.unom.toString()
  });
}

export namespace HeatDistributorLocaleMap {
  export const Info: Record<HeatDistributor.Info, string> = {
    balanceHolder: "Балансодержатель",
    commissioningDate: "Дата ввода в эксплуатацию",
    consumerAddress: "Адрес потребителя",
    district: "Муниципальный район",
    locationType: "Тип по размещению",
    number: "Номер ТП",
    source: "Источник теплоснабжения",
    type: "Вид ТП"
  };
}
