import { type FC } from "react";
import Coin from "@/shared/components/common/Coin";
import { Card } from "@/shared/components/ui/card";

interface OverviewCardProps {
  type: string;
  count: number;
  totalCoins: number;
}

const OverviewCard: FC<OverviewCardProps> = ({ type, count, totalCoins }) => {
  return (
    <Card className="border-none p-3 shadow-none">
      <div className="flex flex-col items-center justify-between gap-1">
        <p className="w-full text-sm font-medium text-gray-900">{type}</p>
        <div className="flex w-full flex-row items-center justify-between">
          <span className="text-cc-app-blue text-2xl font-semibold">
            {count}
          </span>
          <div className="flex items-center justify-between gap-2 rounded-sm bg-amber-100 p-1">
            <Coin />
            <span className="text-lg font-semibold text-orange-500">
              {totalCoins}
            </span>
          </div>
        </div>
      </div>
    </Card>
  );
};

export default OverviewCard;
