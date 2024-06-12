import { Issue } from "./issue.type";
import { Priority } from "./priority.type";

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
    issue: Issue;
    priority: Priority;
    consumerCount: number;
    issues: Issue[];
    incidentCount: number;
  }
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
