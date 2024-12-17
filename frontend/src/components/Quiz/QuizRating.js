import React, { useState } from 'react';
import axios from 'axios';

const QuizRating = ({ quizId, onFinish, setAverageRating }) => {
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSuccess, setIsSuccess] = useState(false);

  const handleRatingChange = (newRating) => {
    console.log(`Star clicked: ${newRating}`);
    setRating(newRating);
  };

  const handleCommentChange = (e) => {
    setComment(e.target.value);
  };

  const submitRating = async () => {
    if (rating === 0) {
      alert('Пожалуйста, поставьте оценку!');
      return;
    }

    setIsSubmitting(true);
    
    try {
      const response = await axios.post(`http://localhost:8081/quizzes/${quizId}/rating`, {
        rating,
        comment,
      });

      console.log('Ответ от сервера:', response.data);

      if (response.data.success) {
        setIsSuccess(true);
        setAverageRating(rating);
        await fetchAverageRating();
        setTimeout(() => {
          onFinish();
        }, 2000);
      } else {
        alert('Спасибо за Ваш отзыв!');
      }
    } catch (error) {
      console.error("Ошибка при отправке оценки:", error);
      alert('Произошла ошибка при отправке отзыва. Пожалуйста, попробуйте снова.');
    } finally {
      setIsSubmitting(false);
    }
  };

  const fetchAverageRating = async () => {
    try {
      const response = await axios.get(`http://localhost:8081/quizzes/${quizId}/average-rating`);
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

  return (
    <div className="quiz-rating-container">
      <h2>Оцените викторину</h2>
      
      <div className="rating-stars">
        {[1, 2, 3, 4, 5].map((star) => (
          <span
            key={star}
            className={`star ${rating >= star ? 'filled' : ''}`}
            onClick={() => handleRatingChange(star)}
          >
            ★
          </span>
        ))}
      </div>

      <textarea
        placeholder="Оставьте комментарий (необязательно)"
        value={comment}
        onChange={handleCommentChange}
      ></textarea>

      <button
        onClick={submitRating}
        disabled={isSubmitting}
        className="submit-rating-btn"
      >
        {isSubmitting ? 'Отправка...' : 'Оценить'}
      </button>
    </div>
  );
};

export default QuizRating;
