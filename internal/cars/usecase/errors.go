package usecase

type Err struct {
	err  string
	code int
}

func (e *Err) Error() string {
	return e.err
}

func (e *Err) Code() int {
	return e.code
}

var (
	ErrNotFound = Err{err: "Not Found", code: 404}
)
