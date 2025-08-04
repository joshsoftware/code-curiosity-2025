import { useRepositoryContributors } from "@/api/queries/Contributors";
import { useParams } from "react-router-dom";
import ContributorsCard from "./ContributorsCard";
import { useState } from "react";
import { Button } from "@/shared/components/ui/button";
import clsx from "clsx";

const ContributorsList = () => {
  const [viewAll, setViewAll] = useState(false);

  const handleViewAll = () => {
    setViewAll(!viewAll);
  };

  const { repoid } = useParams();
  const repoId = Number(repoid);
  const { data, isLoading } = useRepositoryContributors(repoId);
  const contributors = data?.data ?? [];

  const contributorsData = viewAll ? contributors : contributors?.slice(0, 20);

  return (
    <div className="mx-auto flex h-full max-w-4xl flex-col gap-2 overflow-auto p-2 shadow-none">
      <div className="flex items-center justify-between">
        <p className="text-md font-bold text-cc-app-blue">
          Contributors {contributors.length}
        </p>
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
      ) : (
        <div
          className={clsx(
            "flex h-auto flex-1 flex-col gap-2 p-2",
            viewAll ? "no-scrollbar overflow-auto" : ""
          )}
        >
          <div className="*:data-[slot=avatar]:ring-background flex -space-x-4 *:data-[slot=avatar]:ring-2 *:data-[slot=avatar]:grayscale">
            {contributorsData?.map(contributor => (
              <ContributorsCard
                name={contributor.name}
                avatarUrl={contributor.avatarUrl}
                githubUrl={contributor.githubUrl}
                contributions={contributor.contributions}
              />
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default ContributorsList;
