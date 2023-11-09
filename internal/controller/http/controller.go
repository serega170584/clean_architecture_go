package http

type UseCase interface {
	Do()
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{uc: uc}
}

func (c *Controller) Serve() {
	c.uc.Do()
}
