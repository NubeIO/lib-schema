package schema

type Host struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"host ip address"`
	Default string `json:"default" default:"0.0.0.0"`
}

type DeviceObjectId struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"object id"`
	Default int    `json:"default" default:"2508"`
	Min     int    `json:"minLength" default:"0"`
}

type AddressId struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"address id"`
	Default int    `json:"default" default:"1"`
	Min     int    `json:"minLength" default:"0"`
	Max     int    `json:"maxLength" default:"255"`
}

type AddressLength struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"address length"`
	Default int    `json:"default" default:"1"`
}

type NetworkNumber struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"network number"`
	Default int    `json:"default" default:"0"`
	Min     int    `json:"minLength" default:"0"`
}

type DeviceMac struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"device mstp mac address"`
	Default int    `json:"default" default:"0"`
	Min     int    `json:"minLength" default:"0"`
	Max     int    `json:"maxLength" default:"255"`
}
