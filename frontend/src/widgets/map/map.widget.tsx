import "leaflet/dist/leaflet.css";
import "./draw/draw-locale";
import "./map.css";
import { observer } from "mobx-react-lite";
import { MapContainer, TileLayer, Marker, Popup, ZoomControl } from "react-leaflet";
import { MapInitializer } from "./map-initializer";

const Map = observer(() => {
  return (
    <MapContainer
      className="w-full h-screen absolute"
      center={[55.751244, 37.618423]}
      zoom={13}
      zoomControl={false}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
      />
      <Marker position={[51.505, -0.09]}>
        <Popup>
          A pretty CSS3 popup. <br /> Easily customizable.
        </Popup>
      </Marker>
      <ZoomControl position="topright" />
      <MapInitializer />
    </MapContainer>
  );
});

export default Map;
