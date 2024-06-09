import DropdownMultiple from "@/components/DropdownMultiple";
import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import { MapStore } from "@/stores/map.store";
import { observer } from "mobx-react-lite";

export const Filters = observer(() => {
  const vm = MapStore.filters;

  return (
    <div className="gap-3 h-full overflow-hidden flex flex-col">
      <Text.SubtleSemi className="text-slate-500">Фильтры</Text.SubtleSemi>
      <div className="space-y-2 flex-1">
        <DropdownMultiple {...vm.heatNetworks.attributes} label="Тепловая сеть" />
      </div>
      <Button variant="secondary" onClick={() => vm.reset()}>
        Сбросить фильтры
      </Button>
    </div>
  );
});