import { Consumer } from "@/types/consumer.type";
import { Priority } from "@/types/priority.type";
import { LatLngExpression } from "leaflet";
import { FC } from "react";
import { Marker, Polygon, Popup } from "react-leaflet";

export const ConsumerPolygon: FC<{ data: Consumer.Polygon }> = (x) => {
  const getTopLeftCorner = () => {
    return x.data.position.reduce(
      (acc, point) => {
        if (point[1] < acc[1]) {
          return point;
        }
        return acc;
      },
      [99, 99] satisfies LatLngExpression
    );
  };

  return (
    <>
      <Polygon
        positions={x.data.position as LatLngExpression[]}
        color={Priority.ItemMap[x.data.priority].color}
      />
      <Marker position={getTopLeftCorner() as LatLngExpression}>
        <Popup>Polygon {x.data.id}</Popup>
      </Marker>
    </>
  );
};
