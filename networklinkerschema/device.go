package networklinkerschema

import "github.com/NubeIO/lib-schema/schema"

type DeviceSchema struct {
	Enable      schema.Enable `json:"enable"`
	Name        schema.Name   `json:"name"`
	AddressUUID struct {
		Type    string   `json:"type" default:"string"`
		Title   string   `json:"title" default:"address_uuid"`
		Options []string `json:"enum" default:"[]"`
		Help    string   `json:"help" default:"address_uuid"`
	} `json:"address_uuid"`
	AutoMappingEnable schema.AutoMappingEnableDefaultTrue `json:"auto_mapping_enable"`
}

func GetDeviceSchema() *DeviceSchema {
	m := &DeviceSchema{}
	schema.Set(m)
	return m
}
