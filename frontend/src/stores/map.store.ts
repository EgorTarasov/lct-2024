import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import L, { LatLngExpression, Map } from "leaflet";
import { MapFilters, MapFiltersLocale } from "@/constants/map-filters";
import { Filter } from "./filter.vm";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import geojsonvt from "geojson-vt";
import { debounceAsync } from "@/utils/debounce";

class FiltersViewModel implements DisposableVm {
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

class mapStore implements DisposableVm {
  private map: Map | null = null;
  filters = new FiltersViewModel(this);
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

  shouldZoomIn = true;
  polygons: LatLngExpression[][][] = [];
  moveEnd = debounceAsync(async (signal) => {
    if (!this.map) return;

    const zoom = this.map.getZoom();
    const center = this.map.getCenter();

    const radiusKm = 0.5 * Math.pow(2, 14 - zoom);
    if (radiusKm > 0.2) {
      this.shouldZoomIn = true;
      return;
    }
    this.shouldZoomIn = false;
    const res = await MapEndpoint.getProperty(center.lat, center.lng, radiusKm, signal);
    const polygons: LatLngExpression[][][] = [];
    res.forEach((v) => {
      const coords = v.polygon.find((p) => p.Key === "coordinates")?.Value as
        | LatLngExpression[][][]
        | undefined;
      if (!coords) return;

      polygons.push(coords[0].map((v) => [v[1], v[0]]));
    });

    this.polygons = polygons;
  }, 500);

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
