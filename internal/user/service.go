package user

type Service interface {
	GenerateEmailCode(email string) error
	IsEmailAwaitVerify(email string) (bool, error)
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

func (s *service) IsEmailAwaitVerify(email string) (bool, error) {
	return s.repo.Exists(&RegistrationCode{}, "email = ?", email)
}
