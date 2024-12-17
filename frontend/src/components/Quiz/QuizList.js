import React from 'react';

function QuizList({ quizzes = [], openQuiz }) {
  return (
    <div className='quiz-list-container'>
      <h2>Список викторин</h2>
      <div className="services">
        {quizzes.length > 0 ? (
          quizzes.map((quiz, index) => (
            <div className="service-card" key={index} onClick={() => openQuiz(quiz)}>
              <h3>{quiz.title}</h3>
            </div>
          ))
        ) : (
          <p>Нет доступных викторин</p> 
        )}
      </div>
    </div>
  );
}

export default QuizList;