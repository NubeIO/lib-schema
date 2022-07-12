package schema

type FastPollRate struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"fast poll rate (seconds)"`
	Default int    `json:"default" default:"1"`
}

type NormalPollRate struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"normal poll rate (seconds)"`
	Default int    `json:"default" default:"20"`
}

type SlowPollRate struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"slow poll rate seconds"`
	Default int    `json:"default" default:"120"`
}

type MaxPollRate struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"max poll rate seconds"`
	Default int    `json:"default" default:"0.1"`
}

type ZeroMode struct {
	Type  string `json:"type" default:"zero mode"`
	Title string `json:"title" default:"enable"`
}
