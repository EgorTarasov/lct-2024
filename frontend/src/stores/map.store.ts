import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable, toJS, when } from "mobx";
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
import { MapDto } from "@/api/models/map.model";
import { cloneDeep, debounce } from "lodash";
import { ConsumersEndpoint } from "@/api/endpoints/consumers.endpoint";

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
    const filters = await ConsumersEndpoint.getFilters();
    this.filters = filters.map((f) => new Filter(f.filterName, f.values));

    const res = await MapEndpoint.getProperty(0, 0, 0);
    const polygons: MapConstants.PolygonFeature[] = [];
    res.forEach((v) => {
      const feature = buildPropertyFeature(v);

      if (feature) {
        polygons.push(feature);
      }
    });

    this.consumers = polygons;

    this.buildFeatureLayer();
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

  get filteredConsumers() {
    return this.consumers;
  }
  //#endregion

  //#region map
  // get filteredGeozones() {
  // }

  shouldZoomIn = true;
  private map: Map | null = null;

  setMap(m: Map) {
    if (this.map) {
      this.dispose();
    }

    this.map = m;
    this.map.addEventListener("moveend", () => {
      this.fetchVisiblePart();
    });

    if (this.featureLayer) {
      this.featureLayer.addTo(this.map);
    } else {
      when(() => this.featureLayer !== null).then(() => {
        this.featureLayer!.addTo(this.map!);
        console.log(this.featureLayer);
      });
    }
  }

  consumers: MapConstants.PolygonFeature[] = [];

  featureLayer: L.VectorGrid.Slicer | null = null;
  buildFeatureLayer() {
    if (this.featureLayer) {
      console.log("disposed");
      this.featureLayer.remove();
    }

    if (!this.filteredConsumers.length) {
      return null;
    }

    const featureCollection = {
      type: "FeatureCollection" as const,
      features: this.filteredConsumers
    };

    const layer = buildSlicerLayer(
      toJS(this.filteredConsumers),
      MapConstants.PolygonProperties.priority.high,
      {
        onClick: (v) => this.onLayerClick(v)
      }
    );

    this.featureLayer = layer;
  }

  onLayerClick(v: MapConstants._ConsumerFeatureProperty) {
    console.log(v);
  }

  fetchVisiblePart = debounce(() => {
    if (!this.map) return;

    const zoom = this.map.getZoom();
    const center = this.map.getCenter();
    const radiusKm = 0.5 * Math.pow(2, 14 - zoom);
    if (radiusKm > 0.4) {
      this.shouldZoomIn = true;
      return;
    }
    this.shouldZoomIn = false;

    console.log(zoom, center);
  }, 1000);
  //#endregion

  dispose(): void {
    this.map = null;
  }
}

export const MapStore = new mapStore();
