import { useState } from "react";

function RegisterForm() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const userData = { username, email, password };

    try {
      const response = await fetch("http://localhost:8080/insertone", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(userData),
      });

      if (response.ok) {
        alert("User registered successfully");
        window.location.reload(); // Refresh the page
      } else {
        setErrorMessage("Failed to register user");
      }
    } catch (error) {
      console.error("Error:", error);
      setErrorMessage("Error occurred while registering user");
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="bg-white p-8 rounded-lg shadow-xl w-full max-w-md">
        <h2 className="text-3xl font-bold mb-6 text-center text-gray-700">
          Register
        </h2>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Username"
            required
            className="block w-full p-3 mb-4 border border-gray-300 rounded-lg focus:border-blue-500 focus:outline-none shadow-sm transition duration-200"
          />
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            required
            className="block w-full p-3 mb-4 border border-gray-300 rounded-lg focus:border-blue-500 focus:outline-none shadow-sm transition duration-200"
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            required
            className="block w-full p-3 mb-4 border border-gray-300 rounded-lg focus:border-blue-500 focus:outline-none shadow-sm transition duration-200"
          />
          {errorMessage && (
            <div className="text-red-500 mb-4 font-semibold">
              {errorMessage}
            </div>
          )}
          <button
            type="submit"
            className="w-full bg-gradient-to-r from-blue-500 to-purple-500 text-white p-3 rounded-lg hover:from-blue-600 hover:to-purple-600 transition duration-300"
          >
            Register
          </button>
        </form>
      </div>
    </div>
  );
}

export default RegisterForm;
