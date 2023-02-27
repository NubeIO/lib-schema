package schema

type DataType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"Data Type"`
	Options  []string `json:"enum" default:"[\"digital\",\"uint16\",\"int16\",\"uint32\",\"int32\",\"uint64\",\"int64\",\"float32\",\"float64\"]"`
	EnumName []string `json:"enumNames" default:"[\"digital\",\"uint16\",\"int16\",\"uint32\",\"int32\",\"uint64\",\"int64\",\"float32\",\"float64\"]"`
	Default  string   `json:"default" default:"uint16"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type ObjectEncoding struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"Object Encoding (Endianness)"`
	Options  []string `json:"enum" default:"[\"beb_lew\",\"beb_bew\",\"leb_lew\",\"leb_bew\"]"`
	EnumName []string `json:"enumNames" default:"[\"beb_lew\",\"beb_bew\",\"leb_lew\",\"leb_bew\"]"`
	Default  string   `json:"default" default:"beb_lew"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}
