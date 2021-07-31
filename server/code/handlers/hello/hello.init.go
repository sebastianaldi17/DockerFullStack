package hello

type Handler struct {
	helloBusiness helloBusiness
}

func New(helloBusiness helloBusiness) *Handler {
	hello := Handler{
		helloBusiness: helloBusiness,
	}
	return &hello
}
