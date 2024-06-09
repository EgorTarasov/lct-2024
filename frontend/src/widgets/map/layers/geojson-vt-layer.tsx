/* eslint-disable @typescript-eslint/no-explicit-any */
import { FC, useEffect } from "react";
import { useMap } from "react-leaflet";
import geojsonvt from "geojson-vt";
import L from "leaflet";
import "leaflet.vectorgrid";
import { observer } from "mobx-react-lite";

export const GeoJSONVTTileLayer: FC<{ data: geojsonvt.Data[] }> = observer(({ data }) => {
  const map = useMap();

  useEffect(() => {
    const layers = data.map((geojson) => {
      const tileIndex = geojsonvt(geojson, {
        maxZoom: 14,
        tolerance: 3,
        extent: 4096,
        buffer: 64,
        debug: 0,
        indexMaxZoom: 5,
        indexMaxPoints: 100000
      });

      return (L.vectorGrid as any)
        .slicer(tileIndex, {
          rendererFactory: (L.canvas as any).tile,
          vectorTileLayerStyles: {
            sliced: {
              weight: 1,
              color: "#000000",
              opacity: 1,
              fillColor: "#000000",
              fillOpacity: 0.5,
              stroke: true,
              fill: true
            }
          }
        })
        .addTo(map);
    });

    console.log(layers);

    return () => {
      console.log("ok");
      // layers.forEach((layer) => map.removeLayer(layer));
    };
  }, [map, data]);

  return null;
});
