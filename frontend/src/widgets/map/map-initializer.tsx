import { observer } from "mobx-react-lite";
import { useEffect } from "react";
import { useMap } from "react-leaflet";
import { MapViewModel } from "./map.vm";

export const MapInitializer = observer(() => {
  const map = useMap();

  useEffect(() => {
    MapViewModel.setMap(map);
  }, [map]);

  return null;
});
