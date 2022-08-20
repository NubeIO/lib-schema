package schema

type PollPriority struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"poll priority"`
	Options  []string `json:"enum" default:"[\"high\",\"normal\",\"low\"]"`
	EnumName []string `json:"enumNames" default:"[\"high\",\"normal\",\"low\"]"`
	Default  string   `json:"default" default:"normal"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type PollRate struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"poll rate"`
	Options  []string `json:"enum" default:"[\"fast\",\"normal\",\"slow\"]"`
	EnumName []string `json:"enumNames" default:"[\"fast\",\"normal\",\"slow\"]"`
	Default  string   `json:"default" default:"normal"`
	ReadOnly bool     `json:"readOnly" default:"false"`
}

type FastPollRate struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"fast poll rate (seconds)"`
	Default  int    `json:"default" default:"1"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type NormalPollRate struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"normal poll rate (seconds)"`
	Default  int    `json:"default" default:"20"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type SlowPollRate struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"slow poll rate seconds"`
	Default  int    `json:"default" default:"120"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type MaxPollRate struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"max poll rate seconds"`
	Default  int    `json:"default" default:"0.1"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type ZeroMode struct {
	Type     string `json:"type" default:"boolean"`
	Title    string `json:"title" default:"zero mode"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}
