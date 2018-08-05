package textbot

import (
	"errors"
)

var MissingParams = errors.New("Missing one or more parameters.")
var MustBeDataType = errors.New("Must be a Data type.")
