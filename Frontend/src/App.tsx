import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Activities from "./Activities";
import CreateActivity from "./CreateActivity";
import Login from "./Login"; // 👈 importalo

function App() {
  return (
    <BrowserRouter>
      <nav style={{ padding: "1rem", display: "flex", gap: "1rem", background: "#eee" }}>
        <Link to="/login">Login</Link>
        <Link to="/activities">Actividades</Link>
        <Link to="/create">Crear</Link>
      </nav>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/activities" element={<Activities />} />
        <Route path="/create" element={<CreateActivity />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
