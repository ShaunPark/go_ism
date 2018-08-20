package transform

type Length struct {
	dataIndex   int    //dataIndex // -1
	detailIndex int    //= -1;
	fieldIndex  int    //= -1;
	repeatCount int    //= -1;
	diffValue   int    //= 0;
	fieldId     string //;
	offset      int    // = -1;
	length      int    //= -1;
	lenType     int    //= Constants.TOTAL_LENGTH;
}

type Output struct {
	byteMessage  []byte
	input        Input
	dataOffset   []int
	detailOffset [][][]int
}

type ArrayInput struct {
	ByteData   []byte
	Data       [][][]byte
	Detail     [][][][][]byte
	RepeatInfo [][]int
}

type Input struct {
	bData  []byte
	Data   [][]byte
	Detail [][][][]byte

	DataOffset   []int
	DetailOffset []int

	DataLength   []int
	DetailLength [][][]int
}
