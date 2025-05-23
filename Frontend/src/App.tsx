import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import CreateActivity from "./CreateActivity";
import Login from "./Login";
import Dashboard from "./Dashboard";
import ActivityDetail from "./pages/ActivityDetail";
import MyActivities from "./pages/MyActivities";
import AdminRoute from "./components/AdminRoute"; // 👈 importalo

function App() {
  return (
    <BrowserRouter>
      <nav style={{ padding: "1rem", display: "flex", gap: "1rem", background: "#eee" }}>
        <Link to="/login">Login</Link>
        <Link to="/dashboard">Dashboard</Link>
        <Link to="/create">Crear</Link>
        <Link to="/mis-actividades">Mis Actividades</Link>
      </nav>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/create" element={
          <AdminRoute>
            <CreateActivity />
          </AdminRoute>
        } />
        <Route path="/actividad/:id" element={<ActivityDetail />} />
        <Route path="/mis-actividades" element={<MyActivities />} />
        <Route path="*" element={<Login />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
