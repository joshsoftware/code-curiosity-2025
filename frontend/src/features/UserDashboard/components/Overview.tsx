import { useState, type FC } from "react";
import clsx from "clsx";
import { Card } from "@/shared/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger
} from "@/shared/components/ui/dropdown-menu";
import OverviewCard from "@/features/UserDashboard/components/OverviewCard";
import { ChevronDown, TrendingUp } from "lucide-react";
import { format, subMonths } from "date-fns";
import { useOverview } from "@/api/queries/Overview";

interface OverviewProps {
  className?: string;
}

const getLastNMonths = (n: number) => {
  return Array.from({ length: n }).map((_, i) => {
    const date = subMonths(new Date(), i);
    return {
      label: format(date, "MMMM yyyy"),
      month: date.getMonth() + 1,
      year: date.getFullYear()
    };
  });
};

const Overview: FC<OverviewProps> = ({ className }) => {
  const monthOptions = getLastNMonths(3);
  const [selectedPeriod, setSelectedPeriod] = useState<{
    month: number;
    year: number;
  }>(monthOptions[0]);

  const { data, isLoading } = useOverview(selectedPeriod);
  const overview = data?.data ?? [];

  const overviewData = overview?.filter(data => {
    const date = new Date(data.month);
    return (
      date.getFullYear() === selectedPeriod.year &&
      date.getMonth() + 1 === selectedPeriod.month
    );
  });

  return (
    <Card
      className={clsx(
        "bg-cc-app-gray-background flex h-full w-full flex-col gap-2 overflow-auto border-none p-6 shadow-none",
        className
      )}
    >
      <div className="flex items-center justify-between">
        <p className="text-xl font-bold text-gray-900">Overview</p>
        <DropdownMenu>
          <DropdownMenuTrigger className="border-cc-app-mid-blue text-cc-app-blue hover:bg-cc-app-mid-blue/5 flex cursor-pointer items-center gap-2 rounded-sm border px-4 py-2 text-sm font-medium focus:outline-none">
            {monthOptions.find(
              opt =>
                opt.month === selectedPeriod.month &&
                opt.year === selectedPeriod.year
            )?.label ?? "Select Month"}
            <ChevronDown className="h-4 w-4" />
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-48 rounded-md border bg-white p-1">
            {monthOptions.map(option => (
              <DropdownMenuItem
                key={`${option.month}-${option.year}`}
                onClick={() =>
                  setSelectedPeriod({ month: option.month, year: option.year })
                }
                className={clsx(
                  "cursor-pointer rounded-sm px-3 py-2 text-sm hover:bg-gray-100",
                  selectedPeriod.month === option.month &&
                    selectedPeriod.year === option.year &&
                    "bg-cc-app-gray-background"
                )}
              >
                {option.label}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      {isLoading ? (
        <div className="flex h-full w-full items-center justify-center">
          <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
        </div>
      ) : overviewData?.length === 0 ? (
        <div className="flex h-full w-full flex-col items-center justify-center text-center">
          <TrendingUp className="mb-3 h-12 w-12 text-gray-400" />
          <p className="mb-2 text-lg font-medium text-gray-600">
            No overview data
          </p>
          <p className="text-sm text-gray-500">
            No activity found for the selected period.
            <br />
            Try selecting a different month or start contributing!
          </p>
        </div>
      ) : (
        <div className="no-scrollbar grid h-auto grid-cols-2 gap-2 overflow-auto">
          {overviewData?.map(user => (
            <OverviewCard
              key={user.type}
              type={user.type}
              count={user.count}
              totalCoins={user.totalCoins}
            />
          ))}
        </div>
      )}
    </Card>
  );
};

export default Overview;
