import DropdownMultiple from "@/components/DropdownMultiple";
import { Text } from "@/components/typography/Text";
import { MapStore } from "@/stores/map.store";
import { observer } from "mobx-react-lite";

export const Filters = observer(() => {
  const vm = MapStore.filters;

  return (
    <div className="space-y-3">
      <Text.SubtleSemi className="text-slate-500">Фильтры</Text.SubtleSemi>
      <div className="space-y-2">
        <DropdownMultiple {...vm.heatNetworks.attributes} label="Тепловая сеть" />
      </div>
    </div>
  );
});
