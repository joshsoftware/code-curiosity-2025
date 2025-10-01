export const Layout = {
  AuthLayout: "AuthLayout",
  DashboardLayout: "DashboardLayout",
  AdminLayout: "AdminLayout",
  None: "None"
} as const;

export type LayoutType = (typeof Layout)[keyof typeof Layout];
