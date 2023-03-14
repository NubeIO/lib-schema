package networklinkerschema

import (
	"github.com/NubeIO/lib-schema/schema"
)

type NetworkSchema struct {
	Enable      schema.Enable `json:"enable"`
	AddressUUID struct {
		Type    string   `json:"type" default:"string"`
		Title   string   `json:"title" default:"address_uuid"`
		Options []string `json:"enum" default:"[]"`
		Help    string   `json:"help" default:"address_uuid"`
	} `json:"address_uuid"`
	AutoMappingEnable          schema.AutoMappingEnable          `json:"auto_mapping_enable"`
	AutoMappingFlowNetworkName schema.AutoMappingFlowNetworkName `json:"auto_mapping_flow_network_name"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	m.Enable.Default = true
	schema.Set(m)
	return m
}
