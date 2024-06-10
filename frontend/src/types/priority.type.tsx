import PriorityHighIcon from "@/assets/priority-high.svg";
import PriorityMediumIcon from "@/assets/priority-medium.svg";
import PriorityLowIcon from "@/assets/priority-low.svg";

export namespace Priority {
  export enum Item {
    HIGH = "HIGH",
    MEDIUM = "MEDIUM",
    LOW = "LOW"
  }

  export const ItemMap = {
    [Item.HIGH]: {
      locale: "Высокий",
      alternateLocale: "Высокая",
      color: "#EF4444",
      backgroundColor: "rgba(247, 119, 0, 0.2)",
      icon: <PriorityHighIcon />
    },
    [Item.MEDIUM]: {
      locale: "Средний",
      alternateLocale: "Средняя",
      color: "#F78500",
      backgroundColor: "rgba(247, 119, 0, 0.2)",
      icon: <PriorityMediumIcon />
    },
    [Item.LOW]: {
      locale: "Низкий",
      alternateLocale: "Низкая",
      color: "#FFA903",
      backgroundColor: "rgba(247, 119, 0, 0.2)",
      icon: <PriorityLowIcon />
    }
  };
}
