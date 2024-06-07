import { useTheme } from "../hoc/theme-provider";
import { Button } from "../ui/button";
import { MoonIcon, SunIcon } from "lucide-react";

export const ThemeSwitcher = () => {
  const theme = useTheme();

  return (
    <Button
      variant="outline"
      size="icon"
      onClick={() => theme.setTheme(theme.theme === "light" ? "dark" : "light")}>
      {theme.theme === "light" ? <MoonIcon /> : <SunIcon />}
    </Button>
  );
};
