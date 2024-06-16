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
      <TableCaption>Готовые отчёты</TableCaption>
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
      <TableFooter>
        <TableRow>
          <TableCell colSpan={3}>Total</TableCell>
          <TableCell className="text-right">$2,500.00</TableCell>
        </TableRow>
      </TableFooter>
    </Table>
  );
}

const Page = observer(() => {
  return (
    <>
      <main className="mx-auto flex flex-col w-full max-w-screen-xl pt-20 h-full overflow-hidden">
        <Text.H4>Загруженные данные</Text.H4>
        <ScrollArea className="flex-1">
          <TableDemo />
        </ScrollArea>
        <div className="flex flex-col">test</div>
      </main>
    </>
  );
});

export const Route = createFileRoute("/_base/reports")({
  component: Page
});
