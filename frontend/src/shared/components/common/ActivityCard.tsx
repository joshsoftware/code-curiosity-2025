import type { FC } from "react";
import Coin from "@/shared/components/common/Coin";
import { format } from "date-fns";

interface ActivityCardProps {
  contributionType: string;
  repositoryName?: string;
  contributedAt: string;
  balanceChange: number;
  showLine: boolean;
  isRepositoryActivity?: boolean;
}

const ActivityCard: FC<ActivityCardProps> = ({
  contributionType,
  repositoryName,
  contributedAt,
  balanceChange,
  showLine = true,
  isRepositoryActivity
}) => {
  return (
    <div className="relative flex h-full w-full items-start">
      {showLine && (
        <div className="bg-cc-app-blue absolute top-6 left-2 h-full w-0.5"></div>
      )}

      <div className="bg-cc-app-orange relative z-10 mt-1 h-4 w-4 flex-shrink-0 rounded-full"></div>

      <div className="ml-4 flex w-full items-start justify-between pb-4">
        <div className="flex-1">
          <div className="text-sm font-semibold text-gray-900">
            {contributionType.replace(/([A-Z])/g, " $1")}
          </div>
          {isRepositoryActivity ? null : (
            <div className="text-cc-app-mid-blue mt-1 text-xs">
              Contributed to &lt;{repositoryName}&gt;
            </div>
          )}

          <div className="text-cc-app-mid-blue mt-1 text-xs">
            Contributed on {format(new Date(contributedAt), "MMM d yyyy")}
          </div>
        </div>

        {balanceChange && (
          <div
            className={`flex items-center gap-1 text-lg font-bold ${
              balanceChange < 0 ? "text-red-500" : "text-cc-app-orange"
            }`}
          >
            <Coin />
            {balanceChange < 0 ? `-${Math.abs(balanceChange)}` : balanceChange}
          </div>
        )}
      </div>
    </div>
  );
};

export default ActivityCard;
