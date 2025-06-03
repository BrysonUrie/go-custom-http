package customhttp

type Path struct {
	PurePath   string
	PathParams map[string]string
	Children   *[]Path
}
