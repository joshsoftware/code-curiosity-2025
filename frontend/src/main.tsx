import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { QueryClientProvider } from "@tanstack/react-query";
import { Toaster } from "sonner";

import Router from "@/root/Router";
import { queryClient } from "@/api/react-query.ts";
import "./index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <Router />
      <Toaster richColors theme="light" />
    </QueryClientProvider>
  </StrictMode>
);
