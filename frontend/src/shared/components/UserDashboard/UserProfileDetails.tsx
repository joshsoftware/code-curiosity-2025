import { ExternalLink, MoreHorizontal } from "lucide-react";

import Coin from "@/shared/components/common/Coin";
import { Button } from "@/shared/components/ui/button";
import DefaultProfilePic from "@/assets/default-profile-pic.svg"
import { Separator } from "@/shared/components/ui/separator";
import { useLoggedInUser } from "@/api/queries/UserProfileDetails";
import { Link } from "react-router-dom";

const UserProfileDetails = () => {
  const { data } = useLoggedInUser();
  const user = data?.data

  return (
    <div>
      <div className="space-y-4">
        <div className="border-cc-app-sky-blue mx-auto h-48 w-48 rounded-full border-8">
          <div className="h-44 w-44 overflow-hidden rounded-full border-4 border-gray-300 bg-white">
            <img
              src={user?.avatarUrl || DefaultProfilePic}
              alt="Profile"
              className="h-full w-full object-cover"
            />
          </div>
        </div>

        <div className="mb-6 flex items-center justify-center gap-2">
          <h2 className="text-3xl font-medium text-white">{user?.githubUsername || "username"}</h2>

          <Link to={user ? `https://github.com/${user?.githubUsername}` : ''}>
            <ExternalLink className="h-5 w-5 cursor-pointer text-white" />
          </Link>
        </div>
      </div>
      <Separator className="my-6 bg-gradient-to-r from-transparent via-white to-transparent opacity-70" />
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Coin />
          <span className="text-cc-app-orange text-2xl font-bold">{user?.currentBalance || "0"}</span>
        </div>
        <div className="flex gap-2">
          <Button
            variant="ccAppOutlineMidBlue"
            size="sm"
            className="cursor-pointer"
          >
            Redeem
          </Button>
          <Button
            variant="ccAppOutlineMidBlue"
            size="sm"
            className="cursor-pointer"
          >
            <MoreHorizontal className="h-4 w-4" />
          </Button>
        </div>
      </div>
    </div>
  );
};

export default UserProfileDetails;
