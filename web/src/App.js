import './App.css';
import {BrowserRouter, Routes, Route, Link} from "react-router-dom";
import MainPage from './pages/MainPage';
import GetByPhone from "./pages/GetByPhone";
import GetByAddress from "./pages/GetByAddress";
import GetByName from "./pages/GetByName";

function App() {

  return (
    <BrowserRouter>
        <div className="navbar">
            <Link to="/get_by_phone">Найти по телефону</Link>
            <Link to="/get_by_address">Найти по адресу</Link>
            <Link to="/get_by_name">Найти по имени</Link>
        </div>

        <Routes>
            <Route path="/" element={<MainPage />} />
            <Route path="/get_by_phone" element={<GetByPhone /> } />
            <Route path="/get_by_address" element={<GetByAddress /> } />
            <Route path="/get_by_name" element={<GetByName /> } />
        </Routes>
    </BrowserRouter>
  );
}

export default App;
