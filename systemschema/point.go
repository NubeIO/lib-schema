package systemschema

import "github.com/NubeIO/lib-schema/schema"

type PointSchema struct {
	UUID        schema.UUID        `json:"uuid"`
	Name        schema.Name        `json:"name"`
	Description schema.Description `json:"description"`
	Enable      schema.Enable      `json:"enable"`
}

func GetPointSchema() *PointSchema {
	m := &PointSchema{}
	schema.Set(m)
	return m
}
