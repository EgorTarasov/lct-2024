import { ConsumersEndpoint } from "@/api/endpoints/consumers.endpoint";
import { HeatDistributorsEndpoint } from "@/api/endpoints/heat-distributors.endpoint";
import { Consumer } from "@/types/consumer.type";
import { HeatDistributor } from "@/types/heat.type";
import { makeAutoObservable } from "mobx";

export class ConsumerViewModel {
  constructor() {
    makeAutoObservable(this);
  }

  item: Consumer.Item | null = null;
  heatDistributor: HeatDistributor.Item | null = null;
  loading = true;

  async init(consumerUnom: string, heatDistributorUnom: string) {
    this.loading = true;
    try {
      const [consumer, heatDistributor] = await Promise.all([
        ConsumersEndpoint.getByUnom(consumerUnom),
        HeatDistributorsEndpoint.getByUnom(heatDistributorUnom)
      ]);
      this.heatDistributor = HeatDistributor.convertDto(heatDistributor);
      this.item = Consumer.convertDto(consumer);
    } finally {
      this.loading = false;
    }
  }
}
