import { IncidentFilters } from "@/types/incident-filters";
import { makeAutoObservable } from "mobx";

export class IncidentsPageViewModel {
  constructor() {
    makeAutoObservable(this);
  }

  selectedTab: IncidentFilters.Tab = "heat-sources";
}
