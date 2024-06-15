import { observer } from "mobx-react-lite";
import { WeatherWidget } from "../weather/weather.widget";
import { MapStore } from "@/stores/map.store";
import { ELEVATION } from "@/constants/elevation";

export const BottomRightBar = observer(() => {
  return (
    <div
      className="absolute right-4 bottom-4 flex gap-2 z-[11]"
      style={{ zIndex: ELEVATION.WEATHER }}>
      {MapStore.filteredConsumers.length > 0 && (
        <div className="appear flex items-center justify-center bg-card border rounded-xl px-2 gap-2">
          Найдено объектов: {MapStore.filteredConsumers.length}
        </div>
      )}
      <WeatherWidget />
    </div>
  );
});
