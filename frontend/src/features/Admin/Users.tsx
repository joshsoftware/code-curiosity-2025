import { type FC, useEffect, useState } from "react";
import { Card } from "@/shared/components/ui/card";
import defaultAvatar from "@/assets/default-profile-pic.svg";
import {
  Loader,
  User,
  Shield,
  ShieldOff,
  ChevronLeft,
  ChevronRight
} from "lucide-react";
import { useGetAllUsers, useUpdateUserBlockStatus } from "@/api/queries/Admin";
import Coin from "@/shared/components/common/Coin";
import { toast } from "sonner";

const ITEMS_PER_PAGE = 10;

export const AllUsersList: FC = () => {
  const { data, isLoading } = useGetAllUsers();
  const { mutate: updateBlockStatus } = useUpdateUserBlockStatus();

  const [currentPage, setCurrentPage] = useState(1);
  const [users, setUsers] = useState(data?.data || []);

  useEffect(() => {
    if (data?.data) {
      setUsers(data.data);
    }
  }, [data]);

  const totalPages = Math.ceil(users.length / ITEMS_PER_PAGE);
  const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
  const endIndex = startIndex + ITEMS_PER_PAGE;
  const currentUsers = users.slice(startIndex, endIndex);

  const handleBlockToggle = (userId: number, block: boolean) => {
    updateBlockStatus(
      { userId, block },
      {
        onSuccess: () => {
          setUsers(prev =>
            prev.map(user =>
              user.userId === userId ? { ...user, isBlocked: block } : user
            )
          );
          toast.success(
            `User has been ${block ? "blocked" : "unblocked"} successfully.`
          );
        },
        onError: () => {
          toast.error("Something went wrong while updating block status.");
        }
      }
    );
  };

  const goToPage = (page: number) => {
    setCurrentPage(Math.max(1, Math.min(page, totalPages)));
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center py-20">
        <Loader className="h-8 w-8 animate-spin text-blue-500" />
        <span className="ml-2 text-gray-600">Loading users...</span>
      </div>
    );
  }

  if (!users.length) {
    return (
      <div className="flex items-center justify-center py-20">
        <div className="text-center">
          <User className="mx-auto mb-4 h-12 w-12 text-gray-400" />
          <p className="text-gray-600">No users found</p>
        </div>
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-4xl p-6">
      <div className="mb-6">
        <h1 className="mb-2 text-2xl font-bold text-gray-900">
          User Management
        </h1>
        <p className="text-gray-600">
          Showing {startIndex + 1}-{Math.min(endIndex, users.length)} of{" "}
          {users.length} users
        </p>
      </div>

      <div className="mb-6 space-y-3">
        {currentUsers.map(user => (
          <Card
            key={user.userId}
            className="p-4 transition-shadow hover:shadow-md"
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="flex-shrink-0">
                  <img
                    src={user.avatarUrl || defaultAvatar}
                    alt="User Avatar"
                    className="h-10 w-10 rounded-full object-cover"
                  />
                </div>

                <div className="min-w-0 flex-1">
                  <div className="flex items-center space-x-3">
                    <h3 className="truncate text-sm font-medium text-gray-900">
                      {user.githubUsername}
                    </h3>
                    <div className="flex items-center space-x-1 text-sm text-gray-600">
                      <Coin />
                      <span>{user.currentBalance}</span>
                    </div>
                  </div>
                  <div className="mt-1 flex items-center">
                    <span
                      className={`inline-flex items-center rounded-full px-2 py-1 text-xs font-medium ${
                        user.isBlocked
                          ? "bg-red-100 text-red-800"
                          : "bg-green-100 text-green-800"
                      }`}
                    >
                      {user.isBlocked ? (
                        <>
                          <ShieldOff className="mr-1 h-3 w-3" />
                          Blocked
                        </>
                      ) : (
                        <>
                          <Shield className="mr-1 h-3 w-3" />
                          Active
                        </>
                      )}
                    </span>
                  </div>
                </div>
              </div>

              <div className="flex-shrink-0">
                <button
                  onClick={() =>
                    handleBlockToggle(user.userId, !user.isBlocked)
                  }
                  className={`rounded-md px-4 py-2 text-sm font-medium transition-colors ${
                    user.isBlocked
                      ? "bg-green-600 text-white hover:bg-green-700"
                      : "bg-red-600 text-white hover:bg-red-700"
                  }`}
                >
                  {user.isBlocked ? "Unblock" : "Block"}
                </button>
              </div>
            </div>
          </Card>
        ))}
      </div>

      {totalPages > 1 && (
        <div className="flex items-center justify-between border-t pt-6">
          <div className="flex items-center space-x-2">
            <button
              onClick={() => goToPage(currentPage - 1)}
              disabled={currentPage === 1}
              className="flex items-center rounded-md border border-gray-300 bg-white px-3 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
            >
              <ChevronLeft className="mr-1 h-4 w-4" />
              Previous
            </button>

            <div className="flex space-x-1">
              {Array.from({ length: Math.min(5, totalPages) }, (_, i) => {
                let pageNum;
                if (totalPages <= 5) {
                  pageNum = i + 1;
                } else if (currentPage <= 3) {
                  pageNum = i + 1;
                } else if (currentPage >= totalPages - 2) {
                  pageNum = totalPages - 4 + i;
                } else {
                  pageNum = currentPage - 2 + i;
                }

                return (
                  <button
                    key={pageNum}
                    onClick={() => goToPage(pageNum)}
                    className={`rounded-md px-3 py-2 text-sm font-medium ${
                      currentPage === pageNum
                        ? "bg-blue-600 text-white"
                        : "border border-gray-300 bg-white text-gray-700 hover:bg-gray-50"
                    }`}
                  >
                    {pageNum}
                  </button>
                );
              })}
            </div>

            <button
              onClick={() => goToPage(currentPage + 1)}
              disabled={currentPage === totalPages}
              className="flex items-center rounded-md border border-gray-300 bg-white px-3 py-2 text-sm font-medium text-gray-500 hover:bg-gray-50 disabled:cursor-not-allowed disabled:opacity-50"
            >
              Next
              <ChevronRight className="ml-1 h-4 w-4" />
            </button>
          </div>

          <div className="text-sm text-gray-700">
            Page {currentPage} of {totalPages}
          </div>
        </div>
      )}
    </div>
  );
};
