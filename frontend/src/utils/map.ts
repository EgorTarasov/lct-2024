import { MapDto } from "@/api/models/map.model";
import { MapConstants } from "@/constants/map";
import { FeatureCollection } from "geojson";
import geojsonvt from "geojson-vt";
import L from "leaflet";
import "leaflet.vectorgrid";

export const buildPropertyFeature = (v: MapDto.Property): MapConstants.PolygonFeature | null => {
  const coords = v.consumerAddress.border?.find((p) => p.Key === "coordinates")?.Value as
    | number[][][]
    | undefined;
  if (coords) {
    return {
      type: "Feature",
      geometry: {
        type: "Polygon",
        coordinates: [coords[0].map((v) => [v[0], v[1]] as [number, number])]
      },
      properties: {
        type: "consumer",
        data: {
          number: v.number,
          balanceHolder: v.balanceHolder,
          commissioningDate: v.commissioningDate,
          district: v.district,
          locationType: v.locationType,
          Source: v.Source,
          type: v.type,
          consumerAddress: {
            address: v.consumerAddress.address,
            municipalDistrict: v.consumerAddress.municipalDistrict,
            unom: v.consumerAddress.unom,
            center: {
              type: "Point",
              coordinates: v.consumerAddress.center.coordinates as [number, number]
            }
          },
          heatingPointAddress: {
            address: v.heatingPointAddress.address,
            municipalDistrict: v.heatingPointAddress.municipalDistrict,
            unom: v.heatingPointAddress.unom,
            center: {
              type: "Point",
              coordinates: v.heatingPointAddress.center.coordinates as [number, number]
            }
          }
        }
      }
    };
  }

  return null;
};

export const buildSlicerLayer = (
  polygons: MapConstants.PolygonFeature[],
  colors: MapConstants.LayerProperty,
  opt?: {
    onClick?: (e: MapConstants._ConsumerFeatureProperty) => void;
  }
): L.VectorGrid.Slicer => {
  const featureCollection: FeatureCollection = {
    type: "FeatureCollection" as const,
    features: polygons
  };

  const layer = L.vectorGrid.slicer(featureCollection, {
    rendererFactory: L.canvas.tile,
    vectorTileLayerStyles: {
      sliced: {
        weight: 2,
        color: colors.border,
        opacity: 1,
        fillOpacity: 0.6,
        fillColor: colors.fill
      }
    },
    maxZoom: 24,
    maxNativeZoom: 24,
    interactive: true
  });

  layer.on("click", (e: L.LeafletEvent) => {
    if (opt?.onClick) {
      const feature = e.propagatedFrom.feature as MapConstants.PolygonFeature;
      if (feature.properties.type !== "consumer") {
        return;
      }
      opt.onClick(feature.properties.data);
    }
  });

  return layer;
};
