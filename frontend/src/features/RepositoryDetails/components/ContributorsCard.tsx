import { Avatar, AvatarFallback, AvatarImage } from "@radix-ui/react-avatar";
import type { FC } from "react";
import { Link } from "react-router-dom";

interface ContributorsCardProps {
  name: string;
  avatarUrl: string;
  contributions: number;
  githubUrl: string;
}

const ContributorsCard: FC<ContributorsCardProps> = ({
  name,
  avatarUrl,
  contributions,
  githubUrl
}) => {
  return (
    <div className="group *:data-[slot=avatar]:ring-background relative flex -space-x-4 *:data-[slot=avatar]:ring-2 *:data-[slot=avatar]:grayscale">
      <Avatar>
        <Link to={githubUrl} className="relative inline-block">
          <AvatarImage
            src={avatarUrl}
            alt="contributors img"
            className="h-15 w-15 rounded-full"
          />
          <AvatarFallback>Contributors-Image</AvatarFallback>

          {/* Tooltip */}
          <div className="absolute top-full left-1/2 mt-2 hidden -translate-x-1/2 rounded bg-yellow-100 px-3 py-1.5 text-xs whitespace-nowrap text-gray-600 shadow-md group-hover:block">
            <p className="font-medium">{name}</p>
            <p>{contributions} Contributions</p>
          </div>
        </Link>
      </Avatar>
    </div>
  );
};

export default ContributorsCard;
