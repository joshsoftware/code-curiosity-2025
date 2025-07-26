import { useState, type FC } from "react";
import clsx from "clsx";

import { Button } from "@/shared/components/ui/button";
import { Card } from "@/shared/components/ui/card";
import LeaderboardCard from "@/features/UserDashboard/components/LeaderboardCard";
import { useCurrentUserRank, useLeaderboard } from "@/api/queries/Leaderboard";
import { TrendingUp } from "lucide-react";

interface LeaderboardProps {
  className?: string;
}

const Leaderboard: FC<LeaderboardProps> = ({ className }) => {
  const [viewAll, setViewAll] = useState(false);

  const handleViewAll = () => {
    setViewAll(!viewAll);
  };

  const { data, isLoading } = useLeaderboard();
  const leaderboard = data?.data ?? [];

  const leaderboardData = viewAll ? leaderboard : leaderboard?.slice(0, 10);

  const { data: userData } = useCurrentUserRank();
  const currentUser = userData?.data;

  return (
    <Card
      className={clsx(
        "bg-cc-app-gray-background flex h-full w-full flex-col gap-2 overflow-auto border-none p-6 shadow-none",
        className
      )}
    >
      <div className="flex items-center justify-between">
        <p className="text-xl font-bold text-gray-900">Leader Board</p>
        <Button
          variant="ghost"
          className="text-cc-app-blue hover:text-cc-app-blue cursor-pointer bg-transparent text-xs font-semibold hover:bg-transparent hover:underline"
          onClick={handleViewAll}
        >
          {viewAll ? "View Less" : "View All"}
        </Button>
      </div>

      {isLoading ? (
        <div className="flex h-full w-full items-center justify-center">
          <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
        </div>
      ) : leaderboardData?.length === 0 ? (
        <div className="flex h-full w-full flex-col items-center justify-center text-center">
          <TrendingUp className="mb-3 h-12 w-12 text-gray-400" />
          <p className="mb-2 text-lg font-medium text-gray-600">
            No leaderboard data
          </p>
        </div>
      ) : (
        <>
          <div className="no-scrollbar flex h-auto flex-1 flex-col gap-2 overflow-auto">
            {leaderboardData?.map(user => (
              <LeaderboardCard
                key={user.id}
                rank={user.rank}
                username={user.githubUsername}
                repositories={user.contributedReposCount}
                balance={user.currentBalance}
              />
            ))}
          </div>
          {!viewAll && (
            <div className="bg-cc-app-blue mt-auto rounded-xl p-2">
              <LeaderboardCard
                rank={currentUser?.rank || 0}
                username={currentUser?.githubUsername || ""}
                repositories={currentUser?.contributedReposCount || 0}
                balance={currentUser?.currentBalance || 0}
              />
            </div>
          )}
        </>
      )}
    </Card>
  );
};

export default Leaderboard;
