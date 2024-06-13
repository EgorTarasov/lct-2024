import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable, when } from "mobx";
import { Map } from "leaflet";
import { MapEndpoint } from "@/api/endpoints/map.endpoint";
import { ConsumersViewModel } from "../widgets/map/vm/consumers.vm";
import { HeatDistributorsViewModel } from "@/widgets/map/vm/heat-distributors.vm";
import { Priority } from "@/types/priority.type";
import { DateRange } from "react-day-picker";
import { Filter } from "./filter.vm";
import { MapFilters } from "@/types/map-filters";
import type { FeatureCollection, Feature } from "geojson";
import L from "leaflet";
import "leaflet.vectorgrid";
import { buildPropertyFeature, buildSlicerLayer } from "@/utils/map";
import { MapConstants } from "@/constants/map";

// patch canvas for click events
L.Canvas.Tile.include({
  _onClick: function (e: L.LeafletEvent) {
    const point = this._map.mouseEventToLayerPoint(e).subtract(this.getOffset());
    let layer;
    let clickedLayer;

    for (const id in this._layers) {
      layer = this._layers[id];
      if (
        layer.options.interactive &&
        layer._containsPoint(point) &&
        !this._map._draggableMoved(layer)
      ) {
        clickedLayer = layer;
      }
    }
    if (clickedLayer) {
      clickedLayer.fireEvent(e.type, undefined, true);
    }
  }
});

class mapStore implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
    void this.init();
  }

  consumersVm: ConsumersViewModel | null = null;
  heatSourceVm = new HeatDistributorsViewModel();
  datesWithEvents: Date[] = [new Date()];

  //#region filters
  filters: Filter<string>[] = [];
  search = "";
  async init() {
    // const filters = await ConsumersEndpoint.getFilters();
    // this.filters = filters.map((f) => new Filter(f.filterName, f.values));

    this.fetchVisiblePart();
  }
  dateRange: DateRange = {
    from: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000),
    to: new Date()
  };
  layer: MapFilters.Layer = MapFilters.Layer.AllObjects;

  reset() {
    this.dateRange = {
      from: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000),
      to: new Date()
    };
    this.datesWithEvents = [new Date()];
    this.layer = MapFilters.Layer.AllObjects;
    this.filters.forEach((f) => f.reset());
  }
  //#endregion

  //#region map
  allGeozones: FeatureCollection | null = null;

  // get filteredGeozones() {
  // }

  shouldZoomIn = true;
  private map: Map | null = null;

  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
    // if (this.vectorTileLayer) {
    //   console.log("set 2");
    //   this.vectorTileLayer?.addTo(this.map);
    // } else {
    //   console.log("ye");
    //   when(() => this.vectorTileLayer !== null).then(() => {
    //     console.log(this.vectorTileLayer, "yo");
    //     this.vectorTileLayer.addTo(this.map);
    //   });
    // }
    this.map.addEventListener("moveend", () => {
      const zoom = this.map?.getZoom();
      // this.fetchVisiblePart();
    });
  }

  async fetchVisiblePart() {
    // const zoom = this.map.getZoom();
    // const center = this.map.getCenter();

    // const radiusKm = 0.5 * Math.pow(2, 14 - zoom);
    // if (radiusKm > 0.4) {
    //   this.shouldZoomIn = true;
    //   return;
    // }
    // this.shouldZoomIn = false;
    console.log("running");
    const res = await MapEndpoint.getProperty(0, 0, 0);
    const polygons: MapConstants.PolygonFeature[] = [];
    res.forEach((v) => {
      const feature = buildPropertyFeature(v);

      if (feature) {
        polygons.push(feature);
      }
    });

    // const layer = buildSlicerLayer(polygons, MapConstants.PolygonProperties.priority.high);

    const featureCollection = {
      type: "FeatureCollection" as const,
      features: polygons
    };
    this.allGeozones = featureCollection;

    const layer = L.vectorGrid.slicer(featureCollection, {
      rendererFactory: L.canvas.tile,
      vectorTileLayerStyles: {
        sliced: {
          weight: 2,
          color: "#ff0000",
          opacity: 1,
          fillOpacity: 1,
          fillColor: "#ff0000"
        }
      },
      maxZoom: 24,
      maxNativeZoom: 24,
      interactive: true
    });
    if (!this.map) {
      await when(() => this.map !== null);
    }
    layer.on("click", (e: L.LeafletEvent) => {
      console.log(e);
    });
    layer.addTo(this.map!);
    // this.setGeozones();
  }
  //#endregion

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
