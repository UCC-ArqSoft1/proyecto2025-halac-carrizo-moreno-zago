import { useEffect, useState } from "react";

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

export default function Activities() {
  const [activities, setActivities] = useState<Activity[]>([]);

  useEffect(() => {
    fetch("http://localhost:3000/activities")
      .then((res) => res.json())
      .then((data) => setActivities(data));
  }, []);

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
        color: "#000" // 👈 texto negro
      }}>
        <h2 style={{ textAlign: "center", color: "#333" }}>Listado de Actividades</h2>
        {activities.map((a) => (
          <div key={a.id} style={{ marginBottom: "1rem", borderBottom: "1px solid #ddd", paddingBottom: "1rem" }}>
            <h3>{a.name} ({a.intensity})</h3>
            <p><strong>ID:</strong> {a.id}</p>
            <p><strong>Duración:</strong> {a.duration} minutos</p>
            <p><strong>Entrenador:</strong> {a.trainer_id}</p>
            {a.schedule && a.schedule.length > 0 ? (
              <ul>
                {a.schedule.map((s, i) => (
                  <li key={i}>{s.day_of_week}: {s.start_time} - {s.end_time}</li>
                ))}
              </ul>
            ) : (
              <p><em>Sin horarios</em></p>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
