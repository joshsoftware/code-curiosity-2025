import { type FC } from "react";
import type { AdminCredentials } from "@/shared/types/types";
import { useLogInAdmin } from "@/api/queries/Admin";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { ACCESS_TOKEN_KEY } from "@/shared/constants/local-storage";
import { Button } from "@/shared/components/ui/button";

const AdminLogin: FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<AdminCredentials>();
  const navigate = useNavigate();
  const { mutate, isPending, isError, error } = useLogInAdmin();

  const onSubmit = async (
    data: AdminCredentials,
    event?: React.BaseSyntheticEvent
  ) => {
    try {
      event?.preventDefault();
      console.log(" Submitting admin login", data);
      mutate(data, {
        onSuccess: res => {
          console.log(" Admin login success", res);
          localStorage.setItem(ACCESS_TOKEN_KEY, res.data.jwtToken);
          navigate("/admin/users");
        },
        onError: err => {
          console.error(" Admin login error", err);
        }
      });
    } catch (err) {
      console.error(" Unexpected submit error", err);
    }
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="w-full max-w-md space-y-6"
    >
      <div className="space-y-2">
        <label
          htmlFor="email"
          className="block text-sm font-medium text-gray-700"
        >
          Email
        </label>
        <input
          id="email"
          type="email"
          className="w-full rounded-md border border-gray-300 bg-white p-2 text-sm shadow-sm focus:border-blue-500 focus:bg-white focus:ring-0 focus:outline-none"
          {...register("email", { required: "Email is required" })}
        />
        {errors.email && (
          <p className="text-sm text-red-600">{errors.email.message}</p>
        )}
      </div>

      <div className="space-y-2">
        <label
          htmlFor="password"
          className="block text-sm font-medium text-gray-700"
        >
          Password
        </label>
        <input
          id="password"
          type="password"
          className="w-full rounded-md border border-gray-300 bg-white p-2 text-sm shadow-sm focus:border-blue-500 focus:bg-white focus:ring-0 focus:outline-none"
          {...register("password", { required: "Password is required" })}
        />
        {errors.password && (
          <p className="text-sm text-red-600">{errors.password.message}</p>
        )}
      </div>

      {isError && (
        <p className="text-sm text-red-600">
          {error?.message || "Failed to log in"}
        </p>
      )}

      <Button
        type="submit"
        disabled={isPending}
        className="bg-cc-app-orange w-full rounded-md px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:outline-none"
      >
        {isPending ? "Logging in..." : "Log in as Admin"}
      </Button>
    </form>
  );
};

export default AdminLogin;