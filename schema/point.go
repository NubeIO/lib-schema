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

type IoNumber struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io number"`
	Options  []string `json:"enum" default:"[\"UI1\",\"UI2\",\"UI3\",\"UI4\",\"UI5\",\"UI6\",\"UI7\",\"UI8\",\"UO1\",\"UO2\",\"UO3\",\"UO4\",\"UO5\",\"UO6\",\"DO1\",\"DO2\"]"`
	EnumName []string `json:"enumNames" default:"[\"UI1\",\"UI2\",\"UI3\",\"UI4\",\"UI5\",\"UI6\",\"UI7\",\"UI8\",\"UO1\",\"UO2\",\"UO3\",\"UO4\",\"UO5\",\"UO6\",\"DO1\",\"DO2\"]"`
	Default  string   `json:"default" default:"UI1"`
}

type IoType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io type"`
	Options  []string `json:"enum" default:"[\"digital\",\"voltage_dc\",\"thermistor_10k_type_2\",\"current\",\"raw\"]"`
	EnumName []string `json:"enumNames" default:"[\"digital\",\"voltage_dc\",\"thermistor_10k_type_2\",\"current\",\"raw\"]"`
	Default  string   `json:"default" default:"digital"`
}

//type Enable struct {
//	Type    string `json:"type" default:"boolean"`
//	Title   string `json:"title" default:"enable"`
//	Default bool   `json:"default" default:"true"`
//}

type ScaleEnable struct {
	Type  string `json:"type" default:"boolean"`
	Title string `json:"title" default:"scale enable"`
	//Default bool   `json:"default" default:"true"`
}

type ScaleInMin struct {
	Type    string  `json:"type" default:"number"`
	Title   string  `json:"title" default:"scale: input min"`
	Default float64 `json:"default" default:"0"`
}

type ScaleInMax struct {
	Type    string  `json:"type" default:"number"`
	Title   string  `json:"title" default:"scale: input max"`
	Default float64 `json:"default" default:"0"`
}

type ScaleOutMin struct {
	Type    string  `json:"type" default:"number"`
	Title   string  `json:"title" default:"scale: output min"`
	Default float64 `json:"default" default:"0"`
}

type ScaleOutMax struct {
	Type    string  `json:"type" default:"number"`
	Title   string  `json:"title" default:"scale: output max"`
	Default float64 `json:"default" default:"0"`
}

//ScaleEnable struct {
//Type        string `json:"type" default:"bool"`
//Required    bool   `json:"required" default:"true"`
//Default     *bool  `json:"default" default:"false"`
//DisplayName string `json:"display_name" default:"Scale/Limit Enable"`
//} `json:"scale_enable"`
//ScaleInMin struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"0"`
//DisplayName string `json:"display_name" default:"Scale: Input Min"`
//} `json:"scale_in_min"`
//ScaleInMax struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"0"`
//DisplayName string `json:"display_name" default:"Scale: Input Max"`
//} `json:"scale_in_max"`
//ScaleOutMin struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"0"`
//DisplayName string `json:"display_name" default:"Scale/Limit: Output Min"`
//} `json:"scale_out_min"`
//ScaleOutMax struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"0"`
//DisplayName string `json:"display_name" default:"Scale/Limit: Output Max"`
//} `json:"scale_out_max"`
//MultiplicationFactor struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"1"`
//DisplayName string `json:"display_name" default:"Multiplication Factor"`
//} `json:"multiplication_factor"`
//Offset struct {
//Type        string `json:"type" default:"float"`
//Required    bool   `json:"required" default:"false"`
//Default     string `json:"default" default:"0"`
//DisplayName string `json:"display_name" default:"Offset"`
//} `json:"offset"`
