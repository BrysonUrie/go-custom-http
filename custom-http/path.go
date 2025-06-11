package customhttp

type Path struct {
	PurePath string
	Handler  func(*Request) (*Response, error)
	// PathParams map[string]string
	// Children   *[]Path
}
