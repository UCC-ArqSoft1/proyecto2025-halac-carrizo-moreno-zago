// src/pages/LoginPage.tsx

import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function LoginPage() {
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
        credentials: "include", // ¡muy importante para que guarde la cookie!
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (!res.ok) {
        setErrorMsg(data.error || "Credenciales incorrectas");
        return;
      }

      // Si llegamos aquí, la cookie “token” ya está almacenada en el navegador
      navigate("/dashboard");
    } catch (err) {
      console.error("Error al loguearse:", err);
      setErrorMsg("Error de conexión con el servidor");
    }
  };

  return (
    <div style={{ display: "grid", placeItems: "center", height: "100vh" }}>
      <form
        onSubmit={handleLogin}
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "1rem",
          width: 300,
          boxShadow: "0 5px 15px rgba(0,0,0,0.1)",
          padding: "2rem",
          borderRadius: "8px",
          background: "#fff",
        }}
      >
        <h2>Iniciar Sesión</h2>
        {errorMsg && <p style={{ color: "red" }}>🔴 {errorMsg}</p>}
        <input
          type="text"
          placeholder="Usuario"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
          style={{ padding: "0.5rem", borderRadius: "4px", border: "1px solid #ccc" }}
        />
        <input
          type="password"
          placeholder="Contraseña"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          style={{ padding: "0.5rem", borderRadius: "4px", border: "1px solid #ccc" }}
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
