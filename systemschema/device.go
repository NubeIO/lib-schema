package systemschema

import "github.com/NubeIO/lib-schema/schema"

type DeviceSchema struct {
	UUID              schema.UUID                         `json:"uuid"`
	Name              schema.Name                         `json:"name"`
	Description       schema.Description                  `json:"description"`
	Enable            schema.Enable                       `json:"enable"`
	AutoMappingEnable schema.AutoMappingEnableDefaultTrue `json:"auto_mapping_enable"`
}

func GetDeviceSchema() *DeviceSchema {
	m := &DeviceSchema{}
	schema.Set(m)
	return m
}
