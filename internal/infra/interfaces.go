package infra

type FluxCP interface {
	Get(path string) ([]byte, error)
}
