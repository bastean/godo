package exec

type Command struct {
	Name string
	Args []string
}

type Task struct {
	Title, Description string
	Success, Error     string
	Commands           []*Command
}

func New() *Task {
	return &Task{
		Commands: []*Command{},
	}
}
