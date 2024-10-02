package cases

type (
	DataRead interface {
		Run(route string) ([]byte, error)
	}
	DataDecode interface {
		Run(data []byte, target any) error
	}
)
