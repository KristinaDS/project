package usecase

type Usecase struct {
	p              Provider
	quizServiceURL string
}

func NewUsecase(p Provider, quizServiceURL string) *Usecase {
	return &Usecase{
		p:              p,
		quizServiceURL: quizServiceURL,
	}
}
