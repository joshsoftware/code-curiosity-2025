// import { type FC } from "react";
// import type { AdminCredentials } from "@/shared/types/types";
// import { useLogInAdmin } from "@/api/queries/Admin";

// const AdminLogin: FC = () => {
//   const {
//     register,
//     handleSubmit,
//     formState: { errors },
//   } = useForm<AdminCredentials>();

//   const { mutate, isPending, isError, error } = useLogInAdmin();

//   const onSubmit = (data: AdminCredentials) => {
//     mutate(data);
//   };

//   return (
//     <form
//       onSubmit={handleSubmit(onSubmit)}
//       className="w-full max-w-md space-y-6"
//     >
//       <div className="space-y-2">
//         <label htmlFor="email" className="block text-sm font-medium text-gray-700">
//           Email
//         </label>
//         <input
//           id="email"
//           type="email"
//           className="w-full rounded-md border border-gray-300 p-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
//           {...register("email", { required: "Email is required" })}
//         />
//         {errors.email && (
//           <p className="text-sm text-red-600">{errors.email.message}</p>
//         )}
//       </div>

//       <div className="space-y-2">
//         <label htmlFor="password" className="block text-sm font-medium text-gray-700">
//           Password
//         </label>
//         <input
//           id="password"
//           type="password"
//           className="w-full rounded-md border border-gray-300 p-2 text-sm shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
//           {...register("password", { required: "Password is required" })}
//         />
//         {errors.password && (
//           <p className="text-sm text-red-600">{errors.password.message}</p>
//         )}
//       </div>

//       {isError && (
//         <p className="text-sm text-red-600">
//           {(error as Error)?.message || "Failed to log in"}
//         </p>
//       )}

//       <button
//         type="submit"
//         disabled={isPending}
//         className="w-full rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
//       >
//         {isPending ? "Logging in..." : "Log in as Admin"}
//       </button>
//     </form>
//   );
// };

// export default AdminLogin;
