import { ConsumerCard } from "@/components/cards/consumer.card";
import { HeatSourceCard } from "@/components/cards/heat-source.card";
import { Text } from "@/components/typography/Text";
import { Breadcrumb, BreadcrumbItem } from "@/components/ui/breadcrumb";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { createLazyFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";

const vm = MapStore;

const Page = observer(() => (
  <MainSidebar>
    <div className="gap-3 h-full overflow-hidden flex flex-col">
      <Text.SubtleSemi className="px-4 text-muted-foreground">Реестр объектов</Text.SubtleSemi>
      <Breadcrumb className="px-4">
        <BreadcrumbItem>Источники тепла</BreadcrumbItem>
      </Breadcrumb>
      <ScrollArea>
        {vm.heatSources.map((v) => (
          <>
            <HeatSourceCard data={v} />
            <Separator />
          </>
        ))}
      </ScrollArea>
    </div>
  </MainSidebar>
));

export const Route = createLazyFileRoute("/_map/")({
  component: Page
});
