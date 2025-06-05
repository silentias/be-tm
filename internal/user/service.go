package user

type Service interface {
	GenerateEmailCode(email string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GenerateEmailCode(email string) error {
	registrationCode := &RegistrationCode{
		Email: email,
		Code:  999999,
	}
	return s.repo.CreateRegistrationCode(registrationCode)
}
