import React, { useState } from 'react';
import axios from 'axios';

export const authenticate = async (email, password) => {
  try {
    const response = await axios.post(`http://localhost:1234/authenticate`, { email, password });
    return response.data;
  } catch (error) {
    if (error.response) {
      throw new Error(error.response.data.error || 'Ошибка при аутентификации');
    } else {
      throw new Error('Ошибка при аутентификации');
    }
  }
};

function LoginForm({ setToken, switchToRegister }) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const handleLogin = async () => {
    setErrorMessage('');
    try {
      const data = await authenticate(email, password);
      setToken(data.token);
      localStorage.setItem('token', data.token);
      alert("Успешный вход!");
    } catch (error) {
      setErrorMessage(error.message);
    }
  };

  return (
    <div className='auth-container'>
      <h2>Вход</h2>
      <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} required />
      <input type="password" placeholder="Пароль" value={password} onChange={e => setPassword(e.target.value)} required />
      <button onClick={handleLogin} className='auth-button'>Войти</button>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <p>Нет аккаунта? <span onClick={switchToRegister} className="register-link">Зарегистрироваться</span></p>
    </div>
  );
}

export default LoginForm;
