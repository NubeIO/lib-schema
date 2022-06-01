package schema

type UUID struct {
	Type     string `json:"type" default:"string"`
	ReadOnly bool   `json:"read_only"`
}

type Name struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"2"`
	Max      int    `json:"max" default:"50"`
	Default  string `json:"default" default:""`
}

type Username struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"1"`
	Max      int    `json:"max" default:"50"`
	Default  string `json:"default" default:"admin"`
}

type Password struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Default  string `json:"default" default:""`
}

type Description struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
}

type Enable struct {
	Type     string `json:"type" default:"bool"`
	Required bool   `json:"required" default:"true"`
}

type Interface struct {
	Type     string   `json:"type" default:"string"`
	Required bool     `json:"required" default:"true"`
	Options  []string `json:"options" default:"[]"`
	Default  string   `json:"default" default:""`
	Help     string   `json:"help" default:"host network interface card, eg eth0"`
}

type Product struct {
	Type     string   `json:"type" default:"string"`
	Required bool     `json:"required" default:"true"`
	Options  []string `json:"options" default:"[\"RubixCompute\",\"RubixCompute5\",\"RubixComputeIO\",\"Edge28\",\"Nuc\",\"AllLinux\"]"`
	Default  string   `json:"default" default:""`
	Help     string   `json:"help" default:"a nube product type or a general linux server"`
}

type IP struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Default  string `json:"default" default:""`
	Help     string `json:"help" default:"ip address, eg 192.168.15.10 or nube-io.com (https:// is not needed in front of the address)"`
}

type Netmask struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Default  string `json:"default" default:"255.255.255.0"`
	Help     string `json:"help" default:"ip netmask address eg, 255.255.255.0"`
}

type Gateway struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Help     string `json:"help" default:"ip gateway address eg, 192.168.15.1"`
}

type HTTPS struct {
	Type     string `json:"type" default:"bool"`
	Required bool   `json:"required" default:"true"`
}

type Port struct {
	Type     string `json:"type" default:"int"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"2"`
	Max      int    `json:"max" default:"65535"`
	Options  int    `json:"options" default:""`
	Default  int    `json:"default" default:""`
	Help     string `json:"help" default:"ip port, eg port 8080 192.168.15.10:8080"`
}
