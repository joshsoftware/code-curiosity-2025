import Coin from "@/shared/components/common/Coin";
import { Card, CardContent } from "@/shared/components/ui/card";
import { LangColor } from "@/shared/constants/constants";
import { format } from "date-fns";
import type { FC } from "react";
import { useNavigate } from "react-router-dom";

interface RepositoriesCardProps {
  id: number;
  name: string;
  languages: string[];
  description: string;
  updatedOn: string;
  coins: number;
}

const RepositoriesCard: FC<RepositoriesCardProps> = ({
  id,
  name,
  languages,
  description,
  updatedOn,
  coins
}) => {
  const navigate = useNavigate();

  return (
    <Card className="border-none p-4 shadow-none">
      <CardContent className="space-y-1">
        <div
          className="flex items-start justify-between"
          onClick={() => navigate(`/repositories/${id}`)}
        >
          <p className="cursor-pointer text-lg font-semibold">{name}</p>
          <div className="flex items-center justify-between gap-2">
            <Coin />
            <span className="text-cc-app-blue text-lg font-bold">{coins}</span>
          </div>
        </div>

        <div className="flex items-center gap-4">
          {languages?.map((language, index) => {
            const color = LangColor[language] || "bg-gray-400";
            return (
              <div key={index} className="flex items-center gap-1">
                <div className={`h-3 w-3 rounded-full ${color}`}></div>
                <span className="text-sm text-gray-700">{language}</span>
              </div>
            );
          })}
        </div>

        <p className="text-sm leading-relaxed">
          {description || "No description for the given repository"}
        </p>

        <p className="text-xs text-gray-500">
          Updated on {format(new Date(updatedOn), "MMM d, yyyy")}
        </p>
      </CardContent>
    </Card>
  );
};

export default RepositoriesCard;
