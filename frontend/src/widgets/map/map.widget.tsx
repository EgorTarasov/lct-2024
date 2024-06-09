import React, { useRef, useEffect } from "react";
import mapboxgl from "mapbox-gl";
import geojsonvt from "geojson-vt";
import { observer } from "mobx-react-lite";
import { useTheme } from "@/components/hoc/theme-provider";
import { MapStore } from "@/stores/map.store";

mapboxgl.accessToken = import.meta.env.VITE_MAPBOX_API_KEY;

const vm = MapStore;

const Map = observer(() => {
  const mapContainerRef = useRef<HTMLDivElement>(null);
  const theme = useTheme();

  useEffect(() => {
    if (!mapContainerRef.current) return;

    const map = new mapboxgl.Map({
      container: mapContainerRef.current,
      style: "mapbox://styles/mapbox/navigation-night-v1",
      center: [55.751244, 37.618423],
      zoom: 3
    });

    vm.setMap(map);

    // Clean up on unmount
    return () => map.remove();
  }, []);

  useEffect(() => {
    vm.map?.setStyle(
      theme.theme === "dark"
        ? "mapbox://styles/mapbox/dark-v11"
        : "mapbox://styles/mapbox/light-v11"
    );
  }, [theme.theme]);

  return <div ref={mapContainerRef} style={{ width: "100%", height: "100vh" }} />;
});

export default Map;
