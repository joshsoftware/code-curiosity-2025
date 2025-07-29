import { type FC } from "react";
import LanguageCard from "../../RepositoryDetails.tsx/components/LanguagesCard";
import { useRepositoryLanguages } from "@/api/queries/Languages";
import { useParams } from "react-router-dom";

interface LanguagesProps {
  className?: string;
}

const Languages: FC<LanguagesProps> = ({ className }) => {
  const { repoid } = useParams();
  const repoId = Number(repoid);
  const { data } = useRepositoryLanguages(repoId);

  const languagesData = data?.data;
  return (
    <LanguageCard
      title="Languages"
      languages={languagesData ?? []}
      className={className}
    />
  );
};

export default Languages;
export type { LanguagesProps };
