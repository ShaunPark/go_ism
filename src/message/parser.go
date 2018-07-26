package message

type ParserInterface struct {
}

type Parser interface {
	GetInterfaceId(msg string) string
	IsSync(msg string) bool
	GetGID(msg string) string
}
