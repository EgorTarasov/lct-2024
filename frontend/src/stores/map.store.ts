import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import L, { Map } from "leaflet";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import { debounceAsync } from "@/utils/debounce";
import { FiltersViewModel } from "@/widgets/map/vm/filters.vm";
import { Consumer } from "@/types/consumer.type";
import { Priority } from "@/types/priority.type";
import { HeatSource } from "@/types/heat.type";
import { Issue } from "@/types/issue.type";

export class mapStore implements DisposableVm {
  filters = new FiltersViewModel(this);
  constructor() {
    makeAutoObservable(this);
  }

  private map: Map | null = null;
  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
    this.map.addEventListener("moveend", () => this.moveEnd());
  }

  shouldZoomIn = true;
  consumers: Consumer.Item[] = [];
  heatSources: HeatSource.Item[] = [
    {
      address: "ул. Ленина, 1",
      consumerCount: 10,
      id: 1,
      issue: Issue.EMERGENCY,
      issues: [Issue.EMERGENCY, Issue.REPAIR],
      number: "ТЭЦ-1",
      priority: Priority.Item.HIGH
    }
  ];
  moveEnd = debounceAsync(async (signal) => {
    if (!this.map) return;

    const zoom = this.map.getZoom();
    const center = this.map.getCenter();

    const radiusKm = 0.5 * Math.pow(2, 14 - zoom);
    if (radiusKm > 0.4) {
      this.shouldZoomIn = true;
      return;
    }
    this.shouldZoomIn = false;
    const res = await MapEndpoint.getProperty(center.lat, center.lng, radiusKm, signal);
    const polygons: Consumer.Item[] = [];
    res.forEach((v) => {
      const coords = v.polygon.find((p) => p.Key === "coordinates")?.Value as
        | number[][][]
        | undefined;
      if (!coords) return;

      polygons.push({
        id: v.globalID,
        position: coords[0].map((v) => [v[1], v[0]]),
        priority: Priority.Item.LOW
      });
    });

    this.consumers = polygons;
  }, 500);

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
