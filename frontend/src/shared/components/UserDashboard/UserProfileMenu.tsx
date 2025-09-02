import { MoreVertical } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator
} from "@/shared/components/ui/dropdown-menu";
import { Button } from "@/shared/components/ui/button";

interface Props {
  onSettingsClick: () => void;
  onLogoutClick: () => void;
}

const UserProfileMenu = ({ onSettingsClick, onLogoutClick }: Props) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ccAppOutlineMidBlue"
          size="sm"
          className="cursor-pointer"
        >
          <MoreVertical className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-40">
        <DropdownMenuItem
          onClick={onSettingsClick}
          className="hover:cursor-pointer"
        >
          Settings
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem
          onClick={onLogoutClick}
          className="hover:cursor-pointer"
        >
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UserProfileMenu;
