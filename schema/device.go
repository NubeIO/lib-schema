package schema

type Host struct {
	Type     string `json:"type" default:"string"`
	Title    string `json:"title" default:"host ip address"`
	Default  string `json:"default" default:"0.0.0.0"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type Ip struct {
	Type     string `json:"type" default:"string"`
	Title    string `json:"title" default:"host ip address"`
	Default  string `json:"default" default:"0.0.0.0"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type DeviceObjectId struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"object id"`
	Default  int    `json:"default" default:"2508"`
	Min      int    `json:"minLength" default:"0"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type AddressId struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"address id"`
	Default  int    `json:"default" default:"1"`
	Min      int    `json:"minLength" default:"0"`
	Max      int    `json:"maxLength" default:"255"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type AddressLength struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"address length"`
	Default  int    `json:"default" default:"1"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type NetworkNumber struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"network number"`
	Default  int    `json:"default" default:"0"`
	Min      int    `json:"minLength" default:"0"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type DeviceMac struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"device mstp mac address"`
	Default  int    `json:"default" default:"0"`
	Min      int    `json:"minLength" default:"0"`
	Max      int    `json:"maxLength" default:"255"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}
