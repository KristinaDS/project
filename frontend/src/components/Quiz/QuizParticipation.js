import React, { useState, useEffect } from 'react';
import axios from 'axios';
import QuizRating from './QuizRating';

const QuizParticipation = ({ quizId, setView }) => {
  const [quiz, setQuiz] = useState(null);
  const [questions, setQuestions] = useState([]);
  const [answers, setAnswers] = useState([]);
  const [correctAnswersCount, setCorrectAnswersCount] = useState(0);
  const [isCompleted, setIsCompleted] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [correctAnswers, setCorrectAnswers] = useState([]);
  const [averageRating, setAverageRating] = useState(0);

  const handleBackToDashboard = () => {
    setView("list");
  };

  useEffect(() => {
    const fetchQuiz = async () => {
      try {
        const quizResponse = await axios.get(`http://localhost:8081/quizzes/${quizId}`);
        setQuiz(quizResponse.data);

        const questionsResponse = await axios.get(`http://localhost:8081/quizzes/${quizId}/questions`);
        setQuestions(questionsResponse.data);

        setAnswers(new Array(questionsResponse.data.length).fill(''));

        setIsLoading(false);
      } catch (error) {
        console.error('Error fetching quiz or questions:', error);
        setIsLoading(false);
      }
    };

    fetchQuiz();
  }, [quizId]);

  const handleAnswerChange = (index, value) => {
    const newAnswers = [...answers];
    newAnswers[index] = value;
    setAnswers(newAnswers);
  };

  const fetchAnswers = async (questionId) => {
    try {
      const response = await axios.get(`http://localhost:8081/questions/${questionId}/answer`);
      return response.data;
    } catch (error) {
      console.error('Ошибка при загрузке ответов:', error);
      return [];
    }
  };

  const handleSubmit = async () => {
    let count = 0;
    const correctAnswersList = [];

    for (let index = 0; index < questions.length; index++) {
      const question = questions[index];
      const userAnswer = answers[index];

      const possibleAnswers = await fetchAnswers(question.id);

      if (Array.isArray(possibleAnswers)) {
        const correctAnswer = possibleAnswers.find(answer => answer.is_correct);

      if (correctAnswer && userAnswer.toLowerCase() === correctAnswer.answer_text.toLowerCase()) {
        count++;
      }
      correctAnswersList.push({
        questionId: question.id,
        correctAnswer: correctAnswer ? correctAnswer.answer_text : 'Нет правильного ответа'
      });
    } else {
      console.error("Ответы для вопроса не являются массивом", possibleAnswers);
    }
    }
    setCorrectAnswers(correctAnswersList);
    setCorrectAnswersCount(count);
    setIsCompleted(true);
  };

  const handleRatingUpdate = (rating) => {
    setAverageRating(rating);
  };

  if (isLoading) {
    return <div>Загрузка...</div>;
  }

  if (isCompleted) {
    return (
      <div className='result'>
        <h2>Викторина завершена!</h2>
        <h3>Правильных ответов: {correctAnswersCount} из {questions.length}</h3>
        <h4>Правильные ответы:</h4>
        <ul>
          {correctAnswers.map((item, index) => (
            <li key={index}>
              <strong>Вопрос {index + 1}: </strong>{item.correctAnswer}
            </li>
          ))}
        </ul>
        <button onClick={handleBackToDashboard} className="back-but">Назад к викторинам</button>
        <QuizRating quizId={quizId} onFinish={() => setView("list")} setAverageRating={handleRatingUpdate} />
        
      </div>
    );
  }

  return (
    <div className='quiz-participation-container'>
      <h2>{quiz ? quiz.title : 'Загрузка викторины...'}</h2>
      <h4>{quiz ? quiz.description : 'Описание загружается...'}</h4>

      <form onSubmit={(e) => e.preventDefault()}>
        {questions.map((question, index) => (
          <div key={question.id} style={{ marginBottom: '25px' }}>
            <p>{question.question_text}</p>
            <input
              type="text"
              value={answers[index]}
              onChange={(e) => handleAnswerChange(index, e.target.value)}
              placeholder="Ваш ответ"
            />
          </div>
        ))}

        <button onClick={handleSubmit} className='end'>Завершить викторину</button>
        <button onClick={handleBackToDashboard} className="back-btn2">Назад к викторинам</button>
      </form>
    </div>
  );
};

export default QuizParticipation;
