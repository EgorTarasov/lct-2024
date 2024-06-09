import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable } from "mobx";
import { MapFilters, MapFiltersLocale } from "@/constants/map-filters";
import { Filter } from "./filter.vm";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import mapboxgl from "mapbox-gl";
import geojsonvt from "geojson-vt";

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
  map: mapboxgl.Map | null = null;
  filters = new FiltersViewModel();
  constructor() {
    makeAutoObservable(this);
  }

  setMap(m: mapboxgl.Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;

    this.map.on("load", () => {
      const map = this.map!;

      map.getStyle().layers.forEach((layer) => {
        if (layer.id.endsWith("-label")) {
          map.setLayoutProperty(layer.id, "text-field", [
            "coalesce",
            ["get", "name_ru"],
            ["get", "name"]
          ]);
        }
      });

      // Your GeoJSON data
      const geojsonData: GeoJSON.FeatureCollection = {
        type: "FeatureCollection",
        features: [
          {
            type: "Feature",
            geometry: {
              type: "Point",
              coordinates: [-100.486052, 37.830348]
            },
            properties: {
              title: "Mapbox",
              description: "Washington, D.C."
            }
          }
        ]
      };

      // Create the tile index
      const tileIndex = geojsonvt(geojsonData, {});

      map.addSource("geojson-vt-source", {
        type: "vector",
        tiles: ["http://localhost:3000/tiles/{z}/{x}/{y}.pbf"],
        minzoom: 0,
        maxzoom: 14
      });

      map.addLayer({
        id: "geojson-vt-layer",
        type: "circle",
        source: "geojson-vt-source",
        "source-layer": "geojsonLayer",
        paint: {
          "circle-radius": 6,
          "circle-color": "#B42222"
        }
      });
    });
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
