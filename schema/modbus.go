package schema

type DataType struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io type"`
	Options  []string `json:"enum" default:"[\"digital\",\"uint16\",\"int16\",\"uint32\",\"int32\",\"uint64\",\"int64\",\"float32\",\"float64\"]"`
	EnumName []string `json:"enumNames" default:"[\"digital\",\"uint16\",\"int16\",\"uint32\",\"int32\",\"uint64\",\"int64\",\"float32\",\"float64\"]"`
	Default  string   `json:"default" default:"digital"`
}

type ObjectEncoding struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"io type"`
	Options  []string `json:"enum" default:"[\"leb_lew\",\"leb_bew\",\"beb_lew\",\"beb_bew\"]"`
	EnumName []string `json:"enumNames" default:"[\"leb_lew\",\"leb_bew\",\"beb_lew\",\"beb_bew\"]"`
	Default  string   `json:"default" default:"leb_lew"`
}
