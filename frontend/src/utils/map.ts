import geojsonvt from "geojson-vt";

export const buildGeoJsonPolygon = (coordinates: number[][]): geojsonvt.Data => {
  return {
    type: "Feature",
    properties: {},
    geometry: {
      type: "Polygon",
      coordinates: [coordinates]
    }
  };
};
