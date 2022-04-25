package api

const (
	TypeRequest uint8 = iota
	TypeResponse
)

type APIPackage struct {
	Type   uint8
	Action string
	Data   string
}

type ReqInRegiste struct {
	Name      string   `json:"name"`
	Sli1Route []string `json:"route"`
}
