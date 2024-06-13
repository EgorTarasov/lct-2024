import { ELEVATION } from "@/constants/elevation";
import { observer } from "mobx-react-lite";
import { ThemeSwitcher } from "../layoutProfileBar/components/ThemeSwitcher";
import { UserNav } from "../layoutProfileBar/components/UserNav";
import { Link } from "@tanstack/react-router";
import { buttonVariants } from "@/components/ui/button";
import { ChevronLeftIcon } from "lucide-react";
import { FC } from "react";
import { cn } from "@/utils/cn";

export const IncidentsOverlay: FC<{ hasSidebar?: boolean }> = observer((x) => {
  return (
    <>
      <div
        className={cn("absolute top-4 left-4 ml-4 flex", x.hasSidebar && "left-96")}
        style={{ zIndex: ELEVATION.FILTERS }}>
        <Link to="/" className={buttonVariants({ variant: "secondary" })}>
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
});
