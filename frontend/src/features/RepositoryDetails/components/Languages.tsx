import { type FC } from "react";
import { useRepositoryLanguages } from "@/api/queries/Languages";
import { useParams } from "react-router-dom";
import LanguageCard from "./LanguagesCard";

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
