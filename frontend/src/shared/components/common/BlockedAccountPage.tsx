import {
  Card,
  CardContent,
  CardHeader,
  CardFooter
} from "@/shared/components/ui/card";
import { Button } from "@/shared/components/ui/button";
import { clearAccessToken } from "@/shared/utils/local-storage";
import { useNavigate } from "react-router-dom";
import { AlertTriangle } from "lucide-react";

const BlockedAccountPage = () => {
  const navigate = useNavigate();
  clearAccessToken();

  const handleBackToLogin = () => {
    clearAccessToken();
    navigate("/login");
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-100 px-4">
      <Card className="w-full max-w-md bg-white text-center shadow-md">
        <CardHeader className="flex flex-col items-center space-y-4 pt-8">
          <AlertTriangle className="h-12 w-12 text-red-500" />
          <h2 className="text-xl font-semibold text-gray-800">
            Account Blocked
          </h2>
        </CardHeader>

        <CardContent>
          <p className="text-sm text-gray-600">
            Your account has been blocked due to policy violations or unusual
            activity. If you believe this is a mistake, please contact our
            support team.
          </p>
        </CardContent>

        <CardFooter className="flex justify-center pb-8">
          <Button
            onClick={handleBackToLogin}
            className="bg-cc-app-orange text-white"
          >
            Back to Login
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
};

export default BlockedAccountPage;
