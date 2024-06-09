import { Button } from "@/components/ui/button";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { ELEVATION } from "@/constants/elevation";
import { NotificationStore } from "@/stores/notification.store";
import { cn } from "@/utils/cn";
import { BellIcon } from "lucide-react";
import { observer } from "mobx-react-lite";

const vm = NotificationStore;

export const NotificationWidget = observer(() => {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button size="icon" variant="ghost" className="relative">
          <BellIcon />
          {vm.hasNotifications && (
            <div
              className={cn(
                "absolute top-2 border border-card right-2 w-2 h-2 rounded-full bg-primary"
              )}
            />
          )}
        </Button>
      </SheetTrigger>
      <SheetContent>test</SheetContent>
    </Sheet>
  );
});
