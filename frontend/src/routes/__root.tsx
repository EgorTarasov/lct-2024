import { ThemeProvider } from "@/components/hoc/theme-provider";
import { LoadingWrapper } from "@/components/ui/loaders/LoadingWrapper";
import { TooltipProvider } from "@/components/ui/tooltip";
import { ELEVATION } from "@/constants/elevation";
import NotFoundPage from "@/pages/not-found.page";
import { AuthService } from "@/stores/auth.service";
import { createRootRoute, useMatch, useMatches } from "@tanstack/react-router";
import React from "react";

const Toaster = React.lazy(() =>
  import("@/components/ui/sonner").then((m) => ({ default: m.Toaster }))
);

const AnimatedOutlet = React.lazy(() =>
  import("@/components/router/animated-outlet").then((m) => ({ default: m.AnimatedOutlet }))
);

const AnimatePresence = React.lazy(() =>
  import("framer-motion").then((m) => ({ default: m.AnimatePresence }))
);

const Page = () => {
  const matches = useMatches();
  const match = useMatch({ strict: false });
  const nextMatchIndex = matches.findIndex((d) => d.id === match.id) + 1;
  const nextMatch = matches[nextMatchIndex];

  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <TooltipProvider>
        <React.Suspense
          fallback={
            <div
              className="absolute inset-0 flex justify-center items-center"
              style={{ zIndex: ELEVATION.SIDEBAR }}>
              <LoadingWrapper />
            </div>
          }>
          <AnimatePresence mode="popLayout">
            <AnimatedOutlet key={nextMatch.id} />
          </AnimatePresence>
          <Toaster richColors />
        </React.Suspense>
      </TooltipProvider>
    </ThemeProvider>
  );
};

export const Route = createRootRoute({
  component: Page,
  pendingComponent: LoadingWrapper,
  notFoundComponent: NotFoundPage,
  beforeLoad: () => AuthService.waitInit()
});
