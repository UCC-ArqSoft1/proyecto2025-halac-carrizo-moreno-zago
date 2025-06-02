// src/pages/HomePage.tsx

import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export type Activity = {
  id: string;
  name: string;
  duration: number;
  intensity: string;
  trainer_id: string;
  schedule: {
    day_of_week: string;
    start_time: string;
    end_time: string;
  }[];
};

export default function HomePage() {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [search, setSearch] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost:3000/activities", {
      credentials: "include", // <— para enviar la cookie token
    })
      .then(async (res) => {
        if (res.status === 401) {
          // Si no está logueado, redirigimos al login
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
          setActivities(data);
        }
      })
      .catch((err: any) => {
        console.error("Error cargando actividades en HomePage:", err);
        setError(err.message || "Error al cargar actividades");
      });
  }, [navigate]);

  if (error) {
    return <div style={{ color: "red", textAlign: "center", marginTop: "2rem" }}>🔴 {error}</div>;
  }

  const filtered = activities.filter((a) =>
    a.name.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div style={{ padding: "2rem", fontFamily: "Segoe UI" }}>
      <h1>Actividades</h1>
      <input
        type="text"
        placeholder="Buscar por nombre..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
        style={{ padding: ".5rem", width: "100%", maxWidth: "300px", marginBottom: "1rem" }}
      />
      {filtered.map((a) => (
        <div
          key={a.id}
          style={{
            background: "#fff",
            padding: "1rem",
            marginBottom: "1rem",
            borderRadius: "8px",
            boxShadow: "0 2px 6px rgba(0,0,0,0.1)",
          }}
        >
          <h3>{a.name}</h3>
          <p>
            <strong>Entrenador:</strong> {a.trainer_id}
          </p>
          {a.schedule.length > 0 && (
            <p>
              <strong>Horario:</strong> {a.schedule[0].day_of_week} - {a.schedule[0].start_time}
            </p>
          )}
          <button onClick={() => navigate(`/actividad/${a.id}`)} style={{ marginTop: "0.5rem" }}>
            Ver más
          </button>
        </div>
      ))}
    </div>
  );
}
