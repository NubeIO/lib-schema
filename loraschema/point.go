package loraschema

import "github.com/NubeIO/lib-schema/schema"

type PointSchema struct {
	UUID        schema.UUID        `json:"uuid"`
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
	Enable      schema.Enable      `json:"enable"`
	IoType      schema.IoType      `json:"io_type"`

	ScaleEnable          schema.ScaleEnable          `json:"scale_enable"`
	ScaleInMin           schema.ScaleInMin           `json:"scale_in_min"`
	ScaleInMax           schema.ScaleInMax           `json:"scale_in_max"`
	ScaleOutMin          schema.ScaleOutMin          `json:"scale_out_min"`
	ScaleOutMax          schema.ScaleOutMax          `json:"scale_out_max"`
	Offset               schema.Offset               `json:"offset"`
	MultiplicationFactor schema.MultiplicationFactor `json:"multiplication_factor"`
	Fallback             schema.Fallback             `json:"fallback"`
}

func GetPointSchema() *PointSchema {
	m := &PointSchema{}

	m.IoType.Default = "raw"
	m.IoType.EnumName = []string{"raw", "thermistor_10k_type_2", "digital", "voltage_dc"}
	m.IoType.Options = []string{"raw", "thermistor_10k_type_2", "digital", "voltage_dc"}
	schema.Set(m)
	return m
}
