import { GrantsService } from "@/stores/grant.service";
import { MapLoading } from "@/widgets/map/map-loading";
import { Outlet, createFileRoute } from "@tanstack/react-router";
import React, { Suspense, useState } from "react";
import {
  SidebarContent,
  MainSidebarContext
} from "@/widgets/layoutMainSidebar/main-sidebar.context";
import { checkGrant } from "@/utils/check-grant";

const MainSidebarView = React.lazy(() =>
  import("@/widgets/layoutMainSidebar/main-sidebar.view").then((x) => ({
    default: x.MainSidebarView
  }))
);
const Map = React.lazy(() => import("@/widgets/map/map.widget"));
const BottomRightBar = React.lazy(() =>
  import("@/widgets/layoutBottomRight/bottom-right-bar").then((x) => ({
    default: x.BottomRightBar
  }))
);

const ProfileBar = React.lazy(() =>
  import("@/widgets/layoutProfileBar/profile-bar.widget").then((m) => ({ default: m.ProfileBar }))
);

const Page = () => {
  const [content, _setContent] = useState<SidebarContent | null>(null);
  const [secondaryContent, setSecondaryContent] = useState<SidebarContent | null>(null);
  const [isOpen, setIsOpen] = useState(true);

  const setContent = (v: SidebarContent | null) => {
    setIsOpen(true);
    _setContent(v);
  };

  return (
    <MainSidebarContext.Provider
      value={{
        isOpen,
        toggleSidebar: () => setIsOpen(!isOpen),
        content,
        setContent,
        secondaryContent,
        setSecondaryContent
      }}>
      <div className="h-full w-full relative">
        <Suspense fallback={<MapLoading />}>
          <Map />
        </Suspense>
        <Suspense fallback={null}>
          <MainSidebarView />
          <BottomRightBar />
          <ProfileBar />
        </Suspense>
        <Outlet />
      </div>
    </MainSidebarContext.Provider>
  );
};

export const Route = createFileRoute("/_map")({
  component: Page,
  beforeLoad: () => checkGrant(GrantsService.canReadMap)
});
