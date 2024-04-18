package errors

var (
	ErrNotFound           = NewErr("Not Found", 404)
	ErrUnpocessableEntity = NewErr("Unprocessable Entity", 422)
	ErrBadRequest         = NewErr("Bad Request", 400)
	ErrForbidden          = NewErr("Forbidden: Not Enough Permissions", 403)
	ErrNotAuthorized      = NewErr("Forbidden: Not Authorized", 401)
	ErrTooManyRequests    = NewErr("Too Many Requests", 429)
	ErrInternal           = NewErr("Internal Server Error", 500)
)
