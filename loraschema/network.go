package loraschema

import (
	"github.com/NubeIO/lib-schema/schema"
)

type NetworkSchema struct {
	UUID                       schema.UUID                       `json:"uuid"`
	Name                       schema.Name                       `json:"name"`
	Description                schema.Description                `json:"description"`
	Enable                     schema.Enable                     `json:"enable"`
	PluginName                 schema.PluginName                 `json:"plugin_name"`
	SerialPort                 schema.SerialPortLora             `json:"serial_port"`
	SerialBaudRate             schema.SerialBaudRate             `json:"serial_baud_rate"`
	AutoMappingEnable          schema.AutoMappingEnable          `json:"auto_mapping_enable"`
	AutoMappingFlowNetworkName schema.AutoMappingFlowNetworkName `json:"auto_mapping_flow_network_name"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	m.SerialBaudRate.ReadOnly = true
	schema.Set(m)
	return m
}
