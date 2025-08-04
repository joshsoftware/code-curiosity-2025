import RepositoryCard from "./RepositoryCard";
import { useRepository } from "@/api/queries/Repository";
import { useParams } from "react-router-dom";

const Repository = () => {
  const { repoid } = useParams();
  const repoId = Number(repoid);
  const { data } = useRepository(repoId);
  const repo = data?.data;

  return (
    <div className=" mx-auto p-8">
      <RepositoryCard
        key={repo?.id}
        name={repo?.repoName || ""}
        languages={repo?.languages || []}
        description={repo?.description || ""}
        updatedOn={repo?.updateDate || ""}
        owner={repo?.ownerName || ""}
        repoUrl={repo?.repoUrl || ""}
      />
    </div>
  );
};

export default Repository;
