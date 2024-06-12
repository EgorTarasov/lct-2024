import PriorityHighIcon from "@/assets/priority-high.svg";
import PriorityMediumIcon from "@/assets/priority-medium.svg";
import PriorityLowIcon from "@/assets/priority-low.svg";

export enum Priority {
  HIGH = "HIGH",
  MEDIUM = "MEDIUM",
  LOW = "LOW"
}

export const PriorityLocaleMap = {
  [Priority.HIGH]: {
    locale: "Высокий",
    alternateLocale: "Высокая",
    color: "#EF4444",
    backgroundColor: "rgba(247, 119, 0, 0.2)",
    icon: <PriorityHighIcon />
  },
  [Priority.MEDIUM]: {
    locale: "Средний",
    alternateLocale: "Средняя",
    color: "#F78500",
    backgroundColor: "rgba(247, 119, 0, 0.2)",
    icon: <PriorityMediumIcon />
  },
  [Priority.LOW]: {
    locale: "Низкий",
    alternateLocale: "Низкая",
    color: "#FFA903",
    backgroundColor: "rgba(247, 119, 0, 0.2)",
    icon: <PriorityLowIcon />
  }
};
