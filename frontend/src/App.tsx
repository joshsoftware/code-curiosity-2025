import "./App.css";
import { UserProvider } from "./context/AuthProvider";
import Router from "./root/Router";

function App() {
  return (
    <UserProvider>
      <Router />
    </UserProvider>
  );
}

export default App;
