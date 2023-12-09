package actions
type Action interface{
	reset(int)
}
type User struct {
	Id int
    Name string `json:"name"`
	// Text string
}