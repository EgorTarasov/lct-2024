import {
  Combobox,
  Transition,
  ComboboxButton,
  ComboboxOptions,
  ComboboxOption,
  Label,
  ComboboxInput
} from "@headlessui/react";
import { Fragment, useEffect, useState } from "react";
import { observer } from "mobx-react-lite";
import { CheckIcon, ChevronDownIcon } from "@radix-ui/react-icons";
import { cn } from "@/utils/cn";

interface ComboboxMultipleProps<T> {
  value: T[];
  onChange: (value: T[]) => void;
  options: readonly T[];
  compare: (value: T) => string;
  render: (value: T) => React.ReactNode;
  label?: string;
  placeholder?: string;
}

const DropdownMultiple = observer(<T,>(p: ComboboxMultipleProps<T>) => {
  const [query, setQuery] = useState("");
  const [inputFocused, setInputFocused] = useState(false);

  const filteredOptions =
    query === ""
      ? p.options
      : p.options.filter((option) =>
          p
            .compare(option)
            .toLowerCase()
            .replace(/\s+/g, "")
            .includes(query.toLowerCase().replace(/\s+/g, ""))
        );
  const placeholder = p.value.map((v) => p.compare(v)).join(", ");

  useEffect(() => {
    const closeOnEscape = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        if (inputFocused) {
          event.stopPropagation();
          setInputFocused(false);
        }
      }
    };

    document.addEventListener("keydown", closeOnEscape);

    return () => {
      document.removeEventListener("keydown", closeOnEscape);
    };
  }, [inputFocused]);

  return (
    <Combobox value={p.value} multiple onChange={p.onChange}>
      <div className="relative text-sm">
        {p.label && <Label className="text-text-primary/60 text-base">{p.label}</Label>}
        <div className={cn("relative h-fit flex items-center w-full", p.label && "mt-3")}>
          <ComboboxInput
            className="whitespace-nowrap w-full cursor-pointer pr-8 text-ellipsis border-2 border-text-primary/20 rounded-lg p-3"
            placeholder={placeholder}
            onFocus={(e) => {
              e.preventDefault();
              setQuery("");
              setInputFocused(true);
            }}
            onBlur={(e) => {
              if (e.relatedTarget instanceof HTMLLIElement) return;

              setInputFocused(false);
            }}
            displayValue={(value: T[]) => (inputFocused ? "" : placeholder)}
            onChange={(event) => setQuery(event.target.value)}
          />
          <ComboboxButton className="h-5 w-5 absolute right-2 text-text-primary/60">
            <ChevronDownIcon />
          </ComboboxButton>
        </div>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100 delay-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
          show={inputFocused}
          afterLeave={() => setQuery("")}>
          <ComboboxOptions
            className="absolute z-10 mt-1 max-h-60 w-full border border-border overflow-auto rounded-xl py-2 bg-white"
            style={{
              scrollbarWidth: "thin"
            }}>
            {filteredOptions.length === 0 && query !== "" ? (
              <div className="px-4 py-2 text-text-secondary">Ничего не найдено</div>
            ) : (
              filteredOptions.map((option, index) => (
                <ComboboxOption
                  key={index}
                  value={option}
                  className={({ active }) =>
                    cn(
                      "p-2 cursor-pointer flex justify-between hover:bg-text-primary/5 items-center",
                      active && "bg-primary/5"
                    )
                  }>
                  {({ selected, active }) => (
                    <>
                      <span>{p.render(option)}</span>
                      {/* {active && "актив"} */}
                      {selected && <CheckIcon className="w-5 h-5" />}
                    </>
                  )}
                </ComboboxOption>
              ))
            )}
            {}
          </ComboboxOptions>
        </Transition>
      </div>
    </Combobox>
  );
});

export default DropdownMultiple;
