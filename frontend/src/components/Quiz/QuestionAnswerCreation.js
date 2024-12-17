import React, { useState } from 'react';
import axios from 'axios';

function QuestionAnswerCreation({ quizId, setView }) {
  const [questions, setQuestions] = useState([]);
  const [newQuestionText, setNewQuestionText] = useState('');
  const [newAnswerText, setNewAnswerText] = useState('');
  const [selectedQuestionId, setSelectedQuestionId] = useState(null);
  const [showAnswerInput, setShowAnswerInput] = useState(false);
  const [showSaveQuizButton, setShowSaveQuizButton] = useState(false);
  const [isAddingQuestion, setIsAddingQuestion] = useState(true);

  const addQuestion = async () => {
    if (!newQuestionText.trim()) {
      alert('Пожалуйста, введите текст вопроса');
      return;
    }

    const quizIdNumber = parseInt(quizId, 10);
    if (isNaN(quizIdNumber)) {
      alert("Некорректный ID викторины");
      return;
    }

    try {
      const response = await axios.post(
        `http://localhost:8081/quizzes/${quizId}/questions`,
        { question_text: newQuestionText }
      );

      if (response.data && response.data.id) {
        const newQuestion = {
          id: response.data.id,
          question_text: newQuestionText,
          answers: [],
        };
        setQuestions([...questions, newQuestion]);
        setSelectedQuestionId(newQuestion.id);
        setNewQuestionText('');
        setShowAnswerInput(true);
        setIsAddingQuestion(false);
      } else {
        alert("Ошибка при добавлении вопроса");
      }
    } catch (error) {
      console.error("Ошибка при добавлении вопроса:", error);
      alert("Ошибка при добавлении вопроса");
    }
  };

  const addAnswer = async () => {
    if (!newAnswerText.trim() || !selectedQuestionId) {
      alert('Пожалуйста, выберите вопрос и введите ответ');
      return;
    }

    try {
      const response = await axios.post(
        `http://localhost:8081/questions/${selectedQuestionId}/answer`,
        { answer_text: newAnswerText, is_correct: true }
      );

      if (response.data && response.data.id) {
        setQuestions((prevQuestions) =>
          prevQuestions.map((question) =>
            question.id === selectedQuestionId
              ? {
                  ...question,
                  answers: [
                    ...question.answers,
                    {
                      id: response.data.id,
                      answer_text: newAnswerText,
                      is_correct: true,
                    },
                  ],
                }
              : question
          )
        );
        setNewAnswerText('');
        setSelectedQuestionId(null); 
        setShowAnswerInput(false);
        setIsAddingQuestion(true);
        setShowSaveQuizButton(true);
      } else {
        alert("Ошибка при добавлении ответа");
      }
    } catch (error) {
      console.error("Ошибка при добавлении ответа:", error);
      alert("Ошибка при добавлении ответа");
    }
  };

  const saveQuiz = async () => {
      alert("Викторина сохранена!");
      setView("list");
  };

  return (
    <div className='quiz-creation-container'>
      <h2>Создание вопросов</h2>

      {isAddingQuestion && (
        <>
          <input
            type="text"
            placeholder="Введите текст вопроса"
            value={newQuestionText}
            onChange={(e) => setNewQuestionText(e.target.value)}
          />
          <button onClick={addQuestion} className='save-quiz-btn'>Добавить вопрос</button>
        </>
      )}

      {showAnswerInput && (
        <div>
          <h3>Добавление ответа</h3>
          <input
            type="text"
            placeholder="Введите ответ"
            value={newAnswerText}
            onChange={(e) => setNewAnswerText(e.target.value)}
          />
          <button onClick={addAnswer} className='save-quiz-btn'>Добавить ответ</button>
        </div>
      )}

      <h3>Список вопросов и ответов</h3>
      <div className="question-list">
        {questions.map((question) => (
          <div key={question.id} className="question-block">
            <p>
              <strong className='list-qacreate'>Вопрос:</strong> {question.question_text}
            </p>
            {question.answers.map((answer) => (
              <p key={answer.id} className="answer-text">
                <strong className='list-qacreate'>Ответ:</strong> {answer.answer_text}
              </p>
            ))}
          </div>
        ))}
      </div>

      {showSaveQuizButton && (
        <button onClick={saveQuiz} className="save-quiz-btn">
          Сохранить викторину
        </button>
      )}

      <button onClick={() => setView("list")} className="back-btn2">
        Назад к викторинам
      </button>
    </div>
  );
}

export default QuestionAnswerCreation;