import { Text } from "@/components/typography/Text";
import { Button } from "@/components/ui/button";
import { AuthService } from "@/stores/auth.service";
import { isLoggedIn } from "@/utils/auth";
import { Link, useNavigate } from "@tanstack/react-router";
import { BellIcon, LogOutIcon } from "lucide-react";
import { observer } from "mobx-react-lite";
import { NotificationWidget } from "@/widgets/notification/notification.widget";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";

export const UserNav = observer(() => {
  const navigate = useNavigate();

  if (!isLoggedIn(AuthService.auth)) return null;

  const logout = () => {
    AuthService.logout();
    navigate({
      to: "/login"
    });
  };

  return (
    <div className="shadow py-2 px-3 bg-card text-card-foreground rounded-xl flex items-center gap-2">
      <Tooltip>
        <TooltipTrigger asChild>
          <Link
            to="/profile"
            className="size-10 flex shadow items-center justify-center rounded-full bg-accent text-accent-foreground">
            <Text.p>
              {[AuthService.auth.user.firstName ?? "К", AuthService.auth.user.lastName ?? "Д"]
                .map((name) => name[0].toUpperCase())
                .join("")}
            </Text.p>
          </Link>
        </TooltipTrigger>
        <TooltipContent>Профиль</TooltipContent>
      </Tooltip>
      <NotificationWidget />
      <Button size="icon" variant="ghost" onClick={logout}>
        <LogOutIcon />
      </Button>
    </div>
  );
});
