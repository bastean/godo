package role

type Reader interface {
	Read(route string) ([]byte, error)
}
