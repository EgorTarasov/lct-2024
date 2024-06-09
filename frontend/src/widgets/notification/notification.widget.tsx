import { Button } from "@/components/ui/button";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { ELEVATION } from "@/constants/elevation";
import { BellIcon } from "lucide-react";
import { observer } from "mobx-react-lite";

export const NotificationWidget = observer(() => {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button size="icon" variant="ghost">
          <BellIcon />
        </Button>
      </SheetTrigger>
      <SheetContent>test</SheetContent>
    </Sheet>
  );
});
