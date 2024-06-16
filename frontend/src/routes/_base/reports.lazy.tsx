import { Text } from "@/components/typography/Text";
import { createFileRoute } from "@tanstack/react-router";
import { observer } from "mobx-react-lite";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow
} from "@/components/ui/table";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Button, buttonVariants } from "@/components/ui/button";
import { PlusIcon } from "lucide-react";
import { Input } from "@/components/ui/input";
import { cn } from "@/utils/cn";

const invoices = [
  {
    invoice: "INV001",
    paymentStatus: "Paid",
    totalAmount: "$250.00",
    paymentMethod: "Credit Card"
  },
  {
    invoice: "INV002",
    paymentStatus: "Pending",
    totalAmount: "$150.00",
    paymentMethod: "PayPal"
  },
  {
    invoice: "INV003",
    paymentStatus: "Unpaid",
    totalAmount: "$350.00",
    paymentMethod: "Bank Transfer"
  },
  {
    invoice: "INV004",
    paymentStatus: "Paid",
    totalAmount: "$450.00",
    paymentMethod: "Credit Card"
  },
  {
    invoice: "INV005",
    paymentStatus: "Paid",
    totalAmount: "$550.00",
    paymentMethod: "PayPal"
  },
  {
    invoice: "INV006",
    paymentStatus: "Pending",
    totalAmount: "$200.00",
    paymentMethod: "Bank Transfer"
  },
  {
    invoice: "INV007",
    paymentStatus: "Unpaid",
    totalAmount: "$300.00",
    paymentMethod: "Credit Card"
  }
];

export function TableDemo() {
  return (
    <Table>
      <TableHeader>
        <TableRow className="*:whitespace-nowrap">
          <TableHead className="w-full">Название файла</TableHead>
          <TableHead>Дата добавления</TableHead>
          <TableHead>Действия с файлом</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {invoices.map((invoice) => (
          <TableRow key={invoice.invoice}>
            <TableCell>{invoice.invoice}</TableCell>
            <TableCell>{invoice.paymentStatus}</TableCell>
            <TableCell>{invoice.paymentMethod}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}

const Page = observer(() => {
  return (
    <>
      <main className="mx-auto flex flex-col w-full max-w-screen-xl pt-20 h-full overflow-hidden px-4">
        <Text.H4>Загруженные данные</Text.H4>
        <ScrollArea className="flex-1 pt-4">
          <TableDemo />
          <div className="flex pt-4">
            <div className={cn(buttonVariants({}), "relative ml-auto")}>
              <input type="file" className="opacity-0 absolute inset-0" />
              <PlusIcon />
              <span>Загрузить файл</span>
            </div>
          </div>
          <div className="pb-8" />
        </ScrollArea>
      </main>
    </>
  );
});

export const Route = createFileRoute("/_base/reports")({
  component: Page
});
