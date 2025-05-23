// src/components/AdminRoute.tsx
import { Navigate } from "react-router-dom";

export default function AdminRoute({ children }: { children: React.ReactNode }) {
  const token = localStorage.getItem("token");

  try {
    if (!token) return <Navigate to="/login" replace />;

    const payload = JSON.parse(atob(token.split(".")[1]));
    if (payload.role !== "admin") return <Navigate to="/dashboard" replace />;

    return <>{children}</>;
  } catch {
    return <Navigate to="/login" replace />;
  }
}
