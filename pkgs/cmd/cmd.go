package cmd

// Represents a `Command` in the design pattern of the same name.
//
// Fire and forget.
type Command interface {
	Execute()
}
