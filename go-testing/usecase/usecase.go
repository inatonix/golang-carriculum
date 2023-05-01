package usecase

type TaskUsecase interface {
	CreateTask(title string) (int, error)
}

type taskUsecase struct {
}

func NewTaskUsecase() TaskUsecase {
	return &taskUsecase{}
}

func (u *taskUsecase) CreateTask(title string) (int, error) {
	return 0, nil
}
