import { ELEVATION } from "@/constants/elevation";
import { observer } from "mobx-react-lite";
import { UserNav } from "./components/UserNav";
import { ThemeSwitcher } from "./components/ThemeSwitcher";
import { LayerSelect } from "./components/LayerSelect";

export const ProfileBar = observer(() => {
  return (
    <div
      className="absolute top-4 right-4 flex gap-2 flex-col lg:flex-row items-end lg:items-start"
      style={{ zIndex: ELEVATION.FILTERS }}>
      <div className="flex gap-2">
        <LayerSelect />
        <ThemeSwitcher />
      </div>
      <UserNav />
    </div>
  );
});
