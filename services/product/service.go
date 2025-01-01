package product

type Service interface {
}

type productService struct {
	r Repository
}

func New(r Repository) Service {
	return &productService{r: r}
}
