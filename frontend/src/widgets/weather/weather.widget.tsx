import { WeatherEndpoint } from "@/api/endpoints/weather.endpoint";
import LoadingEllipsis from "@/components/ui/loaders/LoadingEllipsis";
import { ELEVATION } from "@/constants/elevation";
import { useEffect, useState } from "react";
import WeatherIcon from "./assets/weather.svg";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";

export const WeatherWidget = () => {
  const [temperature, setTemperature] = useState<number | null>(null);

  useEffect(() => {
    const fetchWeather = async () => {
      const res = await WeatherEndpoint.getMoscow();

      setTemperature(res.current.temperature_2m);
    };

    fetchWeather();
  }, []);

  return (
    <Tooltip>
      <TooltipTrigger asChild>
        <div
          className="appear flex items-center justify-center bg-card border rounded-xl px-2 gap-2"
          style={{ zIndex: ELEVATION.FILTERS }}>
          <WeatherIcon />
          {temperature ? (
            `${temperature}°С`
          ) : (
            <LoadingEllipsis className="*:bg-slate-300" size={48} />
          )}
        </div>
      </TooltipTrigger>
      <TooltipContent>Погода в Москве</TooltipContent>
    </Tooltip>
  );
};
