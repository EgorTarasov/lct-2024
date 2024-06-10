import { mapStore } from "@/stores/map.store";

import { MapFilters, MapFiltersLocale } from "@/constants/map-filters";
import { Filter } from "@/stores/filter.vm";
import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";

export class FiltersViewModel implements DisposableVm {
  constructor(private parent: mapStore) {
    makeAutoObservable(this);
  }

  search = "";

  heatNetworks = new Filter(
    Object.values(MapFilters.HeatNetwork),
    MapFiltersLocale.HeatNetwork,
    () => this.updateFilters()
  );
  _layer: MapFilters.Layer = MapFilters.Layer.AllObjects;
  setLayer(layer: MapFilters.Layer) {
    this._layer = layer;
    this.updateFilters();
  }
  get layer() {
    return this._layer;
  }

  updateFilters() {
    const filters = {
      search: this.search,
      heatNetworks: this.heatNetworks.value,
      layer: this._layer
    };

    console.log(filters);
  }

  reset() {
    this.search = "";
    this.heatNetworks.reset();
  }

  dispose(): void {
    this.reset();
  }
}
