import React, { useState, useEffect } from 'react';
import axios from 'axios';
import QuizList from '../Quiz/QuizList';
import QuizDetails from '../Quiz/QuizDetails';
import QuizParticipation from '../Quiz/QuizParticipation';
import QuizCreation from '../Quiz/QuizCreation';
import QuestionAnswerCreation from '../Quiz/QuestionAnswerCreation';


function UserDashboard({ setToken }) {
  const [selectedQuiz, setSelectedQuiz] = useState(null);
  const [quizzes, setQuizzes] = useState([]);
  const [view, setView] = useState("list");

  const handleLogout = () => {
    setToken(null);
    localStorage.removeItem('token');
  };

    useEffect(() => {
      const fetchQuizzes = async () => {
        try {
          const response = await axios.get('http://localhost:8081/quizzes');
          setQuizzes(response.data || []);
        } catch (error) {
          console.error("Ошибка при загрузке викторин:", error);
          setQuizzes([]);
        }
      };
    
      fetchQuizzes();
    }, []);

    const addQuiz = (quiz) => {
      if (quiz && quiz.id) {
        setQuizzes(prevQuizzes => [...prevQuizzes, quiz]);
        setSelectedQuiz(quiz);
        setView("questions");
      } else {
        alert("Ошибка при создании викторины");
      }
    };
  
    const updateQuizList = async () => {
      try {
        const response = await axios.get('http://localhost:8081/quizzes');
        setQuizzes(response.data || []);
      } catch (error) {
        console.error("Ошибка при загрузке викторин:", error);
      }
    };

  return (
    <div className='user-dashboard-container'>

      {view === "list" && (
        <>
          <button onClick={() => setView('create')} className='create-quiz'>Создать викторину</button>
          <QuizList
            quizzes={quizzes}
            openQuiz={(quiz) => {
              setSelectedQuiz(quiz);
              setView("details");
            }}
            //createQuiz={() => setView("create")}
          />
        </>
      )}

      {view === "create" && (
        <QuizCreation setView={setView} addQuiz={addQuiz} />
      )}

      {view === "questions" && selectedQuiz && (
        <QuestionAnswerCreation quizId={selectedQuiz.id} setView={setView} />
      )}

      {selectedQuiz && view === "details" && (
        <QuizDetails
          quiz={selectedQuiz}
          startQuiz={() => setView("participation")}
          setView={setView}
        />
      )}

      {selectedQuiz && view === "participation" && (
        <QuizParticipation quizId={selectedQuiz.id} setView={setView} />
      )}

      <button onClick={handleLogout} className='out'>Выйти из аккаунта</button>

    </div>
  );
}

export default UserDashboard;