import { useState, type FC } from "react";
import clsx from "clsx";
import { Button } from "@/shared/components/ui/button";
import { Card } from "@/shared/components/ui/card";
import ActivityCard from "@/shared/components/common/ActivityCard";
import {  useParams } from "react-router-dom";
import { TrendingUp } from "lucide-react";
import { useRepositoryActivities } from "@/api/queries/RepostoryActivities";
import CoinsInfo from "@/shared/components/common/CoinsInfo";

interface RepositoryActivitiesProps {
  className?: string;
}

const RepositoryActivities: FC<RepositoryActivitiesProps> = ({ className }) => {
  const [viewAll, setViewAll] = useState(false);

  const handleViewAll = () => {
    setViewAll(!viewAll);
  };

  const { repoid } = useParams();
  const repoId = Number(repoid);
  const { data, isLoading } = useRepositoryActivities(repoId);
  const repositoryActivities = data?.data ?? [];
  const repositoryActivitiesData = viewAll
    ? repositoryActivities
    : repositoryActivities?.slice(0, 4);

  let content;
  if (isLoading) {
    content = (
      <div className="flex h-full w-full items-center justify-center">
        <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
      </div>
    );
  } else if (repositoryActivitiesData?.length === 0) {
    content = (
      <div className="flex h-full w-full flex-col items-center justify-center text-center">
        <TrendingUp className="mb-3 h-12 w-12 text-gray-400" />
        <p className="mb-2 text-lg font-medium text-gray-600">
          No recent activities found
        </p>
      </div>
    );
  } else {
    content = (
      <div
        className={clsx(
          "flex h-full flex-col items-center justify-between pt-2",
          viewAll ? "no-scrollbar overflow-auto" : ""
        )}
      >
        {repositoryActivitiesData?.map((activity, index) => (
          <ActivityCard
            key={activity.id ?? `${activity.contributionType}-${activity.contributedAt}-${index}`}
            contributionType={activity.contributionType}
            contributedAt={activity.contributedAt}
            balanceChange={activity.balanceChange}
            showLine={index < repositoryActivities.length - 1}
            isRepositoryActivity={true}
          />
        ))}
         {!viewAll && (
            <div className="w-full text-right">
              <CoinsInfo />
            </div>
          )}
      </div>
    );
  }

  return (
    <Card
      className={clsx(
        "flex h-full w-full flex-col gap-2 overflow-auto border border-gray-300 p-5 shadow-none",
        className
      )}
    >
      <div className="flex items-center justify-between">
        <p className="text-md font-bold text-gray-900">Recent Activities</p>
        <Button
          variant="ghost"
          className="text-cc-app-blue hover:text-cc-app-blue cursor-pointer bg-transparent px-0 text-xs font-semibold hover:bg-transparent hover:underline"
          onClick={handleViewAll}
        >
          {viewAll ? "View Less" : "View All"}
        </Button>
      </div>
      {content}
    </Card>
  );
};

export default RepositoryActivities;
