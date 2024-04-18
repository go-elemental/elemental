package elemental

import (
	"reflect"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchemaOptions struct {
	Collection string
	Database   string
	Connection string
}

type Field struct {
	Type     reflect.Kind
	Required bool
	Default  any
	Min      float64
	Max      float64
	Length   int64
	Regex    string
	Index    options.IndexOptions
	IndexOrder int
}