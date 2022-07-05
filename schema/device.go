package schema

type Host struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"host ip address"`
	Default string `json:"default" default:"192.168.15.10"`
}

type DeviceObjectId struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"object id"`
	Default int    `json:"default" default:"1"`
}

type NetworkNumber struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"network number"`
	Default int    `json:"default" default:"0"`
}

type DeviceMac struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"device mstp mac address"`
	Default int    `json:"default" default:"1"`
}

type Segmentation struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"device segmentation"`
	Options  []string `json:"enum" default:"[\"segmentation_both\",\"no_segmentation\",\"segmentation_transmit\",\"segmentation_receive\"]"`
	EnumName []string `json:"enumNames" default:"[\"segmentation_both\",\"no_segmentation\",\"segmentation_transmit\",\"segmentation_receive\"]"`
}

type MaxADPU struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"device max-adpu"`
	Options  []string `json:"enum" default:"[50, 128, 206, 480, 1024, 1476]"`
	EnumName []string `json:"enumNames" default:"[50, 128, 206, 480, 1024, 1476]"`
}
