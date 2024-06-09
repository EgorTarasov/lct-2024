import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import L, { Map } from "leaflet";
import { MapFilters, MapFiltersLocale } from "@/constants/map-filters";
import { Filter } from "./filter.vm";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";

class FiltersViewModel implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
  }

  search = "";

  // heatNetworks: MapFilters.HeatNetwork[] = [];
  heatNetworks = new Filter(Object.values(MapFilters.HeatNetwork), MapFiltersLocale.HeatNetwork);

  dispose(): void {
    this.search = "";
    this.heatNetworks.reset();
  }
}

class mapStore implements DisposableVm {
  private map: Map | null = null;
  filters = new FiltersViewModel();
  constructor() {
    makeAutoObservable(this);
  }

  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
    this.map.addEventListener("moveend", () => this.moveEnd());
  }

  moveEnd() {
    if (!this.map) return;

    const zoom = this.map.getZoom();
    const center = this.map.getCenter();

    console.log(zoom, center);
    MapEndpoint.getMoek(center.lat, center.lng, 1000);
  }

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
