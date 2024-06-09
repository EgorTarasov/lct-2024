import { AlertTriangleIcon, CalendarClockIcon, NetworkIcon, ThermometerIcon } from "lucide-react";
import React, { ReactNode } from "react";

export type LocaleWithIcon<T extends string> = Record<
  T,
  {
    icon: ReactNode;
    locale: string;
  }
>;

export namespace MapFilters {
  export enum HeatNetwork {
    Main = "main",
    Distribution = "distribution",
    Consumers = "consumers"
  }

  // Все объекты без аварий
  // Аварийные ситуации
  // Прогноз аварийных ситуаций
  // Карта остывания
  export enum Layer {
    AllObjects = "all-objects",
    Emergencies = "emergencies",
    EmergenciesForecast = "emergencies-forecast",
    CoolingMap = "cooling-map"
  }
}

export namespace MapFiltersLocale {
  export const HeatNetwork = {
    [MapFilters.HeatNetwork.Main]: "Магистральная сеть",
    [MapFilters.HeatNetwork.Distribution]: "Распределительная сеть",
    [MapFilters.HeatNetwork.Consumers]: "Потребители с ИТП"
  };

  export const Layer: LocaleWithIcon<MapFilters.Layer> = {
    [MapFilters.Layer.AllObjects]: {
      locale: "Все объекты",
      icon: <NetworkIcon />
    },
    [MapFilters.Layer.Emergencies]: {
      locale: "Аварийные ситуации",
      icon: <AlertTriangleIcon />
    },
    [MapFilters.Layer.EmergenciesForecast]: {
      locale: "Прогноз аварийных ситуаций",
      icon: <CalendarClockIcon />
    },
    [MapFilters.Layer.CoolingMap]: {
      locale: "Карта остывания",
      icon: <ThermometerIcon />
    }
  };
}
