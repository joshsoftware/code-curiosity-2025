import { useState, type FC } from "react";
import clsx from "clsx";
import { Button } from "@/shared/components/ui/button";
import { Card } from "@/shared/components/ui/card";
import ActivityCard from "@/shared/components/common/ActivityCard";
import { useRecentActivities } from "@/api/queries/RecentActivities";
import { Link } from "react-router-dom";
import { TrendingUp } from "lucide-react";

interface RecentActivitiesProps {
  className?: string;
}

const RecentActivities: FC<RecentActivitiesProps> = ({ className }) => {
  const [viewAll, setViewAll] = useState(false);

  const handleViewAll = () => {
    setViewAll(!viewAll);
  };

  const { data, isLoading } = useRecentActivities();
  const recentActivities = data?.data ?? [];
  const recentActivitiesData = viewAll
    ? recentActivities
    : recentActivities?.slice(0, 4);

  return (
    <Card
      className={clsx(
        "bg-cc-app-gray-background flex h-full w-full flex-col gap-2 overflow-auto border-none p-6 shadow-none",
        className
      )}
    >
      <div className="flex items-center justify-between">
        <p className="text-xl font-bold text-gray-900">Recent Activities</p>
        <Button
          variant="ghost"
          className="text-cc-app-blue hover:text-cc-app-blue cursor-pointer bg-transparent px-0 text-xs font-semibold hover:bg-transparent hover:underline"
          onClick={handleViewAll}
        >
          {viewAll ? "View Less" : "View All"}
        </Button>
      </div>

      {isLoading ? (
        <div className="flex h-full w-full items-center justify-center">
          <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
        </div>
      ) : recentActivitiesData?.length === 0 ? (
        <div className="flex h-full w-full flex-col items-center justify-center text-center">
          <TrendingUp className="mb-3 h-12 w-12 text-gray-400" />
          <p className="mb-2 text-lg font-medium text-gray-600">
            No recent activities found
          </p>
        </div>
      ) : (
        <div
          className={clsx(
            "flex h-full flex-col items-center justify-between",
            viewAll ? "no-scrollbar overflow-auto" : ""
          )}
        >
          {recentActivitiesData?.map((activity, index) => (
            <ActivityCard
              key={index}
              contributionType={activity.contributionType}
              repositoryName={activity.repoName}
              contributedAt={activity.contributedAt}
              balanceChange={activity.balanceChange}
              showLine={index < recentActivitiesData.length - 1}
            />
          ))}
          {!viewAll && (
            <div className="w-full">
              <Link
                to={""}
                className="text-cc-app-blue hover:text-cc-app-blue cursor-pointer bg-transparent text-xs font-semibold hover:bg-transparent hover:underline"
              >
                How does points work?
              </Link>
            </div>
          )}
        </div>
      )}
    </Card>
  );
};

export default RecentActivities;
