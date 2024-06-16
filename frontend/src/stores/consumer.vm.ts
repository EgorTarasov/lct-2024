import { ConsumersEndpoint } from "@/api/endpoints/consumers.endpoint";
import { HeatDistributorsEndpoint } from "@/api/endpoints/heat-distributors.endpoint";
import { Consumer } from "@/types/consumer.type";
import { HeatDistributor } from "@/types/heat.type";
import { makeAutoObservable, when } from "mobx";
import { MapStore } from "./map.store";
import { Priority } from "@/types/priority.type";

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

      await when(() => MapStore.consumers.length > 0);
      const polygon = MapStore.consumers.find((v) => {
        return v.properties.data.consumerAddress.unom.toString() === consumerUnom;
      });

      if (polygon) {
        this.item.name = polygon.properties.data.consumerAddress.address;
        const consumerPolygon = {
          id: polygon.properties.data.consumerAddress.unom.toString(),
          position: polygon.geometry.coordinates[0].map((v) => [v[1], v[0]]),
          data: {
            number: this.item.name,
            source: this.item.consumerType,
            balanceHolder: this.item.address,
            commissioningDate: 0,
            priority: Priority.HIGH
          }
        };
        MapStore.setHighlightedConsumer(consumerPolygon);
      } else {
        MapStore.setHighlightedConsumer(null);
      }
    } finally {
      this.loading = false;
    }
  }
}
