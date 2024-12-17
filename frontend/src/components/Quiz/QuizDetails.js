import React, { useState, useEffect } from 'react';
import axios from 'axios';
import QuizRating from './QuizRating';

function QuizDetails({ quiz, startQuiz, setView }) {
  const [questions, setQuestions] = useState([]);
  const [averageRating, setAverageRating] = useState(0);
  const [comments, setComments] = useState([]);

  const handleBackToList = () => {
    setView("list");
  };

  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        const response = await axios.get(`http://localhost:8081/quizzes/${quiz.id}/questions`);
        setQuestions(response.data);
      } catch (error) {
        console.error("Ошибка при загрузке вопросов:", error);
      }
    };

    if (quiz) {
      fetchQuestions();
    }
  }, [quiz]);

  useEffect(() => {
    const fetchAverageRating = async () => {
      try {
        const response = await axios.get(`http://localhost:8081/quizzes/${quiz.id}/average-rating`);
        console.log("Ответ от сервера при получении среднего рейтинга:", response.data);
        if (response.data && typeof response.data.average_rating === 'number') {
          setAverageRating(response.data.average_rating);
        } else {
          setAverageRating(0);
        }
      } catch (error) {
        console.error("Ошибка при получении среднего рейтинга:", error);
      }
    };

    if (quiz) {
      fetchAverageRating();
    }
  }, [quiz]);

  useEffect(() => {
    if (averageRating !== 0) {
      console.log('Средний рейтинг обновлен:', averageRating);
    }
  }, [averageRating]);

  useEffect(() => {
    const fetchComments = async () => {
      try {
        const response = await axios.get(`http://localhost:8081/quizzes/${quiz.id}/comments`);
        console.log("Комментарии от сервера:", response.data);
        if (Array.isArray(response.data)) {
          setComments(response.data);
        } else {
          setComments([]);
        }
      } catch (error) {
        console.error("Ошибка при получении комментариев:", error);
      }
    };

    if (quiz) {
      fetchComments();
    }
  }, [quiz]);

  return (
    <div className='quiz-details-container'>
      <h2>{quiz.title}</h2>
      <h4>{quiz.description}</h4>
      <h3>Количество вопросов: {questions.length}</h3>
      <h4>Средний рейтинг: {averageRating ? averageRating.toFixed(1) : '0.0'} / 5</h4>
      <button onClick={startQuiz} className="start-but">Начать викторину</button>
      <button onClick={handleBackToList} className="back-but">Назад к викторинам</button>
      <QuizRating quizId={quiz.id} onFinish={() => setView("list")} setAverageRating={setAverageRating} />
      <div className="comments-section">
        <h3>Комментарии:</h3>
        {comments.length === 0 ? (
          <p>Нет комментариев.</p>
        ) : (
          comments.map((comment, index) => (
            <div key={index} className="comment">
              <p><strong>Комментарий:</strong></p>
              <p>{comment.comment}</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default QuizDetails;
