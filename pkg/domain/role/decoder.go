package role

type Decoder interface {
	Decode(data []byte, target any) error
}
