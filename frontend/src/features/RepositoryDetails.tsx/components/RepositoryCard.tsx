import { LangColor } from "@/shared/constants/constants";
import { ExternalLink } from "lucide-react";
import type { FC } from "react";

interface RepositoriesCardProps {
  name: string;
  languages: string[];
  description: string;
  updatedOn: string;
  owner: string;
  repoUrl: string;
}

const RepositoryCard: FC<RepositoriesCardProps> = ({
  name,
  languages,
  description,
  owner,
  repoUrl
}) => {
  return (
    <div>
      <div className="flex items-start justify-between">
        <div className="flex items-center gap-2">
          {name}
          <a
            href={repoUrl}
            target="_blank"
            rel="noopener noreferrer"
            className="text-xl font-medium text-blue-600 hover:underline"
          >
            <ExternalLink size={14} className="mt-[1px] text-blue-600" />
          </a>
        </div>
        <p className="text-l text-gray-700">
          Owned By: <span className="font-semibold">{owner}</span>
        </p>
      </div>

      <div className="mt-1 flex items-center gap-4">
        {languages?.map((language, index) => {
          const color = LangColor[language] || "bg-gray-400";
          return (
            <div key={index} className="text-md flex items-center gap-1">
              <div className={`h-2.5 w-2.5 rounded-full ${color}`} />
              <span className="text-l text-gray-700">{language}</span>
            </div>
          );
        })}
      </div>

      <p className="text-l mt-2 leading-snug text-gray-700">
        {description ||
          "No description for the given repository. Lorem ipsum dolor sit amet."}
      </p>
    </div>
  );
};

export default RepositoryCard;
