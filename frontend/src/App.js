import React, { useState } from 'react';
import './styles.css';
import LoginForm from './components/Auth/LoginForm';
import RegisterForm from './components/Auth/RegisterForm';
import UserDashboard from './components/User/UserDashboard';

function App() {
  const [token, setToken] = useState(null);
  const [view, setView] = useState("login");
  const [theme, setTheme] = useState('light');

  const toggleTheme = () => {
    setTheme(prevTheme => prevTheme === 'light' ? 'dark' : 'light');
  };

  return (
    <div className={`app-container ${theme}`}>

      <button className="theme-toggle" onClick={toggleTheme}>
        {theme === 'light' ? '🌙' : '🌞'} {}
      </button>

      {!token ? (
        view === "login" ? (
          <LoginForm setToken={setToken} switchToRegister={() => setView("register")} />
        ) : (
          <RegisterForm switchToLogin={() => setView("login")} />
        )
      ) : (
        <UserDashboard setToken={setToken} />
      )}
    </div>
  );
}

export default App;