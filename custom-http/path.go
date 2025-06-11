package customhttp

type Path struct {
	PurePath string
	Handler  func(*Request) string
	// PathParams map[string]string
	// Children   *[]Path
}
