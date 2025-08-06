import { useState } from "react";
import { ExternalLink, MoreVertical } from "lucide-react";

import Coin from "@/shared/components/common/Coin";
import { Button } from "@/shared/components/ui/button";
import DefaultProfilePic from "@/assets/default-profile-pic.svg";
import { Separator } from "@/shared/components/ui/separator";
import { useLoggedInUser } from "@/api/queries/UserProfileDetails";
import { Link, useNavigate } from "react-router-dom";

import UserProfileMenu from "./UserProfileMenu";
import UserEmail from "./UserEmail";
import SettingsDialog from "./SettingsDialog";
import { clearAccessToken } from "@/shared/utils/local-storage";

const UserProfileDetails = () => {
  const navigate = useNavigate();
  const { data } = useLoggedInUser();
  const user = data?.data;

  const [showSettingsDialog, setShowSettingsDialog] = useState(false);
  const [showEmailDialog, setShowEmailDialog] = useState(false);

  const handleSettingsClick = () => setShowSettingsDialog(true);
  const handleLogoutClick = () => {
    clearAccessToken();
    navigate("/login");
  };

  const handleUpdateEmail = () => {
    setShowSettingsDialog(false);
    setShowEmailDialog(true);
  };

  return (
    <div>
      <div className="space-y-4">
        <div className="relative flex items-start justify-center">
          <div className="border-cc-app-sky-blue mx-auto h-48 w-48 rounded-full border-8">
            <div className="h-44 w-44 overflow-hidden rounded-full border-4 border-gray-300 bg-white">
              <img
                src={user?.avatarUrl || DefaultProfilePic}
                alt="Profile"
                className="h-full w-full object-cover"
              />
            </div>
          </div>

          <div className="absolute top-0 right-0">
            <UserProfileMenu
              onSettingsClick={handleSettingsClick}
              onLogoutClick={handleLogoutClick}
            />
          </div>
        </div>

        <div className="mb-6 flex items-center justify-center gap-2">
          <h2 className="text-3xl font-medium text-white">
            {user?.githubUsername || "username"}
          </h2>
          <Link to={user ? `https://github.com/${user?.githubUsername}` : ""}>
            <ExternalLink className="h-5 w-5 cursor-pointer text-white" />
          </Link>
        </div>
      </div>

      <Separator className="my-6 bg-gradient-to-r from-transparent via-white to-transparent opacity-70" />

      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2">
          <Coin />
          <span className="text-cc-app-yellow text-2xl font-bold">
            {user?.currentBalance || "0"}
          </span>
        </div>
        <div className="flex gap-2">
          <Button variant="ccAppOutlineMidBlue" size="sm">
            Redeem
          </Button>
          <Button variant="ccAppOutlineMidBlue" size="sm">
            <MoreVertical className="h-4 w-4" />
          </Button>
        </div>
      </div>

      <SettingsDialog
        open={showSettingsDialog}
        onClose={() => setShowSettingsDialog(false)}
        onUpdateEmail={handleUpdateEmail}
      />
      {showEmailDialog && (
        <UserEmail
          defaultEmail={user?.email || ""}
          onClose={() => setShowEmailDialog(false)}
        />
      )}
    </div>
  );
};

export default UserProfileDetails;
