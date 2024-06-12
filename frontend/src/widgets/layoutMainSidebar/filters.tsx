import DropdownMultiple from "@/components/DropdownMultiple";
import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import { MapStore } from "@/stores/map.store";
import { observer } from "mobx-react-lite";

export const MainSidebarFilters = observer(() => {
  const vm = MapStore.heatSourceVm;

  return (
    <div className="px-4 pb-4 gap-3 h-full overflow-hidden flex flex-col">
      <Text.UiMedium className="text-muted-foreground">Фильтры</Text.UiMedium>
      <div className="space-y-2 flex-1">
        <DropdownMultiple {...vm.heatNetworks.attributes} label="Тепловая сеть" />
      </div>
      <Button variant="secondary" onClick={() => vm.reset()}>
        Сбросить фильтры
      </Button>
    </div>
  );
});
