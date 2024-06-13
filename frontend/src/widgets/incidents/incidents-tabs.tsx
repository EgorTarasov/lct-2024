import { cn } from "@/utils/cn";
import { FCVM } from "@/utils/vm";
import { observer } from "mobx-react-lite";
import { IncidentsPageViewModel } from "./incidents.page.vm";

export const IncidentsTabs: FCVM<IncidentsPageViewModel> = observer(({ vm }) => {
  return (
    <ul className="flex rounded-md bg-card p-1 border w-fit text-sm gap-1">
      <li>
        <button
          className={cn(
            "hover:bg-muted px-3 py-1.5 rounded-sm transition-colors",
            vm.selectedTab === "heat-sources" && "bg-muted"
          )}
          onClick={() => (vm.selectedTab = "heat-sources")}>
          Источники тепла
        </button>
      </li>
      <li>
        <button
          className={cn(
            "hover:bg-muted px-3 py-1.5 rounded-sm transition-colors",
            vm.selectedTab === "consumers" && "bg-muted"
          )}
          onClick={() => (vm.selectedTab = "consumers")}>
          Потребители
        </button>
      </li>
    </ul>
  );
});
