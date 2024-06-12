import { ConsumerCard } from "@/components/cards/consumer.card";
import { Text } from "@/components/typography/Text";
import {
  Breadcrumb,
  BreadcrumbEllipsis,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator
} from "@/components/ui/breadcrumb";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { LoadingWrapper } from "@/components/ui/loaders/LoadingWrapper";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { HeatDistributor } from "@/types/heat.type";
import { cn } from "@/utils/cn";
import { FCVM } from "@/utils/vm";
import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { ConsumersViewModel } from "@/widgets/map/vm/consumers.vm";
import { PaginationWidget } from "@/widgets/pagination/pagination.widget";
import { Link, Outlet, createFileRoute } from "@tanstack/react-router";
import { ChevronDownIcon } from "lucide-react";
import { observer } from "mobx-react-lite";
import React, { FC, useEffect } from "react";

const PriorityDropdown: FCVM<ConsumersViewModel> = observer(({ vm }) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button variant="outline" className="bg-card flex gap-2 w-56 text-left px-3 mx-4">
        <Text.Subtle className="flex-1">
          {vm.showPriorityFirst ? "Сначала приоритетные" : "Сначала остывшие"}
        </Text.Subtle>
        <ChevronDownIcon className="size-4 text-muted-foreground" />
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent className="w-56">
      <DropdownMenuItem onClick={() => (vm.showPriorityFirst = true)}>
        Сначала приоритетные
      </DropdownMenuItem>
      <DropdownMenuItem onClick={() => (vm.showPriorityFirst = false)}>
        Сначала остывшие
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
));

const PageBreadcrumbs: FC<{
  heatDistributorId: string;
  heatSource: HeatDistributor.Item | null;
}> = (x) => (
  <Breadcrumb className="px-4">
    <BreadcrumbList>
      <BreadcrumbItem>
        <Link to="/">Источники тепла</Link>
      </BreadcrumbItem>
      <BreadcrumbSeparator />
      <DropdownMenu>
        <DropdownMenuTrigger>
          <BreadcrumbEllipsis className="size-4" />
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuItem asChild>
            <Link
              to="/heat_distributor/$heatDistributorId"
              params={{ heatDistributorId: x.heatDistributorId }}
              className={cn("w-full h-full cursor-pointer", !x.heatSource && "text-destructive")}>
              {x.heatSource?.number ?? "Неизвестный источник"}
            </Link>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <BreadcrumbSeparator />
      <BreadcrumbItem>
        <BreadcrumbPage className="cursor-default">Потребители тепла</BreadcrumbPage>
      </BreadcrumbItem>
    </BreadcrumbList>
  </Breadcrumb>
);

const Page = observer(() => {
  const heatDistributorId = Route.useParams().heatDistributorId;
  const heatSource = MapStore.heatSourceVm.items.find((v) => v.id.toString() === heatDistributorId);
  const vm = MapStore.consumersVm;

  useEffect(() => {
    if (heatSource) {
      MapStore.consumersVm = new ConsumersViewModel(heatSource);
    }
  }, [heatDistributorId]);

  return (
    <>
      <MainSidebar>
        <div className="gap-3 h-full overflow-hidden flex flex-col">
          <Text.UiMedium className="px-4 text-muted-foreground">Реестр объектов</Text.UiMedium>
          <PageBreadcrumbs heatDistributorId={heatDistributorId} heatSource={heatSource ?? null} />
          {heatSource && vm ? (
            <>
              <PriorityDropdown vm={vm} />
              <ScrollArea>
                {vm.items.map((v) => (
                  <React.Fragment key={v.id}>
                    <ConsumerCard heatDistributorId={heatDistributorId} data={v} />
                    <Separator />
                  </React.Fragment>
                ))}
                {vm.loading && <LoadingWrapper />}
              </ScrollArea>
              <PaginationWidget
                className="mt-auto mb-2"
                currentPage={1}
                onPageChange={() => void 0}
                totalPages={1}
              />
            </>
          ) : (
            <div className="flex flex-col gap-2 text-center">
              <Text.Large>Источник тепла не найден</Text.Large>
              <span className="text-muted-foreground">Возможно он был удалён</span>
            </div>
          )}
        </div>
      </MainSidebar>
      <Outlet />
    </>
  );
});

export const Route = createFileRoute(
  "/_map/_heat_distributors/heat_distributor/$heatDistributorId/_consumers"
)({
  component: Page
});
