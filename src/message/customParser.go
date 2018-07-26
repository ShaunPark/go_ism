package message

type MyParser struct {
	ParserInterface
}

const GidOffset int = 6
const GidLength int = 4
const InfIdOffset int = 0
const InfIdLength int = 6
const SyncOffset int = 10
const SyncLength int = 1

func (parser *MyParser) IsSync(msg string) bool {
	sync := msg[SyncOffset : SyncOffset+SyncLength]
	if sync == "S" || sync == "s" {
		return true
	}
	return false
}
func (parser *MyParser) GetInterfaceId(msg string) string {
	return msg[InfIdOffset : InfIdOffset+InfIdLength]
}

func (parser *MyParser) GetGID(msg string) string {
	return msg[GidOffset : GidOffset+GidLength]
}
