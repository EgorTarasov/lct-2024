import { IncidentsEndpoint } from "@/api/endpoints/incidents.endpoint";
import { IncidentFilters } from "@/types/incident-filters";
import { Incident } from "@/types/incident.type";
import { Issue } from "@/types/issue.type";
import { Priority } from "@/types/priority.type";
import { debounce } from "@/utils/debounce";
import { makeAutoObservable, toJS } from "mobx";

class incidentsPageViewModel {
  drawerOpen = false;
  selected: Incident.Item | null = null;

  constructor() {
    makeAutoObservable(this);
  }

  async init() {
    if (this.items.length > 0) return;

    this.loading = true;
    try {
      const res = await IncidentsEndpoint.getRecents();
      this.items = res.map((v) => Incident.convertDto(v));
    } finally {
      this.loading = false;
    }
  }

  //#region filters
  selectedTab: IncidentFilters.Tab = "heat-source";
  issueType: Issue | null = null;
  sort: "new" | "old" | "priority" = "priority";
  loading = true;
  search = "";
  setSearch = debounce((v: string) => {
    this.search = v;
  }, 300);

  items: Incident.Item[] = [];

  get heatSources(): Incident.HeatItem[] {
    if (this.selectedTab !== "heat-source") {
      return [];
    }

    const items = this.items.filter(
      (i) =>
        i.type === "heat-source" &&
        (!this.issueType || i.incidentIssue === this.issueType) &&
        i.unom.toString().includes(this.search)
    ) as Incident.HeatItem[];
    return items;
  }

  get consumers(): Incident.ConsumerItem[] {
    if (this.selectedTab === "consumer") {
      console.log(toJS(this.items.filter((i) => i.type === "consumer")));
      return this.items.filter((i) => i.type === "consumer") as Incident.ConsumerItem[];
    }
    return [];
  }
  //#endregion

  async select(unom: string) {
    this.selected = this.items.find((i) => i.unom.toString() === unom) || null;
    this.drawerOpen = false;
  }

  async resolve() {
    console.log("resolve");
  }
}

export const IncidentsPageViewModel = new incidentsPageViewModel();
