import { useState } from "react";
import "./Login.css";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMsg, setErrorMsg] = useState("");

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg("");

    try {
      const res = await fetch("http://localhost:3000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (!res.ok) {
        setErrorMsg(data.error || "Login fallido");
        return;
      }

      localStorage.setItem("token", data.token);
      alert("Login exitoso");
      window.location.href = "/dashboard";
    } catch (err) {
      setErrorMsg("Error al conectar con el servidor");
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleLogin}>
        <h2>Iniciar Sesión</h2>
        {errorMsg && <p className="error">{errorMsg}</p>}
        <input
          type="text"
          placeholder="Usuario"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Ingresar</button>
      </form>
    </div>
  );
}
