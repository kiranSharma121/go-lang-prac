import { useState } from "react";
import "./App.css";
import RegisterForm from "./RegisterForm";
import LoginForm from "./login";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <ToastContainer />

      <LoginForm />
    </>
  );
}

export default App;
