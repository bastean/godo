package out

func Print(printer func(string), messages ...string) {
	for _, message := range messages {
		if message != "" {
			printer(message)
		}
	}
}
