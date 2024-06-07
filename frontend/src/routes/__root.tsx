import { ThemeProvider } from "@/components/hoc/theme-provider";
import { LoadingWrapper } from "@/components/ui/loaders/LoadingWrapper";
import { Toaster } from "@/components/ui/sonner";
import NotFoundPage from "@/pages/not-found.page";
import { createRootRoute, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";

export const Route = createRootRoute({
  component: () => (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Outlet />
      <Toaster richColors />
      <TanStackRouterDevtools />
    </ThemeProvider>
  ),
  pendingComponent: LoadingWrapper,
  notFoundComponent: NotFoundPage
});
