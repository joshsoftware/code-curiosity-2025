import { type FC, useEffect, useState } from "react";
import {
  type ColumnDef,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  useReactTable
} from "@tanstack/react-table";
import defaultAvatar from "@/assets/default-profile-pic.svg";
import {
  Loader,
  User,
  Shield,
  ShieldOff,
  ChevronLeft,
  ChevronRight,
  ChevronsLeft,
  ChevronsRight,
  Search,
  X
} from "lucide-react";
import { useGetAllUsers, useUpdateUserBlockStatus } from "@/api/queries/Admin";
import Coin from "@/shared/components/common/Coin";
import { toast } from "sonner";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow
} from "@/shared/components/ui/table";
import { Button } from "@/shared/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from "@/shared/components/ui/select";
import { Input } from "@/shared/components/ui/input";

interface User {
  userId: number;
  githubUsername: string;
  avatarUrl?: string;
  currentBalance: number;
  isBlocked: boolean;
}

export const AllUsersList: FC = () => {
  const { data, isLoading } = useGetAllUsers();
  const { mutate: updateBlockStatus } = useUpdateUserBlockStatus();
  const [users, setUsers] = useState<User[]>(data?.data || []);

  useEffect(() => {
    if (data?.data) {
      setUsers(data.data);
    }
  }, [data]);

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

  const columns: ColumnDef<User>[] = [
    {
      accessorKey: "githubUsername",
      header: "User",
      cell: ({ row }) => {
        const user = row.original;
        return (
          <div className="flex items-center space-x-3">
            <img
              src={user.avatarUrl || defaultAvatar}
              alt="User Avatar"
              className="h-8 w-8 rounded-full object-cover"
            />
            <span className="font-medium text-gray-900">
              {user.githubUsername}
            </span>
          </div>
        );
      }
    },
    {
      accessorKey: "currentBalance",
      header: "Balance",
      cell: ({ row }) => {
        return (
          <div className="flex items-center space-x-1 text-gray-700">
            <Coin />
            <span>{row.getValue("currentBalance")}</span>
          </div>
        );
      }
    },
    {
      accessorKey: "isBlocked",
      header: "Status",
      cell: ({ row }) => {
        const isBlocked = row.getValue("isBlocked") as boolean;
        return (
          <span
            className={`inline-flex w-[80px] max-w-[80px] items-center rounded-full px-2 py-1 text-xs font-medium ${
              isBlocked
                ? "bg-red-100 text-red-800"
                : "bg-green-100 text-green-800"
            }`}
          >
            {isBlocked ? (
              <>
                <ShieldOff className="mr-1 h-3 w-3 max-w-3" />
                Blocked
              </>
            ) : (
              <>
                <Shield className="mr-1 h-3 w-3 max-w-3" />
                Active
              </>
            )}
          </span>
        );
      }
    },
    {
      id: "actions",
      header: () => <div className="text-right">Actions</div>,
      cell: ({ row }) => {
        const user = row.original;
        return (
          <div className="text-right">
            <Button
              size="sm"
              variant={user.isBlocked ? "ccAppOutline" : "success"}
              onClick={() => handleBlockToggle(user.userId, !user.isBlocked)}
              className="w-[80px] max-w-[80px]"
            >
              {user.isBlocked ? "Unblock" : "Block"}
            </Button>
          </div>
        );
      }
    }
  ];

  const table = useReactTable({
    data: users,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    initialState: {
      pagination: {
        pageSize: 10
      }
    }
  });

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
    <div className="mx-auto max-w-5xl p-6">
      <div className="mb-6">
        <h1 className="mb-2 text-2xl font-bold text-gray-900">
          User Management
        </h1>
        <div className="flex items-center justify-between">
          <p className="text-gray-600">
            Showing{" "}
            {table.getState().pagination.pageIndex *
              table.getState().pagination.pageSize +
              1}
            -
            {Math.min(
              (table.getState().pagination.pageIndex + 1) *
                table.getState().pagination.pageSize,
              table.getFilteredRowModel().rows.length
            )}{" "}
            of {table.getFilteredRowModel().rows.length} users
            {table.getState().globalFilter &&
              ` (filtered from ${users.length} total)`}
          </p>
        </div>
      </div>

      <div className="mb-4 flex items-center gap-4">
        <div className="relative max-w-sm flex-1">
          <Search className="text-muted-foreground absolute top-2.5 left-2 h-4 w-4" />
          <Input
            placeholder="Search users..."
            value={table.getState().globalFilter ?? ""}
            onChange={e => table.setGlobalFilter(e.target.value)}
            className="pl-8"
          />
          {table.getState().globalFilter && (
            <Button
              variant="ghost"
              size="sm"
              className="absolute top-1 right-1 h-6 w-6 p-0"
              onClick={() => table.setGlobalFilter("")}
            >
              <X className="h-3 w-3" />
            </Button>
          )}
        </div>

        <Select
          value={
            table.getColumn("isBlocked")?.getFilterValue() === undefined
              ? "all"
              : table.getColumn("isBlocked")?.getFilterValue() === true
                ? "blocked"
                : "active"
          }
          onValueChange={value => {
            if (value === "all") {
              table.getColumn("isBlocked")?.setFilterValue(undefined);
            } else if (value === "blocked") {
              table.getColumn("isBlocked")?.setFilterValue(true);
            } else {
              table.getColumn("isBlocked")?.setFilterValue(false);
            }
          }}
        >
          <SelectTrigger className="w-[150px]">
            <SelectValue placeholder="Filter status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All Users</SelectItem>
            <SelectItem value="active">Active Only</SelectItem>
            <SelectItem value="blocked">Blocked Only</SelectItem>
          </SelectContent>
        </Select>
      </div>

      <div className="rounded-md border">
        <Table>
          <TableHeader>
            {table.getHeaderGroups().map(headerGroup => (
              <TableRow key={headerGroup.id}>
                {headerGroup.headers.map(header => {
                  return (
                    <TableHead key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                            header.column.columnDef.header,
                            header.getContext()
                          )}
                    </TableHead>
                  );
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map(row => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map(cell => (
                    <TableCell key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>

      {/* Pagination */}
      <div className="flex items-center justify-between space-x-2 py-4">
        <div className="flex items-center space-x-2">
          <p className="text-sm font-medium">Rows per page</p>
          <Select
            value={`${table.getState().pagination.pageSize}`}
            onValueChange={value => {
              table.setPageSize(Number(value));
            }}
          >
            <SelectTrigger className="h-8 w-[70px]">
              <SelectValue placeholder={table.getState().pagination.pageSize} />
            </SelectTrigger>
            <SelectContent side="top">
              {[10, 20, 30, 40, 50].map(pageSize => (
                <SelectItem key={pageSize} value={`${pageSize}`}>
                  {pageSize}
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
        </div>

        <div className="flex items-center space-x-6 lg:space-x-8">
          <div className="flex w-[100px] items-center justify-center text-sm font-medium">
            Page {table.getState().pagination.pageIndex + 1} of{" "}
            {table.getPageCount()}
          </div>
          <div className="flex items-center space-x-2">
            <Button
              variant="outline"
              className="hidden h-8 w-8 p-0 lg:flex"
              onClick={() => table.setPageIndex(0)}
              disabled={!table.getCanPreviousPage()}
            >
              <span className="sr-only">Go to first page</span>
              <ChevronsLeft className="h-4 w-4" />
            </Button>
            <Button
              variant="outline"
              className="h-8 w-8 p-0"
              onClick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            >
              <span className="sr-only">Go to previous page</span>
              <ChevronLeft className="h-4 w-4" />
            </Button>
            <Button
              variant="outline"
              className="h-8 w-8 p-0"
              onClick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            >
              <span className="sr-only">Go to next page</span>
              <ChevronRight className="h-4 w-4" />
            </Button>
            <Button
              variant="outline"
              className="hidden h-8 w-8 p-0 lg:flex"
              onClick={() => table.setPageIndex(table.getPageCount() - 1)}
              disabled={!table.getCanNextPage()}
            >
              <span className="sr-only">Go to last page</span>
              <ChevronsRight className="h-4 w-4" />
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
};
