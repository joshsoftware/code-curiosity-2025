import type { FC } from "react";

import Coin from "@/shared/components/common/Coin";
import { Card, CardContent } from "@/shared/components/ui/card";

interface LeaderboardCardProps {
  rank: number;
  username: string;
  repositories: number;
  balance: number;
}

const LeaderboardCard: FC<LeaderboardCardProps> = ({
  rank,
  username,
  repositories,
  balance
}) => {
  return (
    <div className="flex justify-end">
      <Card className="w-[95%] border-none bg-white p-1 shadow-none">
        <CardContent className="relative flex items-center justify-between bg-white p-2">
          <div className="border-cc-app-mid-blue text-cc-app-mid-blue absolute -left-5 flex h-10 w-10 items-center justify-center rounded-full border-4 bg-white text-sm font-bold">
            {rank}
          </div>
          <div className="ml-6 space-y-1">
            <div className="text-sm font-medium">{username}</div>
            <div className="text-cc-app-mid-blue text-xs">
              Contributed to{" "}
              <span className="font-semibold">
                {repositories}{" "}
                {repositories > 1 ? "Repositories" : "Repository"}
              </span>
            </div>
          </div>
          <div className="text-cc-app-orange flex items-center gap-2 text-lg font-bold">
            <Coin />
            {balance}
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default LeaderboardCard;