import { observer } from "mobx-react-lite";
import { useSidebar } from "./main-sidebar.context";
import { twMerge } from "tailwind-merge";
import { ELEVATION } from "@/constants/elevation";
import { cn } from "@/utils/cn";
import { Button } from "@/components/ui/button";
import { ChevronLeft, CrossIcon, FilterIcon, SearchIcon, XIcon } from "lucide-react";
import { useCallback, useEffect, useState } from "react";
import { IconInput, Input } from "@/components/ui/input";
import { MainSidebarFilters } from "./filters";
import { AnimatePresence } from "framer-motion";
import { useIsPresent, motion } from "framer-motion";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { throttle } from "@/utils/debounce";
import { SecondarySidebarView } from "./SecondarySidebar/secondary-sidebar.view";

const transitionProps = {
  initial: { opacity: 0, translateY: 20 },
  animate: { opacity: 1, translateY: 0 },
  exit: { opacity: 0, translateY: 20 }
};

export const MainSidebarView = observer(() => {
  const [filtersOpen, _setFiltersOpen] = useState(false);
  const ctx = useSidebar();

  const setFiltersOpen = useCallback(
    throttle((v: boolean) => {
      _setFiltersOpen(v);
    }, 350),
    []
  );

  return (
    <>
      <SecondarySidebarView />
      <div
        className="flex absolute gap-2 top-4 left-4 w-[412px]"
        style={{ zIndex: ELEVATION.SEARCHBAR }}>
        <IconInput
          containerClassName="flex-1"
          className={ctx.isOpen ? "bg-background shadow-none" : "bg-card"}
          rightIcon={<SearchIcon />}
          placeholder="Введите данные объекта"
        />
        <Button
          onClick={() => {
            if (filtersOpen) {
              setFiltersOpen(false);
            } else {
              if (!ctx.isOpen) {
                ctx.toggleSidebar();
              }
            }
            setFiltersOpen(!filtersOpen);
          }}
          variant={filtersOpen ? "destructive" : "outline"}
          size="icon"
          className={cn(
            filtersOpen ? "bg-destructive" : "bg-card",
            ctx.isOpen ? "shadow-none" : "shadow-sm"
          )}>
          {filtersOpen ? <XIcon /> : <FilterIcon />}
        </Button>
        <Tooltip>
          <TooltipTrigger asChild>
            <Button
              variant="outline"
              size="icon"
              className="ml-4 bg-card"
              onClick={() => {
                if (filtersOpen) {
                  setFiltersOpen(false);
                }
                ctx.toggleSidebar();
              }}>
              <ChevronLeft className={cn(!ctx.isOpen && "rotate-180")} />
            </Button>
          </TooltipTrigger>
          <TooltipContent side="right">
            {ctx.isOpen ? "Скрыть" : "Показать"} боковую панель
          </TooltipContent>
        </Tooltip>
      </div>
      <aside
        className={cn(
          "absolute transition-transform flex left-0 bottom-0 top-0 -translate-x-96 overflow-hidden",
          (ctx.isOpen || filtersOpen) && "translate-x-0"
        )}
        style={{ zIndex: ELEVATION.SIDEBAR }}>
        <div className={"w-96 flex h-full bg-card text-card-foreground shadow-md pt-[72px]"}>
          <AnimatePresence mode="popLayout">
            {filtersOpen ? (
              <motion.div className="w-full p-4" key="filters" {...transitionProps}>
                <MainSidebarFilters />
              </motion.div>
            ) : (
              <motion.div className={"w-full"} key="content" {...transitionProps}>
                {ctx.content?.content}
              </motion.div>
            )}
          </AnimatePresence>
        </div>
      </aside>
    </>
  );
});
