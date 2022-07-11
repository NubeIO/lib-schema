package schema

type ObjectId struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"object id"`
	Default int    `json:"default" default:"1"`
}

type ObjectType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"object type"`
	Options  []string `json:"enum" default:"[\"analog_input\",\"analog_value\",\"analog_output\",\"binary_input\",\"binary_value\",\"binary_output\",\"multi_state_input\",\"multi_state_value\",\"multi_state_output\"]"`
	EnumName []string `json:"enumNames" default:"[\"analog input\",\"analog value\",\"analog output\",\"binary input\",\"binary value\",\"binary output\",\"multi state input\",\"multi state value\",\"multi state output\"]"`
	Default  string   `json:"default" default:"analog_value"`
}

type WriteMode struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"write mode"`
	Options  []string `json:"enum" default:"[\"read_only\",\"write_only\",\"write_once_then_read\"]"`
	EnumName []string `json:"enumNames" default:"[\"read only\",\"write only\",\"write once then read\"]"`
	Default  string   `json:"default" default:"read_only"`
}

type WritePriority struct {
	Type     string   `json:"type" default:"number"`
	Title    string   `json:"title" default:"write priority"`
	Options  []string `json:"enum" default:"[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]"`
	EnumName []string `json:"enumNames" default:"[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]"`
	Default  int      `json:"default" default:"16"`
}
