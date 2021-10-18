package helper

const (
	CAB_ALREADY_CREATED = "Cab already created"
	CAB_NOT_FOUND       = "Cab not found"
)

type Error struct {
	Code    int
	Message string
	Error   string
}

type BadRequest struct {
	Code    int
	Message string
	Error   string
}
