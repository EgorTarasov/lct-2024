import { ELEVATION } from "@/constants/elevation";
import { observer } from "mobx-react-lite";
import { ThemeSwitcher } from "../layoutProfileBar/components/ThemeSwitcher";
import { UserNav } from "../layoutProfileBar/components/UserNav";
import { Link } from "@tanstack/react-router";
import { buttonVariants } from "@/components/ui/button";
import { ChevronLeftIcon } from "lucide-react";
import { FC } from "react";
import { cn } from "@/utils/cn";
import { IncidentsPageViewModel } from "./incidents.page.vm";

export const IncidentsOverlay: FC<{ hasSidebar?: boolean; vm?: IncidentsPageViewModel }> = observer(
  (x) => {
    return (
      <>
        <div
          className={cn(
            "absolute top-4 left-4 ml-4 flex pointer-events-none",
            x.hasSidebar && "md:left-96"
          )}
          style={{ zIndex: ELEVATION.FILTERS }}>
          <Link
            to="/"
            className={cn(
              buttonVariants({ variant: "secondary" }),
              "hidden md:flex pointer-events-auto"
            )}>
            <ChevronLeftIcon />
            Вернуться к карте
          </Link>
        </div>
        <div className="absolute top-4 right-4 flex gap-2" style={{ zIndex: ELEVATION.FILTERS }}>
          <div className="flex gap-2">
            <ThemeSwitcher />
            <UserNav />
          </div>
        </div>
      </>
    );
  }
);
