import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from "@/shared/components/ui/dialog";
import { Button } from "@/shared/components/ui/button";
import { useState } from "react";
import DeleteAccount from "./DeleteAccount";

interface Props {
  open: boolean;
  onClose: () => void;
  onUpdateEmail: () => void;
  
}

const SettingsDialog = ({
  open,
  onClose,
  onUpdateEmail,
}: Props) => {
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  return (
    <Dialog open={open} onOpenChange={onClose} >
      <DialogContent className="bg-white">
        <DialogHeader>
          <DialogTitle className="text-black">Account Settings</DialogTitle>
        </DialogHeader>

        <div className="space-y-4 py-4">
          <Button
            variant="ccAppOutlineMidBlue"
            className="hover: w-full cursor-pointer"
            onClick={onUpdateEmail}
          >
            Update Email
          </Button>
          <Button
            variant="ccAppOutlineRed"
            className="w-full cursor-pointer"
            onClick={() => setShowDeleteConfirm(true)}
          >
            Delete Account
          </Button>

          <DeleteAccount
            open={showDeleteConfirm}
            onClose={() => setShowDeleteConfirm(false)}
          />
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
            Cancel
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default SettingsDialog;
