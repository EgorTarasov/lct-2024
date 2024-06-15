import { HeatDistributorCard } from "@/components/cards/heat-distributor.card";
import { Text } from "@/components/typography/Text";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage
} from "@/components/ui/breadcrumb";
import { LoadingWrapper } from "@/components/ui/loaders/LoadingWrapper";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { MapStore } from "@/stores/map.store";
import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { PaginationWidget } from "@/widgets/pagination/pagination.widget";
import { Link, Outlet, createLazyFileRoute } from "@tanstack/react-router";
import { AnimatePresence, motion } from "framer-motion";
import { observer } from "mobx-react-lite";
import React from "react";

const vm = MapStore.heatSourceVm;
const transitionProps = {
  initial: { opacity: 0, translateY: 20 },
  animate: { opacity: 1, translateY: 0 },
  exit: { opacity: 0, translateY: 20 }
};

const Page = observer(() => {
  return (
    <>
      <MainSidebar>
        <div className="gap-3 h-full overflow-hidden flex flex-col">
          <Text.UiMedium className="px-4 text-muted-foreground">Реестр объектов</Text.UiMedium>
          <Breadcrumb className="px-4">
            <BreadcrumbList>
              <BreadcrumbItem>
                <BreadcrumbPage>Источники тепла</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
          <ScrollArea>
            {vm.items.map((v) => (
              <React.Fragment key={v.id}>
                <Link
                  to="/heat_distributor/$heatDistributorId"
                  params={{
                    heatDistributorId: v.id.toString()
                  }}>
                  <HeatDistributorCard data={v} />
                </Link>
                <Separator />
              </React.Fragment>
            ))}
            {vm.loading && <LoadingWrapper />}
          </ScrollArea>
          <AnimatePresence mode="popLayout" initial={false}>
            {vm.totalPages && (
              <motion.div className="mt-auto mb-4" {...transitionProps}>
                <PaginationWidget
                  currentPage={vm.page + 1}
                  onPageChange={(v) => (vm.page = v - 1)}
                  totalPages={vm.totalPages + 1}
                />
              </motion.div>
            )}
          </AnimatePresence>
        </div>
      </MainSidebar>
      <Outlet />
    </>
  );
});

export const Route = createLazyFileRoute("/_map/_heat_distributors")({
  component: Page
});
