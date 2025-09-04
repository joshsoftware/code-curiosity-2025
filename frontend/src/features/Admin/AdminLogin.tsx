import { type FC, useState } from "react";
import type { AdminCredentials } from "@/shared/types/types";
import { useLogInAdmin } from "@/api/queries/Admin";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import {
  ACCESS_TOKEN_KEY,
  USER_DATA_KEY
} from "@/shared/constants/local-storage";
import { Button } from "@/shared/components/ui/button";
import { toast } from "sonner";
import { Eye, EyeOff } from "lucide-react";

const AdminLogin: FC = () => {
  const { register, handleSubmit } = useForm<AdminCredentials>();
  const navigate = useNavigate();
  const { mutate, isPending } = useLogInAdmin();
  const [showPassword, setShowPassword] = useState(false);

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
          localStorage.setItem(USER_DATA_KEY, JSON.stringify(res.data));
          navigate("/admin/users");
        },
        onError: err => {
          toast.error("Invalid credentials");
          console.error("Admin login error", err);
        }
      });
    } catch (err) {
      console.error("Unexpected submit error", err);
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
          Email <span className="text-red-500">*</span>
        </label>
        <input
          id="email"
          type="email"
          placeholder="example@gmail.com"
          className="w-full rounded-md border border-gray-300 bg-white p-2 text-sm shadow-sm focus:border-blue-500 focus:bg-white focus:ring-0 focus:outline-none"
          {...register("email", { required: "Email is required" })}
        />
      </div>

      <div className="space-y-2">
        <label
          htmlFor="password"
          className="block text-sm font-medium text-gray-700"
        >
          Password <span className="text-red-500">*</span>
        </label>
        <div className="relative">
          <input
            id="password"
            type={showPassword ? "text" : "password"}
            placeholder="password"
            className="w-full rounded-md border border-gray-300 bg-white p-2 pr-10 text-sm shadow-sm focus:border-blue-500 focus:bg-white focus:ring-0 focus:outline-none"
            {...register("password", { required: "Password is required" })}
          />
          <button
            type="button"
            onClick={() => setShowPassword(prev => !prev)}
            className="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-500 hover:text-gray-700 focus:outline-none"
          >
            {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
          </button>
        </div>
      </div>

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
