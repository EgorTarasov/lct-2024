import { HeatSourceCard } from "@/components/cards/heat-source.card";
import { Text } from "@/components/typography/Text";
import { Breadcrumb, BreadcrumbItem } from "@/components/ui/breadcrumb";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { Outlet, createLazyFileRoute } from "@tanstack/react-router";
import React from "react";

const vm = MapStore;

export const Route = createLazyFileRoute("/_map/_heat_sources")({
  component: () => (
    <>
      <MainSidebar>
        <div className="gap-3 h-full overflow-hidden flex flex-col">
          <Text.SubtleSemi className="px-4 text-muted-foreground">Реестр объектов</Text.SubtleSemi>
          <Breadcrumb className="px-4">
            <BreadcrumbItem>Источники тепла</BreadcrumbItem>
          </Breadcrumb>
          <ScrollArea>
            {vm.heatSources.map((v) => (
              <React.Fragment key={v.id}>
                <HeatSourceCard data={v} />
                <Separator />
              </React.Fragment>
            ))}
          </ScrollArea>
        </div>
      </MainSidebar>
      <Outlet />
    </>
  )
});
