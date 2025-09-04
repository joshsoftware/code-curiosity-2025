import { type ReactNode } from "react";

interface AppLayoutProps {
  children: ReactNode;
}

const AppLayout = ({ children }: AppLayoutProps) => {
  return (
    <div className="w-screen h-screen flex items-center justify-center bg-black">
      <div className="h-full w-full max-w-[1900px] max-h-[1000px] bg-white shadow-lg ">
        {children}
      </div>
    </div>
  );
};

export default AppLayout;
