import { ELEVATION } from "@/constants/elevation";
import { observer } from "mobx-react-lite";
import { UserNav } from "./components/UserNav";
import { ThemeSwitcher } from "./components/ThemeSwitcher";

export const ProfileBar = observer(() => {
  return (
    <div
      className="absolute top-4 right-4 flex gap-2 items-center"
      style={{ zIndex: ELEVATION.FILTERS }}>
      <ThemeSwitcher />
      <UserNav />
    </div>
  );
});
