import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogDescription
} from "@/shared/components/ui/dialog";
import { Button } from "@/shared/components/ui/button";
import {
  useLoggedInUser,
  useSoftDeleteUser
} from "@/api/queries/UserProfileDetails";

interface Props {
  open: boolean;
  onClose: () => void;
}

const DeleteAccount = ({ open, onClose }: Props) => {
  const { data } = useLoggedInUser();
  const user = data?.data;
  const { mutate: softDeleteUser, isPending } = useSoftDeleteUser();

  const handleDelete = () => {
    if (!user?.userId) return;
    softDeleteUser(user.userId, {
      onSuccess: () => {
        onClose();
      }
    });
  };

  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Delete Account</DialogTitle>
          <DialogDescription>
            Deleting your account will not immediately erase your data. We will
            retain your information for 3 months in case you choose to return.
            After that period, your data will be permanently removed from our
            platform.
          </DialogDescription>
        </DialogHeader>

        <DialogFooter>
          <Button variant="outline" onClick={onClose}>
            Cancel
          </Button>
          <Button
            variant="destructive"
            onClick={handleDelete}
            disabled={isPending}
          >
            {isPending ? "Deleting..." : "Yes, Delete My Account"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default DeleteAccount;
