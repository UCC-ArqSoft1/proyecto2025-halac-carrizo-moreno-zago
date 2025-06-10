import { Link } from "react-router-dom";

export default function WelcomePage() {
  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        background: "linear-gradient(135deg, #74ebd5, #9face6)",
        fontFamily: `"Segoe UI", Tahoma, Geneva, Verdana, sans-serif`
      }}
    >
      <div
        style={{
          background: "#fff",
          padding: "3rem 2.5rem",
          borderRadius: "18px",
          boxShadow: "0 8px 32px rgba(0,0,0,0.13)",
          maxWidth: 420,
          textAlign: "center"
        }}
      >
        <h1 style={{ color: "#234075", fontSize: "2.3rem", marginBottom: "1.1rem" }}>
          ¡Bienvenido/a a <span style={{ color: "#2cc4c4" }}>Gimnasio.go</span>!
        </h1>
        <p style={{ color: "#555", fontSize: "1.15rem", marginBottom: "2.5rem" }}>
          Nos alegra tenerte aquí.<br />¿Qué te gustaría hacer ahora?
        </p>
        <div style={{ display: "flex", flexDirection: "column", gap: "1.3rem" }}>
          <Link
            to="/dashboard"
            style={{
              background: "linear-gradient(90deg, #43cea2, #185a9d)",
              color: "white",
              fontWeight: "bold",
              fontSize: "1.13rem",
              padding: "1rem 0",
              border: "none",
              borderRadius: "10px",
              textDecoration: "none",
              boxShadow: "0 2px 8px rgba(67,206,162,0.10)",
              transition: "background 0.3s",
            }}
          >
            Inscribirme a una Actividad
          </Link>
          <Link
            to="/mis-actividades"
            style={{
              background: "linear-gradient(90deg, #2b5876, #4e4376)",
              color: "white",
              fontWeight: "bold",
              fontSize: "1.13rem",
              padding: "1rem 0",
              border: "none",
              borderRadius: "10px",
              textDecoration: "none",
              boxShadow: "0 2px 8px rgba(67,206,162,0.10)",
              transition: "background 0.3s",
            }}
          >
            Ver mis Actividades
          </Link>
        </div>
      </div>
    </div>
  );
}
