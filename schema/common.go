package schema

type UUID struct {
	Type     string `json:"type" default:"string"`
	Title    string `json:"title" default:"uuid"`
	ReadOnly bool   `json:"readOnly" default:"true"`
}

type Name struct {
	Type  string `json:"type" default:"string"`
	Title string `json:"title" default:"name"`
	Min   int    `json:"minLength" default:"2"`
	Max   int    `json:"maxLength" default:"50"`
}

type Username struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"username"`
	Min     int    `json:"minLength" default:"1"`
	Max     int    `json:"maxLength" default:"50"`
	Default string `json:"default" default:"admin"`
}

type Password struct {
	Type  string `json:"type" default:"string"`
	Title string `json:"title" default:"password"`
}

type Description struct {
	Type  string `json:"type" default:"string"`
	Title string `json:"title" default:"description"`
}

type Enable struct {
	Type  string `json:"type" default:"boolean"`
	Title string `json:"title" default:"enable"`
}

type Interface struct {
	Type    string   `json:"type" default:"string"`
	Title   string   `json:"title" default:"interface"`
	Options []string `json:"enum" default:"[]"`
	Help    string   `json:"help" default:"host network interface card, eg eth0"`
}

type Product struct {
	Type     string   `json:"type" default:"string"`
	Title    string   `json:"title" default:"product"`
	Options  []string `json:"enum" default:"[\"RubixCompute\",\"RubixCompute5\",\"RubixComputeIO\",\"Edge28\",\"Nuc\",\"AllLinux\"]"`
	EnumName []string `json:"enumNames" default:"[\"RubixCompute\",\"RubixCompute5\",\"RubixComputeIO\",\"Edge28\",\"Nuc\",\"AllLinux\"]"`
	Help     string   `json:"help" default:"a nube product type or a general linux server"`
}

type IP struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"ip address"`
	Default string `json:"default" default:"192.168.15.10"`
	Help    string `json:"help" default:"ip address, eg 192.168.15.10 or nube-io.com (https:// is not needed in front of the address)"`
}

type Netmask struct {
	Type    string `json:"type" default:"string"`
	Title   string `json:"title" default:"netmask"`
	Default string `json:"default" default:"255.255.255.0"`
	Help    string `json:"help" default:"ip netmask address eg, 255.255.255.0"`
}

type Gateway struct {
	Type  string `json:"type" default:"string"`
	Title string `json:"title" default:"gateway"`
	Help  string `json:"help" default:"ip gateway address eg, 192.168.15.1"`
}

type HTTPS struct {
	Type  string `json:"type" default:"boolean"`
	Title string `json:"title" default:"enable https"`
}

type Port struct {
	Type    string `json:"type" default:"number"`
	Title   string `json:"title" default:"port"`
	Min     int    `json:"minLength" default:"2"`
	Max     int    `json:"maxLength" default:"65535"`
	Default int    `json:"default" default:"1662"`
	Help    string `json:"help" default:"ip port, eg port 8080 192.168.15.10:8080"`
}
