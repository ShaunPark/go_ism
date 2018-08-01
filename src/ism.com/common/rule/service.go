package rule

type Service struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	SvcBlock      string `json:"svcBlock"`
	InDstrId      string `json:"inDstrId"`
	OutDstrId     string `json:"outDstrId"`
	ErrDstrId     string `json:"errDstrId"`
	ServiceType   string `json:"svcType"`
	InterfaceType string `json:"infType"`
	SendRcvType   string `json:"srType"`
}
