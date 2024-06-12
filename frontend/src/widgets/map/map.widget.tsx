import "leaflet/dist/leaflet.css";
import "./draw/draw-locale";
import "./map.css";
import { observer } from "mobx-react-lite";
import { MapContainer, TileLayer, Marker, Popup, ZoomControl, Polygon } from "react-leaflet";
import { MapInitializer } from "./map-initializer";
import { MapStore } from "@/stores/map.store";
import { useEffect } from "react";
import { toast } from "sonner";
import { ConsumerPolygon } from "./polygon";

const vm = MapStore;

const Map = observer(() => {
  useEffect(() => {
    if (vm.shouldZoomIn) {
      setTimeout(() => {
        toast.info("Для отображения объектов увеличьте карту", {
          id: "map-zoom-notifier",
          dismissible: false,
          duration: Infinity,
          important: true
        });
      }, 300);
    } else {
      toast.dismiss("map-zoom-notifier");
    }

    return () => {
      toast.dismiss("map-zoom-notifier");
    };
  }, [vm.shouldZoomIn]);

  return (
    <>
      <MapContainer
        className="w-full h-screen absolute"
        center={[55.751244, 37.618423]}
        zoom={13}
        zoomControl={false}>
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        />
        <ZoomControl position="topright" />
        {vm.consumerGeozones.map((v) => (
          <ConsumerPolygon key={v.id} data={v} />
        ))}
        <Polygon
          positions={[
            [55.427278678, 37.116022179],
            [55.427192471, 37.116086668],
            [55.427181919, 37.116043016],
            [55.427181915, 37.116043019],
            [55.427165807, 37.115976363],
            [55.427160519, 37.115980205],
            [55.427149873, 37.115936399],
            [55.42715516, 37.115932399],
            [55.427136867, 37.115856604],
            [55.427164554, 37.115835942],
            [55.427159098, 37.115813426],
            [55.427217615, 37.115769738],
            [55.427266639, 37.115972409],
            [55.427268122, 37.115978529],
            [55.42726812, 37.115978531],
            [55.427278678, 37.116022179]
          ]}
          color="blue"
        />
        <MapInitializer />
      </MapContainer>
    </>
  );
});

export default Map;
