import { useRepositories } from "@/api/queries/Repositories";
import { Separator } from "@/shared/components/ui/separator";
import RepositoriesCard from "./RepositoriesCard";

const Repositories = () => {
  const { data } = useRepositories();
  const repositoriesData = data?.data;

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
