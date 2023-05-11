package networklinkerschema

import (
	"github.com/NubeIO/lib-schema/schema"
)

type NetworkSchema struct {
	Enable      schema.Enable `json:"enable"`
	Name        schema.Name   `json:"name"`
	AddressUUID struct {
		Type    string               `json:"type" default:"string"`
		Title   string               `json:"title" default:"Networks"`
		Options []schema.OptionOneOf `json:"oneOf"`
		Help    string               `json:"help" default:"Select the network pair"`
	} `json:"address_uuid"`
	AutoMappingEnable          schema.AutoMappingEnable          `json:"auto_mapping_enable"`
	AutoMappingFlowNetworkName schema.AutoMappingFlowNetworkName `json:"auto_mapping_flow_network_name"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	schema.Set(m)
	return m
}
