import { MapStore } from "@/stores/map.store";
import { Consumer } from "@/types/consumer.type";
import { PriorityLocaleMap } from "@/types/priority.type";
import { LatLngBounds, LatLngBoundsLiteral, LatLngExpression } from "leaflet";
import { observer } from "mobx-react-lite";
import { FC, useEffect } from "react";
import { Marker, Polygon, Popup } from "react-leaflet";

const vm = MapStore;

export const ConsumerPolygon: FC<{ data: Consumer.Polygon }> = observer((x) => {
  useEffect(() => {
    if (vm.map) {
      const bounds = x.data.position as LatLngBoundsLiteral;
      vm.map.fitBounds(bounds, {
        maxZoom: 17,
        animate: true,
        paddingTopLeft: [300, 0]
      });
    }
  }, [x.data]);

  return (
    <>
      <Polygon positions={x.data.position as LatLngExpression[]} color="#16A34A" />
      {/* <Marker position={getTopLeftCorner() as LatLngExpression}>
        <Popup>Polygon {x.data.id}</Popup>
      </Marker> */}
    </>
  );
});
