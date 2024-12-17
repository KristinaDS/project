import React, { useState } from 'react';
import axios from 'axios';

function QuizCreation({ setView, addQuiz }) {
  const [quizTitle, setQuizTitle] = useState('');
  const [quizDescription, setQuizDescription] = useState('');
  const [quizCreated, setQuizCreated] = useState(false);

  const handleBackToDashboard = () => {
    setView("list");
  };

  const createQuiz = async () => {
    if (!quizTitle.trim() || !quizDescription.trim()) {
      alert('Пожалуйста, заполните все поля');
      return;
    }
  
    try {
      const response = await axios.post('http://localhost:8081/quizzes', {
        title: quizTitle,
        description: quizDescription,
      });
  
      if (response.data && response.data.id) {
        const createdQuiz = {
          id: response.data.id,
          title: quizTitle,
          description: quizDescription,
        };
  
        addQuiz(createdQuiz);  
        setView("questions");
      } else {
        alert("Ошибка при создании викторины");
      }
    } catch (error) {
      console.error("Ошибка при создании викторины:", error);
      alert("Ошибка при создании викторины");
    }
  };

  return (
    <div className='quiz-creation-container'>
      <h2>Создание викторины</h2>
      
      <div>
        <input
          type="text"
          placeholder="Название викторины"
          value={quizTitle}
          onChange={(e) => setQuizTitle(e.target.value)}
        />
        <textarea
          placeholder="Описание"
          value={quizDescription}
          onChange={(e) => setQuizDescription(e.target.value)}
        />
        <button onClick={createQuiz} className='save-quiz-btn'>Создать викторину</button>
        <button onClick={handleBackToDashboard} className="back-btn2">Назад к викторинам</button>
      </div>

      {quizCreated && (
        <div className="notification">
          <p>Викторина успешно создана!</p>
        </div>
      )}
    </div>
  );
}

export default QuizCreation;


