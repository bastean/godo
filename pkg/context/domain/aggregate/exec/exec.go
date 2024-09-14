package exec

type (
	Command  = string
	Commands = []string
)

type Exec struct {
	Title, Description string
	Commands           []string
}

func New() *Exec {
	return &Exec{
		Commands: []string{},
	}
}
