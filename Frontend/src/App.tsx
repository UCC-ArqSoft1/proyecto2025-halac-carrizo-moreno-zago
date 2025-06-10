import { BrowserRouter, Routes, Route, Link, useLocation } from "react-router-dom";
import CreateActivity from "./CreateActivity";
import Login from "./Login";
import Dashboard from "./Dashboard";
import ActivityDetail from "./pages/ActivityDetail";
import AdminRoute from "./components/AdminRoute";
import MyActivities from "./pages/MyActivities";
import WelcomePage from "./pages/WelcomePage"; 

function Navbar() {
  const location = useLocation();
  const routes = [
    { to: "/login", label: "Login" },
    { to: "/", label: "Inicio" },
    { to: "/dashboard", label: "Dashboard" },
    { to: "/create", label: "Crear" },
    { to: "/mis-actividades", label: "Mis Actividades" },
  ];
  return (
    <nav
      style={{
        padding: "0.8rem 1.5rem",
        display: "flex",
        gap: "1.3rem",
        alignItems: "center",
        background: "linear-gradient(90deg, #43cea2 0%, #185a9d 100%)",
        boxShadow: "0 2px 14px rgba(60,80,120,0.08)",
        borderRadius: "0 0 18px 18px",
        marginBottom: "2rem",
        overflowX: "auto",
      }}
    >
      <span
        style={{
          fontWeight: "bold",
          fontSize: "1.32rem",
          color: "#fff",
          letterSpacing: "1px",
          marginRight: "2.5rem",
          textShadow: "0 2px 6px rgba(30,30,60,0.07)"
        }}
      >
        Gimnasio.go
      </span>
      {routes.map(({ to, label }) => (
        <Link
          key={to}
          to={to}
          style={{
            color: location.pathname === to ? "#fff" : "rgba(255,255,255,0.85)",
            background: location.pathname === to ? "rgba(0,0,0,0.13)" : "transparent",
            borderRadius: "8px",
            padding: "0.45rem 1.1rem",
            fontWeight: location.pathname === to ? "bold" : 500,
            fontSize: "1.08rem",
            textDecoration: "none",
            boxShadow: location.pathname === to ? "0 2px 10px rgba(67,206,162,0.08)" : "none",
            transition: "background 0.18s, color 0.15s, box-shadow 0.17s"
          }}
          onMouseOver={e => (e.currentTarget.style.background = "rgba(44,196,196,0.18)")}
          onMouseOut={e => (e.currentTarget.style.background = location.pathname === to ? "rgba(0,0,0,0.13)" : "transparent")}
        >
          {label}
        </Link>
      ))}
    </nav>
  );
}

function App() {
  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<WelcomePage />} />
        <Route path="/login" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route
          path="/create"
          element={
            <AdminRoute>
              <CreateActivity />
            </AdminRoute>
          }
        />
        <Route path="/actividad/:id" element={<ActivityDetail />} />
        <Route path="/mis-actividades" element={<MyActivities />} />
        <Route path="*" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
