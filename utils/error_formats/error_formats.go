package error_formats

import (
	"assignment-4/utils/error_utils"
	"strings"
)

func ParseError(err error) error_utils.MessageErr {

	if strings.Contains(err.Error(), "no rows in result set") {
		return error_utils.NewNotFoundError("no record found")
	} else if strings.Contains(err.Error(), "violates unique constraint") {
		return error_utils.NewBadRequest("email has been taken, try another one")
	}

	return error_utils.NewInternalServerError("something went wrong")
}
