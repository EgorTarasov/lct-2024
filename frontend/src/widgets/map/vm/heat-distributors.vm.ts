import { MapFilters, MapFiltersLocale } from "@/constants/map-filters";
import { Filter } from "@/stores/filter.vm";
import { DisposableVm } from "@/utils/vm";
import { makeAutoObservable, reaction } from "mobx";
import { HeatDistributor } from "@/types/heat.type";
import { Issue } from "@/types/issue.type";
import { debounceAsync } from "@/utils/debounce";
import { Priority } from "@/types/priority.type";

export class HeatDistributorsViewModel implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
    const r = reaction(
      () => [this.search, this.heatNetworks.value, this.layer, this.page],
      () => this.fetchList()
    );
  }

  //#region filters
  search = "";
  heatNetworks = new Filter(Object.values(MapFilters.HeatNetwork), MapFiltersLocale.HeatNetwork);
  //#endregion

  page = 0;
  totalPages: number | null = 2;
  loading = false;
  prevPage() {
    if (this.page === 0) return;
    this.page--;
  }
  items: HeatDistributor.Item[] = [
    {
      address: "ул. Ленина, 1",
      consumerCount: 10,
      id: 1,
      issue: Issue.EMERGENCY,
      issues: [Issue.EMERGENCY, Issue.REPAIR],
      number: "ТЭЦ-1",
      priority: Priority.HIGH,
      incidentCount: 3
    }
  ];

  fetchList = debounceAsync(async () => {
    const filters = {
      search: this.search,
      heatNetworks: this.heatNetworks.value,
      layer: this.layer
    };

    console.log(filters);
  }, 1000);

  reset() {
    this.search = "";
    this.heatNetworks.reset();
  }

  dispose(): void {
    this.reset();
  }
}
