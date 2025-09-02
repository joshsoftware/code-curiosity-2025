import { useRepositories } from "@/api/queries/Repositories";
import { Separator } from "@/shared/components/ui/separator";
import RepositoriesCard from "./RepositoriesCard";

const Repositories = () => {
  const { data, isLoading } = useRepositories();
  const repositoriesData = data?.data;
  if (isLoading)
    return (
      <div className="flex h-full w-full items-center justify-center">
        <div className="border-t-cc-app-blue h-12 w-12 animate-spin rounded-full border-4 border-gray-200"></div>
      </div>
    );
  return (
    <div className="mx-auto max-w-4xl p-6">
      {repositoriesData?.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-20 text-center">
          <p className="text-lg font-semibold text-gray-700">
            You haven't contributed to any repositories yet
          </p>
          <p className="mt-2 text-sm text-gray-500">
            Start contributing to earn coins and track your progress 
          </p>
        </div>
      ) : (
        repositoriesData?.map(repo => (
          <div key={repo.id}>
            <RepositoriesCard
              id={repo.id}
              name={repo.repoName}
              languages={repo.languages}
              description={repo.description}
              updatedOn={repo.updateDate}
              coins={repo.totalCoinsEarned}
            />
            <Separator />
          </div>
        ))
      )}
    </div>
  );
};

export default Repositories;
