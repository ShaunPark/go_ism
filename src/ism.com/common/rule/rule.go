package rule

type Interface struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type DataStructure struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Data        []Data            `json:"data"`
	Lengths     []LengthFieldInfo `json:"lengths"`
	MessageType int               `json:"messageType"`
}

type Data struct {
	Id                 string
	DataStrtId         string
	MasterFieldGroupId string
	Detail             []Detail
	RecordDelimeter    string
	Master             FieldGroup
	DataIndex          int
}

type Detail struct {
	Id                    string
	Detail                FieldGroup
	RepeatCount           int
	RepeatCountDataIndex  int
	RepeatCountFieldIndex int
	GroupType             int
}

type LengthFieldInfo struct {
	LengthDataIndex   int
	LengthDetailIndex int
	LengthFieldIndex  int

	DataIndex   int
	DetailIndex int
	FieldIndex  int
	DiffValue   int
}

type FieldGroup struct {
}
