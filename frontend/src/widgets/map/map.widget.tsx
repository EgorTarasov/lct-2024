import "leaflet/dist/leaflet.css";
import "./draw/draw-locale";
import "./map.css";
import { observer } from "mobx-react-lite";
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  ZoomControl,
  Polygon,
  GeoJSON
} from "react-leaflet";
import { MapInitializer } from "./map-initializer";
import { MapStore } from "@/stores/map.store";
import { useEffect } from "react";
import { toast } from "sonner";
import { ELEVATION } from "@/constants/elevation";

const vm = MapStore;

const Map = observer(() => {
  // useEffect(() => {
  //   if (vm.shouldZoomIn) {
  //     setTimeout(() => {
  //       toast.info("Для отображения объектов увеличьте карту", {
  //         id: "map-zoom-notifier",
  //         dismissible: false,
  //         duration: Infinity,
  //         important: true
  //       });
  //     }, 300);
  //   } else {
  //     toast.dismiss("map-zoom-notifier");
  //   }

  //   return () => {
  //     toast.dismiss("map-zoom-notifier");
  //   };
  // }, [vm.shouldZoomIn]);

  return (
    <>
      <MapContainer
        className="w-full h-screen absolute"
        center={[55.751244, 37.618423]}
        zoom={13}
        style={{ zIndex: ELEVATION.MAP }}
        maxZoom={18}
        zoomControl={false}>
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        <ZoomControl position="topright" />
        <MapInitializer setMap={(v) => vm.setMap(v)} />
      </MapContainer>
    </>
  );
});

export default Map;
