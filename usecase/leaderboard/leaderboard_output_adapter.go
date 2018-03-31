package leaderboard

// OutputAdapter interface
type OutputAdapter interface {
	Handle(Output, error) (interface{}, error)
}
