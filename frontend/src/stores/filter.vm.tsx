import { ComboboxMultipleProps } from "@/components/DropdownMultiple";
import { LocaleWithIcon } from "@/constants/map-filters";
import { makeAutoObservable } from "mobx";

export class Filter<T extends string> {
  constructor(
    readonly options: T[],
    readonly locale: Record<T, string>,
    private onChange?: () => void
  ) {
    makeAutoObservable(this);
  }

  private _value: T[] = [];

  get value() {
    return this._value;
  }

  get attributes(): ComboboxMultipleProps<T> {
    return {
      value: this.value,
      placeholder: this._value.length === 0 ? "Все" : undefined,
      options: this.options,
      compare: (v: T) => this.locale[v],
      onChange: (v: T[]) => {
        if (v.length === this.options.length) {
          v = [];
        }
        this._value = v;

        this.onChange?.();
      },
      render: (v: T) => this.locale[v]
    };
  }

  reset() {
    this._value = [];
  }
}