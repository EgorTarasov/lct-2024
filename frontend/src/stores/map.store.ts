import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import { Map } from "leaflet";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import { debounceAsync } from "@/utils/debounce";
import { Consumer } from "@/types/consumer.type";
import { ConsumersViewModel } from "../widgets/map/vm/consumers.vm";
import { HeatDistributorsViewModel } from "@/widgets/map/vm/heat-distributors.vm";
import { Priority } from "@/types/priority.type";
import { DateRange } from "react-day-picker";
import { MapFilters } from "@/constants/map-filters";

class mapStore implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
  }

  consumersVm: ConsumersViewModel | null = null;
  heatSourceVm = new HeatDistributorsViewModel();
  dateRange: DateRange = {
    from: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000),
    to: new Date()
  };
  datesWithEvents: Date[] = [new Date()];
  layer: MapFilters.Layer = MapFilters.Layer.AllObjects;

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
        priority: Priority.LOW
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
