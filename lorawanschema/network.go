package lorawanschema

import (
	"github.com/NubeIO/lib-schema/schema"
)

type NetworkSchema struct {
	UUID          schema.UUID          `json:"uuid"`
	Name          schema.Name          `json:"name"`
	Description   schema.Description   `json:"description"`
	Enable        schema.Enable        `json:"enable"`
	PluginName    schema.PluginName    `json:"plugin_name"`
	HistoryEnable schema.HistoryEnable `json:"history_enable"`
}

func GetNetworkSchema() *NetworkSchema {
	m := &NetworkSchema{}
	schema.Set(m)
	return m
}
