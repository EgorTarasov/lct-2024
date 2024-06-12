import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import { Map } from "leaflet";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import { debounceAsync } from "@/utils/debounce";
import { HeatSourcesViewModel } from "@/widgets/map/vm/filters.vm";
import { Consumer } from "@/types/consumer.type";
import { Priority } from "@/types/priority.type";
import { ConsumersViewModel } from "../widgets/map/vm/consumers.vm";

export class mapStore implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
  }

  consumersVm: ConsumersViewModel | null = null;
  heatSourceVm = new HeatSourcesViewModel(this);

  //#region map
  consumerGeozones: Consumer.Polygon[] = [];
  shouldZoomIn = true;
  private map: Map | null = null;

  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
    this.map.addEventListener("moveend", () => this.fetchVisiblePart());
  }

  fetchVisiblePart = debounceAsync(async (signal) => {
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
    const polygons: Consumer.Polygon[] = [];
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

    this.consumerGeozones = polygons;
  }, 500);
  //#endregion

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
