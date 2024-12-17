import React, { useState } from 'react';
import axios from 'axios';

export const register = async (username, email, password) => {
  try {
    const response = await axios.post(`http://localhost:1234/register`, { username, email, password });
    return response.data;
  } catch (error) {
    if (error.response) {
      throw new Error(error.response.data.error || 'Ошибка при регистрации');
    } else {
      throw new Error('Ошибка при регистрации');
    }
  }
};

function RegisterForm({ switchToLogin }) {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');

  const handleRegister = async () => {
    setLoading(true);
    setErrorMessage('');

    if (username.trim() === '') {
      setErrorMessage('Имя пользователя не может состоять только из пробелов');
      setLoading(false);
      return;
    }
    if (email.trim() === '') {
      setErrorMessage('Email не может состоять только из пробелов');
      setLoading(false);
      return;
    }
    if (password.trim() === '') {
      setErrorMessage('Пароль не может состоять только из пробелов');
      setLoading(false);
      return;
    }

    const containsSpaces = (str) => /\s/.test(str); // Проверка на наличие пробела внутри строки

    if (containsSpaces(username)) {
      setErrorMessage('Имя пользователя не может содержать пробелы');
      setLoading(false);
      return;
    }
    if (containsSpaces(email)) {
      setErrorMessage('Email не может содержать пробелы');
      setLoading(false);
      return;
    }
    if (containsSpaces(password)) {
      setErrorMessage('Пароль не может содержать пробелы');
      setLoading(false);
      return;
    }

    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!emailPattern.test(email)) {
      setErrorMessage('Неверный формат email');
      setLoading(false);
      return;
    }

    try {
      await register(username, email, password);
      alert("Успешная регистрация!");
      switchToLogin();
    } catch (error) {
      setErrorMessage(error.message);
      alert("Ошибка регистрации")
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className='auth-container'>
      <h2>Регистрация</h2>
      <input
        type="text"
        placeholder="Имя пользователя"
        value={username}
        onChange={(e) => setUsername(e.target.value.trim())}
        required
      />
      <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value.trim())} required />
      <input type="password" placeholder="Пароль" value={password} onChange={e => setPassword(e.target.value.trim())} required />
      <button onClick={handleRegister} className='auth-button'>Зарегистрироваться</button>
      {errorMessage && <p className="error-message">{errorMessage}</p>}
      <p>
        Уже есть аккаунт?{' '}
        <span onClick={switchToLogin} className='login-link'>
          Войти
        </span>
      </p>
    </div>
  );
}

export default RegisterForm;
