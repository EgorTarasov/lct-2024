import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { MapFilters, MapFiltersLocale } from "@/types/map-filters";
import { MapStore } from "@/stores/map.store";
import { ChevronDownIcon } from "lucide-react";
import { observer } from "mobx-react-lite";

export const LayerSelect = observer(() => {
  const vm = MapStore;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" className="bg-card flex gap-2 w-72 text-left px-3">
          <span className="*:size-4">{MapFiltersLocale.Layer[vm.layer].icon}</span>

          <Text.Subtle className="flex-1">{MapFiltersLocale.Layer[vm.layer].locale}</Text.Subtle>
          <ChevronDownIcon className="size-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-72">
        <DropdownMenuGroup>
          {Object.entries(MapFiltersLocale.Layer).map(([key, v], i) => (
            <DropdownMenuItem key={i} asChild>
              <Button
                variant="ghost"
                className="flex gap-2 w-full text-left"
                onClick={() => (vm.layer = key as MapFilters.Layer)}>
                <span className="*:size-4">{v.icon}</span>
                <Text.Subtle className="flex-1">{v.locale}</Text.Subtle>
              </Button>
            </DropdownMenuItem>
          ))}
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
});
