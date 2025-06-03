package customhttp

type Router struct {
	Paths *[]Path
}

func CreateRouter() (*Router, error) {
	return &Router{
		Paths: &[]Path{},
	}, nil
}

func (router *Router) RegisterPath(path string, handler func(Request) string) error {
	return nil
}
