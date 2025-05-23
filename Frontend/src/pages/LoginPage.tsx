// src/pages/LoginPage.tsx

import { useState } from "react";
import React from "react"; // También funciona si usás JSX tradicional

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      if (res.ok) {
        const data = await res.json();
        // Guardar token en localStorage o cookie
        localStorage.setItem("token", data.token);
        alert("Login exitoso");
        // Redireccionar a dashboard o similar
      } else {
        alert("Credenciales incorrectas");
      }
    } catch (err) {
      console.error("Error al loguearse:", err);
    }
  };

  return (
    <div style={{ display: "grid", placeItems: "center", height: "100vh" }}>
      <form onSubmit={handleLogin} style={{ display: "flex", flexDirection: "column", gap: "1rem", width: 300 }}>
        <h2>Iniciar Sesión</h2>
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
