import { MainSidebar } from "@/widgets/layoutMainSidebar/main-sidebar.widget";
import { Outlet, createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_map/_with_filters")({
  component: () => (
    <>
      <MainSidebar>
        <>test</>
      </MainSidebar>
      <Outlet />
    </>
  )
});
