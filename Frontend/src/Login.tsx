// Login.tsx

import { useState } from "react";
import "./Login.css";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMsg, setErrorMsg] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg("");

    try {
      const res = await fetch("http://localhost:3000/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",          
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (!res.ok) {
        setErrorMsg(data.error || "Login fallido");
        return;
      }

      
      navigate("/"); 
    } catch (err) {
      console.error("Error al conectar con el servidor:", err);
      setErrorMsg("Error de conexión con el servidor");
    }
  };

  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleLogin}>
        <h2>Iniciar Sesión</h2>
        {errorMsg && <p className="error">🔴 {errorMsg}</p>}
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
        <button
          type="submit"
          style={{
            padding: "0.7rem",
            borderRadius: "4px",
            border: "none",
            background: "#2193b0",
            color: "#fff",
            fontWeight: 600,
            cursor: "pointer",
          }}
        >
          Ingresar
        </button>
      </form>
    </div>
  );
}
