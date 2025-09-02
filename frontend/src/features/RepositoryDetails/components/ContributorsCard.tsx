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
    <div className="*:data-[slot=avatar]:ring-background group relative: flex -space-x-4 *:data-[slot=avatar]:ring-2 *:data-[slot=avatar]:grayscale">
      <Avatar>
        <Link to={githubUrl}>
          <AvatarImage
            src={avatarUrl}
            alt="contributors img"
            className="h-15 w-15 rounded-full"
          />
          <AvatarFallback>Contributors-Image</AvatarFallback>
          <div className="left absolute bottom-[-5.5rem] z-1 hidden w-max -translate-y-140 rounded bg-yellow-100 px-2 py-1 text-xs text-gray-600 shadow group-hover:block">
            {name} <br />
            {contributions} Contributions
            <br />
          </div>
        </Link>
      </Avatar>
    </div>
  );
};

export default ContributorsCard;
