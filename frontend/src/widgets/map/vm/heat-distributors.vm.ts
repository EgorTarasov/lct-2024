import { Filter } from "@/stores/filter.vm";
import { DisposableVm } from "@/utils/vm";
import { IReactionDisposer, makeAutoObservable, reaction } from "mobx";
import { HeatDistributor } from "@/types/heat.type";
import { Issue } from "@/types/issue.type";
import { debounceAsync } from "@/utils/debounce";
import { Priority } from "@/types/priority.type";
import { MapFilters, MapFiltersLocale } from "@/types/map-filters";

export class HeatDistributorsViewModel implements DisposableVm {
  constructor() {
    makeAutoObservable(this);
    this.r = reaction(
      () => [this.search, this.heatNetworks.values, this.page],
      () => this.fetchList()
    );

    this.init();
  }

  r: IReactionDisposer;
  async init() {
    await this.fetchList();
  }

  //#region filters
  search = "";
  heatNetworks = new Filter(
    "Тепловая сеть",
    Object.values(MapFilters.HeatNetwork),
    MapFiltersLocale.HeatNetwork
  );
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
      incidentCount: 3,
      unom: "123"
    }
  ];

  fetchList = debounceAsync(async () => {
    const filters = {
      search: this.search,
      heatNetworks: this.heatNetworks.values
    };
  }, 1000);

  reset() {
    this.search = "";
    this.heatNetworks.reset();
  }

  dispose(): void {
    this.reset();
  }
}
