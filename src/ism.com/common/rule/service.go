package rule

type Service struct {
	Id            string
	Name          string
	SvcBlock      string
	InDstrId      string
	OutDstrId     string
	ErrDstrId     string
	ServiceType   string
	InterfaceType string
	SendRcvType   string
}
