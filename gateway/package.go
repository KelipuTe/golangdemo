package gateway

const (
	PackageTypeReq  = 1
	PackageTypeResp = 2
)

type Package struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Type    int    `json:"type"`
	Service string `json:"service"`
	Uri     string `json:"uri"`
	Data    string `json:"data"`
}
