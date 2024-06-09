import { observer } from "mobx-react-lite";
import { useEffect, useRef } from "react";
import { useMap } from "react-leaflet";
import { MapStore } from "../../stores/map.store";

export const MapInitializer = observer(() => {
  const map = useMap();
  const initialized = useRef(false);

  useEffect(() => {
    if (!initialized.current) {
      initialized.current = true;
      MapStore.setMap(map);
    }
  }, [map]);

  return null;
});
