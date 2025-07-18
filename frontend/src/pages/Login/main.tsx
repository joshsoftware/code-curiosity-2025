import { useContext } from "react";
import { UserContext } from "../../context/AuthProvider";

function Login() {
  const { user } = useContext(UserContext);

  return <div>login page username : {user.githubUsername}</div>;
}
export default Login;
