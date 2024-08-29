package types

import "strings"

const UNDEFINED_PLACEHOLDER = "<UNDEFINED TYPE>"

type Type int

const (
	Undefined Type = iota
	Boolean
	Number
	ID
	UUID
	String
	Text
	Time
	Date
	DateTime
)

type TypeRegistry map[string]Type

func (t TypeRegistry) GetType(name string) Type {
	name = strings.ToLower(name)

	if v, ok := t[name]; ok {
		return v
	}

	return Undefined
}

var Registered = TypeRegistry{
	"undefined": Undefined,
	"boolean":   Boolean,
	"number":    Number,
	"id":        ID,
	"uuid":      UUID,
	"string":    String,
	"text":      Text,
	"time":      Time,
	"date":      Date,
	"datetime":  DateTime,
}
