// src/pages/MyActivities.tsx
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

type Schedule = {
  day_of_week: string;
  start_time: string;
  end_time: string;
};

type Activity = {
  id: string;
  name: string;
  duration: number;
  intensity: string;
  trainer_id: string;
  schedule: Schedule[];
};

export default function MyActivities() {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [error, setError] = useState<string>("");
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost:3000/user/activities", {
      credentials: "include",
    })
      .then(async (res) => {
        if (res.status === 401) {
          // Redirigimos inmediatamente a /login
          navigate("/login");
          return null;
        }
        if (!res.ok) {
          const err = await res.json();
          throw new Error(err.error || "Error desconocido");
        }
        return res.json();
      })
      .then((data) => {
        if (data) {
          setActivities(data.activities || []);
        }
      })
      .catch((err: any) => {
        console.error("❌ Error cargando actividades inscritas:", err);
        setError(err.message || "Error al cargar actividades");
      });
  }, [navigate]);

  if (error) {
    return (
      <div style={{ color: "red", textAlign: "center", marginTop: "2rem" }}>
        🔴 {error}
      </div>
    );
  }

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh",
        width: "100vw",
        background: "linear-gradient(135deg, #74ebd5, #9face6)",
        fontFamily: `"Segoe UI", Tahoma, Geneva, Verdana, sans-serif`,
      }}
    >
      <div
        style={{
          background: "white",
          padding: "2.5rem 2rem",
          borderRadius: "12px",
          boxShadow: "0 12px 30px rgba(0, 0, 0, 0.15)",
          width: "100%",
          maxWidth: 600,
          color: "#000",
        }}
      >
        <h2 style={{ textAlign: "center", color: "#333" }}>Mis Actividades</h2>
        {activities.length === 0 ? (
          <p style={{ textAlign: "center" }}>
            <em>No estás inscrito en ninguna actividad.</em>
          </p>
        ) : (
          activities.map((a) => (
            <div
              key={a.id}
              style={{
                marginBottom: "1.5rem",
                borderBottom: "1px solid #ddd",
                paddingBottom: "1rem",
              }}
            >
              <h3>
                {a.name} ({a.intensity})
              </h3>
              <p>
                <strong>Duración:</strong> {a.duration} minutos
              </p>
              <p>
                <strong>Entrenador:</strong> {a.trainer_id}
              </p>
              {a.schedule.length > 0 ? (
                <ul>
                  {a.schedule.map((s, i) => (
                    <li key={i}>
                      🕒 {s.day_of_week}: {s.start_time} - {s.end_time}
                    </li>
                  ))}
                </ul>
              ) : (
                <p>
                  <em>Sin horarios disponibles</em>
                </p>
              )}
            </div>
          ))
        )}
      </div>
    </div>
  );
}
