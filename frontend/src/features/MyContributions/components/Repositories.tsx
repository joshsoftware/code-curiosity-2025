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
      {repositoriesData?.map(repo => (
        <>
          <RepositoriesCard
            key={repo.id}
            id={repo.id}
            name={repo.repoName}
            languages={repo.languages}
            description={repo.description}
            updatedOn={repo.updateDate}
            coins={repo.totalCoinsEarned}
          />
          <Separator />
        </>
      ))}
    </div>
  );
};

export default Repositories;
