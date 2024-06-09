import { GrantsService } from "@/stores/grant.service";
import { MapLoading } from "@/widgets/map/map-loading";
import { MainSidebarView } from "@/widgets/layoutMainSidebar/sidebar-container.widget";
import {
  MainSidebarContent,
  MainSidebarContext
} from "@/widgets/layoutMainSidebar/sidebar.context";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import React, { Suspense, useState } from "react";
import { WeatherWidget } from "@/widgets/weather/weather.widget";

const Map = React.lazy(() => import("@/widgets/map/map.widget"));

const Page = () => {
  const [content, _setContent] = useState<MainSidebarContent | null>(null);
  const [isOpen, setIsOpen] = useState(false);

  const setContent = (v: MainSidebarContent | null) => {
    setIsOpen(true);
    _setContent(v);
  };

  return (
    <MainSidebarContext.Provider
      value={{
        isOpen,
        toggleSidebar: () => setIsOpen(!isOpen),
        content,
        setContent
      }}>
      <div className="h-full w-full relative">
        <MainSidebarView />
        <Suspense fallback={<MapLoading />}>
          <Map />
        </Suspense>
        <WeatherWidget />
        <Outlet />
      </div>
    </MainSidebarContext.Provider>
  );
};

export const Route = createFileRoute("/_map")({
  component: () => <Page />,
  beforeLoad: async (x) => {
    if (GrantsService.canReadMap) {
      return;
    }

    x.navigate({
      to: "/login"
    });
  }
});
