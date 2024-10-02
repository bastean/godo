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

type Exec struct {
	Tasks []*Task
}
