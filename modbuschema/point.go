package modbuschema

import "github.com/NubeIO/lib-schema/schema"

type PointSchema struct {
	UUID             schema.UUID             `json:"uuid"`
	Name             schema.Name             `json:"name"`
	Description      schema.Description      `json:"description"`
	Enable           schema.Enable           `json:"enable"`
	ObjectTypeModbus schema.ObjectTypeModbus `json:"object_type"`
	DataType         schema.DataType         `json:"data_type"`
	WriteMode        schema.WriteMode        `json:"write_mode"`
	AddressId        schema.AddressId        `json:"address_id"`
	AddressLength    schema.AddressLength    `json:"address_length"`
	PollPriority     schema.PollPriority     `json:"poll_priority"`
	PollRate         schema.PollRate         `json:"poll_rate"`

	ScaleEnable          schema.ScaleEnable          `json:"scale_enable"`
	ScaleInMin           schema.ScaleInMin           `json:"scale_in_min"`
	ScaleInMax           schema.ScaleInMax           `json:"scale_in_max"`
	ScaleOutMin          schema.ScaleOutMin          `json:"scale_out_min"`
	ScaleOutMax          schema.ScaleOutMax          `json:"scale_out_max"`
	Offset               schema.Offset               `json:"offset"`
	MultiplicationFactor schema.MultiplicationFactor `json:"multiplication_factor"`
	Fallback             schema.Fallback             `json:"fallback"`
	Decimal              schema.Decimal              `json:"decimal"`
}

func GetPointSchema() *PointSchema {
	m := &PointSchema{}
	schema.Set(m)
	return m
}
