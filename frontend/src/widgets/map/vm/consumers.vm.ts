import { Consumer } from "@/types/consumer.type";
import { HeatSource } from "@/types/heat.type";
import { makeAutoObservable } from "mobx";

export class ConsumersViewModel {
  constructor(private heatSource: HeatSource.Item) {
    makeAutoObservable(this);
  }

  items: Consumer.Item[] = [];
  loading = false;
  showPriorityFirst = true;
}
