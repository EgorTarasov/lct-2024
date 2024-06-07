import { MapLoading } from "@/widgets/map/map-loading";
import { SidebarContainer } from "@/widgets/sidebar/sidebar-container.widget";
import { SidebarContent, SidebarContext } from "@/widgets/sidebar/sidebar.context";
import { Sidebar } from "@/widgets/sidebar/sidebar.widget";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import React, { ReactNode, Suspense, useState } from "react";

const Map = React.lazy(() => import("@/widgets/map/map.widget"));

const Page = () => {
  const [content, setContent] = useState<SidebarContent | null>(null);
  const [isOpen, setIsOpen] = useState(false);

  return (
    <SidebarContext.Provider
      value={{
        isOpen,
        toggleSidebar: () => setIsOpen(!isOpen),
        content,
        setContent
      }}>
      <div className="h-full w-full relative">
        <SidebarContainer />
        <Suspense fallback={<MapLoading />}>
          <Map />
        </Suspense>
        <Outlet />
      </div>
    </SidebarContext.Provider>
  );
};

export const Route = createFileRoute("/_map")({
  component: () => <Page />
});
