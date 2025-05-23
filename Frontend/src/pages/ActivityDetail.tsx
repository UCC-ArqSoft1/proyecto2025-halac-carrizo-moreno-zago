import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import type { Activity } from "./HomePage";

export default function ActivityDetail() {
  const { id } = useParams();
  const [activity, setActivity] = useState<Activity | null>(null);
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch(`http://localhost:3000/activities/${id}`)
      .then((res) => res.json())
      .then((data) => {
        console.log("🔍 Detalle de actividad:", data);
        setActivity(data);
      })
      .catch((err) => {
        console.error("❌ Error al obtener detalle de actividad:", err);
      });
  }, [id]);

  const handleInscription = async () => {
    try {
      const res = await fetch(`http://localhost:3000/activities/${id}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });
      if (res.ok) {
        setMessage("¡Inscripción exitosa!");
      } else {
        setMessage("Error al inscribirse en la actividad.");
      }
    } catch (err) {
      setMessage("Error de conexión con el servidor.");
    }
  };

  if (!activity) return (
    <div style={{ display: "grid", placeItems: "center", minHeight: "100vh", fontFamily: "Segoe UI" }}>
      <p>Cargando actividad...</p>
    </div>
  );

  return (
    <div style={{
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      minHeight: "100vh",
      width: "100vw",
      background: "linear-gradient(135deg, #74ebd5, #9face6)",
      fontFamily: `"Segoe UI", Tahoma, Geneva, Verdana, sans-serif`
    }}>
      <div style={{
        background: "white",
        padding: "2.5rem 2rem",
        borderRadius: "12px",
        boxShadow: "0 12px 30px rgba(0, 0, 0, 0.15)",
        width: "100%",
        maxWidth: 500,
        color: "#000"
      }}>
        <h2 style={{ textAlign: "center", color: "#333" }}>{activity.name}</h2>
        <p><strong>Intensidad:</strong> {activity.intensity}</p>
        <p><strong>Duración:</strong> {activity.duration} minutos</p>
        <p><strong>Entrenador:</strong> {activity.trainer_id}</p>

        {activity.schedule.length > 0 ? (
          <>
            <p><strong>Horarios:</strong></p>
            <ul>
              {activity.schedule.map((s, i) => (
                <li key={i}>🕒 {s.day_of_week}: {s.start_time} - {s.end_time}</li>
              ))}
            </ul>
          </>
        ) : (
          <p><em>Sin horarios disponibles.</em></p>
        )}

        <button
          onClick={handleInscription}
          style={{
            marginTop: "1rem",
            background: "#2193b0",
            color: "white",
            border: "none",
            padding: "0.7rem 1.2rem",
            borderRadius: "8px",
            fontWeight: "bold",
            cursor: "pointer"
          }}
        >
          Inscribirse
        </button>

        {message && (
          <p style={{
            marginTop: "1rem",
            fontWeight: "bold",
            color: message.includes("Error") || message.includes("conexión") ? "red" : "green"
          }}>
            {message}
          </p>
        )}
      </div>
    </div>
  );
}
