package schema

type TransportTypeStruct struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"network transport type"`
	Options  []string `json:"enum" default:"[\"serial\",\"ip\",\"LoRa\"]"`
	EnumName []string `json:"enumNames" default:"[\"serial\",\"ip\",\"LoRa\"]"`
	Default  string   `json:"default" default:"serial"`
}

type SerialPort struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"serial port"`
	Options  []string `json:"enum" default:"[\"/dev/ttyAMA0\",\"/dev/ttyRS485-1\",\"/dev/ttyRS485-2\",\"/data/socat/loRa1\",\"/dev/ttyUSB0\",\"/dev/ttyUSB1\",\"/dev/ttyUSB2\",\"/dev/ttyUSB3\",\"/dev/ttyUSB4\",\"/data/socat/loRa1\",\"/data/socat/serialBridge1\"]"`
	EnumName []string `json:"enumNames" default:"[\"/dev/ttyAMA0\",\"/dev/ttyRS485-1\",\"/dev/ttyRS485-2\",\"/data/socat/loRa1\",\"/dev/ttyUSB0\",\"/dev/ttyUSB1\",\"/dev/ttyUSB2\",\"/dev/ttyUSB3\",\"/dev/ttyUSB4\",\"/data/socat/loRa1\",\"/data/socat/serialBridge1\"]"`
	Default  string   `json:"default" default:"/dev/ttyAMA0"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type SerialBaudRate struct {
	Type     string   `json:"type" default:"number"`
	Title    string   `json:"title" default:"serial baud rate"`
	Options  []string `json:"enum" default:"[9600, 38400, 57600, 115200]"`
	EnumName []string `json:"enumNames" default:"[9600, 38400, 57600, 115200]"`
	//Default  int      `json:"default" default:"38400"`
	ReadOnly bool `json:"readOnly" default:"false"`
}

type SerialParity struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"serial parity"`
	Options  []string `json:"enum" default:"[\"odd\",\"even\",\"none\"]"`
	EnumName []string `json:"enumNames" default:"[\"odd\",\"even\",\"none\"]"`
	Default  string   `json:"default" default:"none"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type SerialDataBits struct {
	Type     string   `json:"type" default:"number"`
	Title    string   `json:"title" default:"serial data bit"`
	Options  []string `json:"enum" default:"[7, 8]"`
	EnumName []string `json:"enumNames" default:"[7, 8]"`
	Default  int      `json:"default" default:"8"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}
type SerialStopBits struct {
	Type     string   `json:"type" default:"number"`
	Title    string   `json:"title" default:"serial stop bit"`
	Options  []string `json:"enum" default:"[1, 2]"`
	EnumName []string `json:"enumNames" default:"[1, 2]"`
	Default  int      `json:"default" default:"1"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type SerialTimeout struct {
	Type    string  `json:"type" default:"number"`
	Title   string  `json:"title" default:"serial timeout"`
	Default float64 `json:"default" default:"0.1"`
}
