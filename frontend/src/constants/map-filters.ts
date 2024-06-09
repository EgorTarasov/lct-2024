export namespace MapFilters {
  export enum HeatNetwork {
    Main = "main",
    Distribution = "distribution",
    Consumers = "consumers"
  }
}

export namespace MapFiltersLocale {
  export const HeatNetwork = {
    [MapFilters.HeatNetwork.Main]: "Магистральная сеть",
    [MapFilters.HeatNetwork.Distribution]: "Распределительная сеть",
    [MapFilters.HeatNetwork.Consumers]: "Потребители с ИТП"
  };
}
