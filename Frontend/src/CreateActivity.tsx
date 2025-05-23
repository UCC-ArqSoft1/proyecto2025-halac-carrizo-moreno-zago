import { useState } from "react";

export default function CreateActivity() {
  const [activity, setActivity] = useState({
    id: "",
    name: "",
    duration: 0,
    intensity: "",
    trainer_id: "",
    schedule: [
      {
        day_of_week: "",
        start_time: "",
        end_time: "",
      },
    ],
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    const parsedValue = name === "duration" ? parseInt(value) : value;
    setActivity({ ...activity, [name]: parsedValue });
  };

  const handleScheduleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    const newSchedule = [...activity.schedule];
    newSchedule[0] = { ...newSchedule[0], [name]: value };
    setActivity({ ...activity, schedule: newSchedule });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await fetch("http://localhost:3000/activities", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(activity),
    });

    if (res.ok) {
      alert("Actividad creada con éxito");
      setActivity({
        id: "",
        name: "",
        duration: 0,
        intensity: "",
        trainer_id: "",
        schedule: [{ day_of_week: "", start_time: "", end_time: "" }],
      });
    } else {
      alert("Error al crear actividad");
    }
  };

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
      <form
        onSubmit={handleSubmit}
        style={{
          background: "white",
          padding: "2.5rem 2rem",
          borderRadius: "12px",
          boxShadow: "0 12px 30px rgba(0, 0, 0, 0.15)",
          width: "100%",
          maxWidth: 360,
          display: "flex",
          flexDirection: "column",
          gap: "1.2rem"
        }}
      >
        <h2 style={{ textAlign: "center", color: "#333" }}>Crear Actividad</h2>
        <input name="id" placeholder="ID" value={activity.id} onChange={handleChange} required />
        <input name="name" placeholder="Nombre" value={activity.name} onChange={handleChange} required />
        <input name="duration" type="number" placeholder="Duración (min)" value={activity.duration} onChange={handleChange} required />
        <select name="intensity" value={activity.intensity} onChange={handleChange} required>
          <option value="">Intensidad</option>
          <option value="low">Baja</option>
          <option value="medium">Media</option>
          <option value="high">Alta</option>
        </select>
        <input name="trainer_id" placeholder="ID del entrenador" value={activity.trainer_id} onChange={handleChange} required />
        <input name="day_of_week" placeholder="Día de la semana" value={activity.schedule[0].day_of_week} onChange={handleScheduleChange} required />
        <input name="start_time" placeholder="Hora de inicio (HH:MM)" value={activity.schedule[0].start_time} onChange={handleScheduleChange} required />
        <input name="end_time" placeholder="Hora de fin (HH:MM)" value={activity.schedule[0].end_time} onChange={handleScheduleChange} required />
        <button type="submit"
          style={{
            background: "linear-gradient(135deg, #74ebd5, #9face6)",
            color: "white",
            padding: "0.9rem",
            border: "none",
            borderRadius: "8px",
            cursor: "pointer",
            fontWeight: 600,
            transition: "transform 0.2s ease"
          }}
        >Crear Actividad</button>
      </form>
    </div>
  );
}
