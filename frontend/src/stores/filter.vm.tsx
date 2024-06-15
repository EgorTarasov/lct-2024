import { ComboboxMultipleProps } from "@/components/DropdownMultiple";
import { makeAutoObservable } from "mobx";

export class Filter<T extends string> {
  constructor(
    public readonly name: string,
    public readonly options: T[],
    public readonly locale?: Record<T, string>
  ) {
    makeAutoObservable(this);
  }

  private _value: T[] = [];

  get attributes(): ComboboxMultipleProps<T> {
    return {
      value: this._value,
      placeholder: this._value.length === 0 ? "Все" : undefined,
      options: this.options,
      label: this.name,
      compare: (v: T) => (this.locale ? this.locale[v] : v),
      onChange: (v: T[]) => {
        if (v.length === this.options.length) {
          v = [];
        }
        this._value = v;
      },
      render: (v: T) => (this.locale ? this.locale[v] : v)
    };
  }

  get values(): T[] {
    return this._value.length === 0 ? this.options : this._value;
  }

  reset() {
    this._value = [];
  }
}
