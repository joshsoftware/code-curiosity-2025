import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter
} from "@/shared/components/ui/dialog";
import { Input } from "@/shared/components/ui/input";
import { Button } from "@/shared/components/ui/button";
import { useState } from "react";
import { useUpdateUserEmail } from "@/api/queries/UserProfileDetails";
import { queryClient } from "@/api/react-query";
import { LOGGED_IN_USER_QUERY_KEY } from "@/shared/constants/query-keys";
import { toast } from "sonner";

interface Props {
  defaultEmail: string;
  onClose: () => void;
}

const UserEmail = ({ defaultEmail, onClose }: Props) => {
  const [email, setEmail] = useState(defaultEmail);
  const { mutate: updateEmail, isPending } = useUpdateUserEmail();

  const isValidEmail = (email: string) =>
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);

  const handleUpdate = () => {
    if (!isValidEmail(email)) return;
    updateEmail(email, {
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: [LOGGED_IN_USER_QUERY_KEY]
        });
        toast.success("email updated successfully");
        onClose();
      }
    });
  };

  return (
    <Dialog open onOpenChange={onClose}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Update Email</DialogTitle>
        </DialogHeader>

        We need your email, to send you notifications

        <Input
          type="email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          placeholder="Enter new email"
          className="selection:bg-cc-app-blue"
        />

        <DialogFooter>
          <Button
            variant="ccAppOutlineMidBlue"
            onClick={handleUpdate}
            disabled={isPending || !email || !isValidEmail(email)}
          >
            {isPending ? "Updating..." : "Update"}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default UserEmail;
