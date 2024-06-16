import { FCVM } from "@/utils/vm";
import { observer } from "mobx-react-lite";
import { IncidentsPageViewModel } from "./incidents.page.vm";
import { IconInput, Input } from "@/components/ui/input";
import { MenuIcon, SearchIcon } from "lucide-react";
import { Text } from "@/components/typography/Text";
import { IncidentsTabs } from "./incidents-tabs";
import { Drawer, DrawerContent, DrawerHeader, DrawerTrigger } from "@/components/ui/drawer";
import { Button } from "@/components/ui/button";
import React, { FC } from "react";
import { LayerSelect } from "../layoutProfileBar/components/LayerSelect";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Link } from "@tanstack/react-router";
import { HeatDistributorCard } from "@/components/cards/heat-distributor.card";
import { Separator } from "@/components/ui/separator";
import { LoadingWrapper } from "@/components/ui/loaders/LoadingWrapper";
import { ConsumerCard } from "@/components/cards/consumer.card";
import { IssueSelect } from "./IssueSelect";

const vm = IncidentsPageViewModel;

export const IncidentsContent: FC<{ isMobile?: boolean }> = observer(({ isMobile }) => {
  return (
    <>
      <div className="flex flex-col px-4 pb-4 gap-3">
        <IconInput
          containerClassName="flex-1"
          className="bg-background"
          rightIcon={<SearchIcon />}
          onChange={(v) => vm.setSearch(v.target.value)}
          placeholder="Введите данные UNOM"
        />
        {!isMobile && (
          <Text.UiMedium className="text-muted-foreground pt-2">Инциденты</Text.UiMedium>
        )}
        <IncidentsTabs />
        <IssueSelect
          value={vm.issueType}
          onChange={(v) => {
            vm.issueType = v;
          }}
        />
      </div>
      <ScrollArea>
        {vm.heatSources.map((v) => (
          <React.Fragment key={v.number}>
            <Link
              to="/incidents/$unom"
              params={{
                unom: v.unom.toString()
              }}>
              <HeatDistributorCard data={v.data} className="px-4" />
            </Link>
            <Separator />
          </React.Fragment>
        ))}
        {vm.consumers.map((v) => (
          <React.Fragment key={v.number}>
            <Link
              to="/incidents/$unom"
              params={{
                unom: v.unom.toString()
              }}>
              <ConsumerCard data={v.data} className="px-4" />
            </Link>
            <Separator />
          </React.Fragment>
        ))}
        {vm.loading && <LoadingWrapper />}
      </ScrollArea>
    </>
  );
});

export const IncidentsSidebar = observer(() => {
  return (
    <aside className="h-full md:flex flex-col w-96 bg-card py-4 overflow-hidden hidden">
      <IncidentsContent />
    </aside>
  );
});

export const IncidentsSidebarMobile = observer(() => {
  return (
    <Drawer open={vm.drawerOpen} onOpenChange={(v) => (vm.drawerOpen = v)}>
      <DrawerTrigger asChild>
        <Button
          variant="outline"
          size="icon"
          className="flex md:hidden absolute left-4 top-4 z-10 bg-card">
          <MenuIcon />
        </Button>
      </DrawerTrigger>
      <DrawerContent className="flex flex-col text-foreground pb-10 max-h-[90%] overflow-hidden">
        <ScrollArea className="flex-1">
          <DrawerHeader>Инциденты</DrawerHeader>
          <IncidentsContent isMobile />
        </ScrollArea>
      </DrawerContent>
    </Drawer>
  );
});
