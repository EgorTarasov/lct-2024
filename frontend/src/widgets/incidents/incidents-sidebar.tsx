import { FCVM } from "@/utils/vm";
import { observer } from "mobx-react-lite";
import { IncidentsPageViewModel } from "./incidents.page.vm";
import { IconInput, Input } from "@/components/ui/input";
import { SearchIcon } from "lucide-react";
import { Text } from "@/components/typography/Text";
import { cn } from "@/utils/cn";
import { IncidentsTabs } from "./incidents-tabs";

export const IncidentsSidebar: FCVM<IncidentsPageViewModel> = observer(({ vm }) => {
  return (
    <aside className="h-full flex flex-col w-96 bg-card py-4 overflow-hidden">
      <div className="flex flex-col px-4">
        <IconInput
          containerClassName="flex-1"
          className="bg-background"
          rightIcon={<SearchIcon />}
          placeholder="Введите данные UNOM"
        />
        <Text.UiMedium className="text-muted-foreground pt-5 pb-3">Инциденты</Text.UiMedium>
        <IncidentsTabs vm={vm} />
      </div>
    </aside>
  );
});
