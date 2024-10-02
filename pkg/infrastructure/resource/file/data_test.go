package file_test

type (
	Path = string
	URL  = string
)

type (
	Model struct {
		Ok bool
	}
	Local struct {
		Success Path
		Failure Path
	}
	Remote struct {
		Success URL
		Failure URL
	}
	Data struct {
		*Model
		*Local
		*Remote
	}
)
