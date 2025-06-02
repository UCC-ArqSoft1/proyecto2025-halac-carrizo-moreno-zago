// src/components/AdminRoute.tsx
import React, { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";

export default function AdminRoute({ children }: { children: React.ReactNode }) {
  const [status, setStatus] = useState<"pending" | "ok" | "denied">("pending");

  useEffect(() => {
    fetch("http://localhost:3000/check-auth", {
      credentials: "include",
    })
      .then(async (res) => {
        if (res.status === 401 || res.status === 403) {
          setStatus("denied");
          return;
        }
        const data = await res.json();
        if (data.role === "admin") {
          setStatus("ok");
        } else {
          setStatus("denied");
        }
      })
      .catch(() => setStatus("denied"));
  }, []);

  if (status === "pending") {
    return <div style={{ textAlign: "center", marginTop: "2rem" }}>Cargando...</div>;
  }

  if (status === "denied") {
    return <Navigate to="/login" replace />;
  }

  return children;
}
