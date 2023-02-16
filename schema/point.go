package schema

type ObjectId struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"object id"`
	Default  int    `json:"default" default:"1"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type ObjectTypeModbus struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"object type"`
	Options  []string `json:"enum" default:"[\"read_coil\",\"write_coil\",\"read_discrete_input\",\"read_register\",\"read_holding\",\"write_holding\"]"`
	EnumName []string `json:"enumNames" default:"[\"read_coil\",\"write_coil\",\"read_discrete_input\",\"read_register\",\"read_holding\",\"write_holding\"]"`
	Default  string   `json:"default" default:"read_coil"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type ObjectType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"object type"`
	Options  []string `json:"enum" default:"[\"analog_input\",\"analog_value\",\"analog_output\",\"binary_input\",\"binary_value\",\"binary_output\",\"multi_state_input\",\"multi_state_value\",\"multi_state_output\"]"`
	EnumName []string `json:"enumNames" default:"[\"analog input\",\"analog value\",\"analog output\",\"binary input\",\"binary value\",\"binary output\",\"multi state input\",\"multi state value\",\"multi state output\"]"`
	Default  string   `json:"default" default:"analog_value"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type WriteMode struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"write mode"`
	Options  []string `json:"enum" default:"[\"read_only\",\"write_once\",\"write_once_read_once\",\"write_always\",\"write_once_then_read\",\"write_and_maintain\"]"`
	EnumName []string `json:"enumNames" default:"[\"read only\",\"write once\",\"write once read once\",\"write always\",\"write once then read\",\"write and maintain\"]"`
	Default  string   `json:"default" default:"read_only"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type WritePriority struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"write priority"`
	Min      int    `json:"minLength" default:"1"`
	Max      int    `json:"maxLength" default:"16"`
	Default  int    `json:"default" default:"16"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type IoNumber struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io number"`
	Options  []string `json:"enum" default:"[\"UI1\",\"UI2\",\"UI3\",\"UI4\",\"UI5\",\"UI6\",\"UI7\",\"UI8\",\"UO1\",\"UO2\",\"UO3\",\"UO4\",\"UO5\",\"UO6\",\"DO1\",\"DO2\"]"`
	EnumName []string `json:"enumNames" default:"[\"UI1\",\"UI2\",\"UI3\",\"UI4\",\"UI5\",\"UI6\",\"UI7\",\"UI8\",\"UO1\",\"UO2\",\"UO3\",\"UO4\",\"UO5\",\"UO6\",\"DO1\",\"DO2\"]"`
	Default  string   `json:"default" default:"UI1"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type IoType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io type"`
	Options  []string `json:"enum" default:"[\"digital\",\"voltage_dc\",\"thermistor_10k_type_2\",\"current\",\"raw\"]"`
	EnumName []string `json:"enumNames" default:"[\"digital\",\"voltage_dc\",\"thermistor_10k_type_2\",\"current\",\"raw\"]"`
	Default  string   `json:"default" default:"digital"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type ScaleEnable struct {
	Type     string `json:"type" default:"boolean"`
	Title    string `json:"title" default:"scale enable"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type ScaleInMin struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"scale: input min"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type ScaleInMax struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"scale: input max"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type ScaleOutMin struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"scale: output min"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type ScaleOutMax struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"scale: output max"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type MultiplicationFactor struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"multiplication factor"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type Offset struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"offset"`
	Default  float64 `json:"default" default:"0"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type Fallback struct {
	Type     string   `json:"type" default:"number"`
	Title    string   `json:"title" default:"fallback"`
	Default  *float64 `json:"default" default:""`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type Decimal struct {
	Type     string  `json:"type" default:"number"`
	Title    string  `json:"title" default:"decimal"`
	Default  float64 `json:"default" default:"2"`
	ReadOnly bool    `json:"readOnly" default:"false"`
}

type HistoryEnable struct {
	Type    string `json:"type" default:"boolean"`
	Title   string `json:"title" default:"history enable"`
	Default bool   `json:"default" default:"false"`
}

type HistoryType struct {
	Type    string   `json:"type" default:"string"`
	Title   string   `json:"title" default:"history type"`
	Options []string `json:"enum" default:"[\"COV\",\"INTERVAL\",\"COV_AND_INTERVAL\"]"`
	Default string   `json:"default" default:"INTERVAL"`
}

type HistoryInterval struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"history interval"`
	Default *int   `json:"default" default:"15"`
}
