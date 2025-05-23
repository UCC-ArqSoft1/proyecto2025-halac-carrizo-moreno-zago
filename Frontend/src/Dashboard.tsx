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

export default function Dashboard() {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [search, setSearch] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
      return;
    }
  
    fetch("http://localhost:3000/activities")
      .then((res) => res.json())
      .then((data) => {
        console.log("📦 Actividades cargadas desde backend:", data); // 👀
        setActivities(data);
      })
      .catch((err) => {
        console.error("❌ Error al obtener actividades:", err);
      });
  }, [navigate]);

  const filtered = activities.filter((a) =>
    a.name.toLowerCase().includes(search.toLowerCase())
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
        <h2 style={{ textAlign: "center", color: "#333" }}>Actividades Disponibles</h2>

        <input
          type="text"
          placeholder="Buscar por nombre..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          style={{ padding: ".5rem", width: "100%", marginBottom: "1rem", borderRadius: "8px", border: "1px solid #ccc" }}
        />

        {filtered.length === 0 && <p>No se encontraron actividades.</p>}

        {filtered.map((a) => (
          <div key={a.id} style={{ marginBottom: "1rem", borderBottom: "1px solid #ddd", paddingBottom: "1rem" }}>
            <h3>{a.name} ({a.intensity})</h3>
            <p><strong>Duración:</strong> {a.duration} minutos</p>
            <p><strong>Entrenador:</strong> {a.trainer_id}</p>
            {a.schedule.length > 0 && (
              <ul>
                {a.schedule.map((s, i) => (
                  <li key={i}>🕓 {s.day_of_week}: {s.start_time} - {s.end_time}</li>
                ))}
              </ul>
            )}
            <button onClick={() => navigate(`/actividad/${a.id}`)} style={{
              marginTop: "0.5rem",
              background: "#2193b0",
              color: "white",
              border: "none",
              padding: "0.5rem 1rem",
              borderRadius: "5px",
              cursor: "pointer"
            }}>
              Ver más
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
