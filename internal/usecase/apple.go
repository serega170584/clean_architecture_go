package usecase

type Repository interface {
	Get()
}

type UseCase struct {
	repo Repository
}

func New(r Repository) *UseCase { return &UseCase{r} }

func (uc *UseCase) Do() {
	uc.repo.Get()
}
