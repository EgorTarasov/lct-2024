import { Consumer } from "@/types/consumer.type";
import { HeatDistributor } from "@/types/heat.type";
import { Issue } from "@/types/issue.type";
import { Priority } from "@/types/priority.type";
import { makeAutoObservable } from "mobx";

export class ConsumersViewModel {
  constructor(private heatDistributorId: string) {
    makeAutoObservable(this);
  }

  items: Consumer.Item[] = [
    {
      address: "10-я Парковая ул., д. 15",
      name: 'ГБУ "ЖИЛИЩНИК РАЙОНА ИЗМАЙЛОВО"',
      id: 1,
      info: {
        type: "социальный"
      },
      issue: Issue.PREDICTION,
      priority: Priority.LOW,
      incidentCount: 2,
      unom: "123"
    }
  ];

  loading = false;
  showPriorityFirst = true;
  heatSource: HeatDistributor.Item | null = null;
}
