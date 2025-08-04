package database

import (
	"learning-companion/internal/model/quiz"
	"learning-companion/internal/model/user"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Create a dummy user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	user := user.User{
		Uuid:        uuid.New().String(),
		Username:    "testuser",
		FirstName:   "Test",
		LastName:    "User",
		Email:       "test@example.com",
		CountryCode: "US",
		Password:    string(hashedPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db.Create(&user)

	// Create Subjects
	subjects := []quiz.Subject{
		// icon is from react icon
		{Name: "Mathematics", Slug: "mathematics", Icon: "FiBook", Description: "The study of numbers, quantity, space, structure, and change."},
		{Name: "Science", Slug: "science", Icon: "FaFlask", Description: "The systematic study of the structure and behavior of the physical and natural world through observation and experiment."},
		{Name: "History", Slug: "history", Icon: "GiAncientRuins", Description: "The study of past events, particularly in human affairs."},
		{Name: "Literature", Slug: "literature", Icon: "FaBookOpen", Description: "Written works, especially those considered of superior or lasting artistic merit."},
	}
	db.Create(&subjects)

	// Create Topics
	topics := []quiz.Topic{
		{SubjectID: 1, Name: "Algebra", Slug: "algebra", Description: "A branch of mathematics that substitutes letters for numbers."},
		{SubjectID: 1, Name: "Calculus", Slug: "calculus", Description: "A branch of mathematics that deals with rates of change and accumulation of quantities."},
		{SubjectID: 2, Name: "Physics", Slug: "physics", Description: "The natural science that studies matter, its motion and behavior through space and time, and the related entities of energy and force."},
		{SubjectID: 2, Name: "Chemistry", Slug: "chemistry", Description: "The branch of science that deals with the identification of the substances of which matter is composed."},
	}
	db.Create(&topics)

	// Create Quizzes
	quizzes := []quiz.Quiz{
		{TopicID: 1, SubjectID: 1, Title: "Basic Algebra", Slug: "basic-algebra", Description: "A quiz on the fundamentals of algebra.", Difficulty: "easy", Type: "MCQ"},
		{TopicID: 2, SubjectID: 1, Title: "Derivatives", Slug: "derivatives", Description: "A quiz on the basics of differentiation.", Difficulty: "medium", Type: "MCQ"},
		{TopicID: 3, SubjectID: 2, Title: "Newton's Laws", Slug: "newtons-laws", Description: "A quiz on Newton's laws of motion.", Difficulty: "easy", Type: "MCQ"},
		{TopicID: 4, SubjectID: 2, Title: "The Periodic Table", Slug: "the-periodic-table", Description: "A quiz on the elements of the periodic table.", Difficulty: "medium", Type: "MCQ"},
	}
	db.Create(&quizzes)

	// Create Questions
	questions := []quiz.Question{
		// Algebra Questions
		{QuizID: 1, TopicID: 1, SubjectID: 1, Title: "What is 2 + 2?", Slug: "what-is-2-plus-2"},
		{QuizID: 1, TopicID: 1, SubjectID: 1, Title: "What is x in the equation x + 5 = 10?", Slug: "what-is-x-in-x-plus-5-equals-10"},

		// Calculus Questions
		{QuizID: 2, TopicID: 2, SubjectID: 1, Title: "What is the derivative of x^2?", Slug: "what-is-the-derivative-of-x-squared"},
		{QuizID: 2, TopicID: 2, SubjectID: 1, Title: "What is the derivative of sin(x)?", Slug: "what-is-the-derivative-of-sin-x"},

		// Physics Questions
		{QuizID: 3, TopicID: 3, SubjectID: 2, Title: "What is Newton's first law of motion?", Slug: "what-is-newtons-first-law-of-motion"},
		{QuizID: 3, TopicID: 3, SubjectID: 2, Title: "What is the formula for force?", Slug: "what-is-the-formula-for-force"},

		// Chemistry Questions
		{QuizID: 4, TopicID: 4, SubjectID: 2, Title: "What is the chemical symbol for water?", Slug: "what-is-the-chemical-symbol-for-water"},
		{QuizID: 4, TopicID: 4, SubjectID: 2, Title: "What is the atomic number of Helium?", Slug: "what-is-the-atomic-number-of-helium"},
	}
	db.Create(&questions)

	// Create Answers
	answers := []quiz.Answer{
		// Answers for "What is 2 + 2?"
		{QuestionID: 1, Text: "3", Correct: false},
		{QuestionID: 1, Text: "4", Correct: true},
		{QuestionID: 1, Text: "5", Correct: false},
		{QuestionID: 1, Text: "6", Correct: false},

		// Answers for "What is x in the equation x + 5 = 10?"
		{QuestionID: 2, Text: "3", Correct: false},
		{QuestionID: 2, Text: "5", Correct: true},
		{QuestionID: 2, Text: "10", Correct: false},
		{QuestionID: 2, Text: "15", Correct: false},

		// Answers for "What is the derivative of x^2?"
		{QuestionID: 3, Text: "2x", Correct: true},
		{QuestionID: 3, Text: "x", Correct: false},
		{QuestionID: 3, Text: "x^2/2", Correct: false},
		{QuestionID: 3, Text: "2", Correct: false},

		// Answers for "What is the derivative of sin(x)?"
		{QuestionID: 4, Text: "cos(x)", Correct: true},
		{QuestionID: 4, Text: "-sin(x)", Correct: false},
		{QuestionID: 4, Text: "-cos(x)", Correct: false},
		{QuestionID: 4, Text: "sin(x)", Correct: false},

		// Answers for "What is Newton's first law of motion?"
		{QuestionID: 5, Text: "An object at rest stays at rest and an object in motion stays in motion with the same speed and in the same direction unless acted upon by an unbalanced force.", Correct: true},
		{QuestionID: 5, Text: "The acceleration of an object is dependent upon two variables - the net force acting upon the object and the mass of the object.", Correct: false},
		{QuestionID: 5, Text: "For every action, there is an equal and opposite reaction.", Correct: false},
		{QuestionID: 5, Text: "Energy cannot be created or destroyed.", Correct: false},

		// Answers for "What is the formula for force?"
		{QuestionID: 6, Text: "F = ma", Correct: true},
		{QuestionID: 6, Text: "E = mc^2", Correct: false},
		{QuestionID: 6, Text: "PV = nRT", Correct: false},
		{QuestionID: 6, Text: "v = d/t", Correct: false},

		// Answers for "What is the chemical symbol for water?"
		{QuestionID: 7, Text: "H2O", Correct: true},
		{QuestionID: 7, Text: "CO2", Correct: false},
		{QuestionID: 7, Text: "O2", Correct: false},
		{QuestionID: 7, Text: "NaCl", Correct: false},

		// Answers for "What is the atomic number of Helium?"
		{QuestionID: 8, Text: "1", Correct: false},
		{QuestionID: 8, Text: "2", Correct: true},
		{QuestionID: 8, Text: "3", Correct: false},
		{QuestionID: 8, Text: "4", Correct: false},
	}
	db.Create(&answers)
}
