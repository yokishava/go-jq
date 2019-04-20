package main

import (
	"fmt"
)

//Error  is custom error
type Error struct {
	Message string
}

func (err *Error) Error() string {
	return fmt.Sprintf("ERROR: %s", err.Message)
}
