import { buttonVariants } from "@/components/ui/button";
import { Link, createLazyFileRoute } from "@tanstack/react-router";

export const Route = createLazyFileRoute("/_base/profile")({
  component: () => (
    <div className="w-full h-96 flex items-center justify-center text-2xl gap-8 mx-auto text-center flex-col">
      Тут будет информация о вашем профиле
      <Link
        to="/"
        className={buttonVariants({
          variant: "default"
        })}>
        Перейти на карту
      </Link>
    </div>
  )
});
