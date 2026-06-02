package product

import "context"


type Service interface {
	//context is important for timeouts, cancellation, and passing request-scoped values
		ListProducts(ctx context.Context) ( error)
}

type srvc struct {
	// you can add dependencies here, like a database client, cache, etc.

}

func NewService() Service {
	return &srvc{

	}
}


func (s *srvc) ListProducts(ctx context.Context) (error) {
		return nil
}