import { IncidentFilters } from "@/types/incident-filters";
import { Incident } from "@/types/incident.type";
import { Issue } from "@/types/issue.type";
import { MapFilters } from "@/types/map-filters";
import { Priority } from "@/types/priority.type";
import { makeAutoObservable } from "mobx";

export class IncidentsPageViewModel {
  constructor() {
    makeAutoObservable(this);
  }

  selectedTab: IncidentFilters.Tab = "heat-sources";
  issueType: MapFilters.Layer = MapFilters.Layer.AllObjects;
  sort: "new" | "old" | "priority" = "priority";
  loading = false;

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
            unom: "123"
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
      incidentType: Issue.EMERGENCY,
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
    if (this.selectedTab === "heat-sources") {
      return this.items.filter((i) => i.type === "heat-source") as Incident.HeatItem[];
    }
    return [];
  }

  get consumers(): Incident.ConsumerItem[] {
    if (this.selectedTab === "consumers") {
      return this.items.filter((i) => i.type === "consumer") as Incident.ConsumerItem[];
    }
    return [];
  }

  drawerOpen = false;
}
