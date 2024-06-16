import { MapContainer, TileLayer, ZoomControl } from "react-leaflet";
import "@/widgets/map/map.css";
import { observer } from "mobx-react-lite";
import { FCVM } from "@/utils/vm";
import { IncidentsPageViewModel } from "../incidents.page.vm";
import { Map } from "leaflet";
import { MapInitializer } from "@/widgets/map/map-initializer";
import { useEffect, useState } from "react";
import "leaflet/dist/leaflet.css";
import { ELEVATION } from "@/constants/elevation";
import { ConsumerPolygon } from "@/widgets/map/polygon";

const vm = IncidentsPageViewModel;

export const IncidentsMap = observer(() => {
  return (
    <MapContainer
      center={[55.751244, 37.618423]}
      zoom={13}
      style={{ height: 450, zIndex: ELEVATION.MAP }}>
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      {/* <ConsumerPolygon data={vm.selected?.polygon} /> */}
    </MapContainer>
  );
});

export const MyMap = observer(() => {
  const [map, setMap] = useState<Map | null>(null);

  useEffect(() => {}, []);

  return null;
  // <MapContainer
  //   className="flex-1 h-[450px] overflow-hidden"
  //   center={[55.751244, 37.618423]}
  //   zoom={13}
  //   maxZoom={18}
  //   zoomControl={false}>
  //   <TileLayer
  //     url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
  //     attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
  //   />
  //   <ZoomControl position="topright" />
  //   <MapInitializer
  //     setMap={(v) => {
  //       if (map) {
  //         map.remove();
  //       }
  //       setMap(v);
  //     }}
  //   />
  // </MapContainer>
});
