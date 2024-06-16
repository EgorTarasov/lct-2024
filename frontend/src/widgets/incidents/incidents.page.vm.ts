import { IncidentsEndpoint } from "@/api/endpoints/incidents.endpoint";
import { IncidentFilters } from "@/types/incident-filters";
import { Incident } from "@/types/incident.type";
import { Issue } from "@/types/issue.type";
import { Priority } from "@/types/priority.type";
import { debounce } from "@/utils/debounce";
import { makeAutoObservable, toJS } from "mobx";

export class IncidentsPageViewModel {
  drawerOpen = false;
  selected: Incident.Item | null = null;

  constructor() {
    makeAutoObservable(this);

    void this.init();
  }

  async init() {
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

  items: Incident.Item[] = [
    {
      address: "ул. Ленина, 1",
      comments: ["Комментарий 1", "Комментарий 2"],
      type: "heat-source",
      unom: 123,
      dependentObjects: {
        tps: "ТП-1",
        heatSource: {
          address: "ул. Ленина, 1",
          consumerCount: 10,
          id: 1,
          issue: Issue.EMERGENCY,
          issues: [Issue.EMERGENCY, Issue.REPAIR],
          number: "ТЭЦ-1",
          priority: Priority.HIGH,
          incidentCount: 3,
          unom: "123"
        },
        consumers: [
          {
            address: "10-я Парковая ул., д. 15",
            name: 'ГБУ "ЖИЛИЩНИК РАЙОНА ИЗМАЙЛОВО"',
            id: 1,
            info: {
              type: "социальный"
            },
            issue: Issue.PREDICTION,
            priority: Priority.LOW,
            incidentCount: 2,
            unom: "123",
            consumerType: "социальный"
          }
        ]
      },
      data: {
        address: "ул. Ленина, 1",
        consumerCount: 10,
        id: 1,
        issue: Issue.EMERGENCY,
        issues: [Issue.EMERGENCY, Issue.REPAIR],
        number: "ТЭЦ-1",
        priority: Priority.HIGH,
        incidentCount: 3,
        unom: "123"
      },
      date: new Date(),
      incidentStatus: "active",
      incidentTitle: "Инцидент 1",
      incidentIssue: Issue.EMERGENCY,
      info: [
        ["Тип ТП", "ТЭЦ"],
        ["Балансодержатель", "ООО Рога и копыта"],
        ["Дата ввода в эксплуатацию", "01.01.2000"]
      ],
      number: "123",
      priority: Priority.HIGH
    }
  ];

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
  }

  async resolve() {
    console.log("resolve");
  }
}
