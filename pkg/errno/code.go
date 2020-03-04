package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	ErrPageInvalid     = &Errno{Code: 30001, Message: "非法的页码."}
	ErrIdInvalid       = &Errno{Code: 30002, Message: "非法的ID."}
	ErrNewNotFound     = &Errno{Code: 30003, Message: "丢失的新闻."}
	ErrArticleNotFound = &Errno{Code: 30004, Message: "丢失的文章."}
)
