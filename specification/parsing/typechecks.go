package parsing

import (
	"slices"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateIsID(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.ID
	}
}

func CreateNoneID(t types.TypeRegistry) func([]models.Field) []models.Field {
	return func(f []models.Field) []models.Field {
		copiedFields := slices.Clone(f)

		return slices.DeleteFunc(copiedFields, func(e models.Field) bool {
			return t.GetType(e.Type) == types.ID
		})
	}
}

func CreateIsUndefined(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Undefined
	}
}

func CreateIsBoolean(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Boolean
	}
}

func CreateHasBoolean(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Boolean {
				return true
			}
		}
		return false
	}
}

func CreateIsNumber(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Number
	}
}

func CreateHasNumber(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Number {
				return true
			}
		}
		return false
	}
}

func CreateIsString(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.String
	}
}

func CreateHasString(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.String {
				return true
			}
		}
		return false
	}
}

func CreateIsText(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Text
	}
}

func CreateHasText(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Text {
				return true
			}
		}
		return false
	}
}

func CreateIsNotTimeBased(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) != types.Time &&
			t.GetType(f.Type) != types.Date &&
			t.GetType(f.Type) != types.DateTime
	}
}

func CreateIsTime(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Time
	}
}

func CreateHasTime(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Time {
				return true
			}
		}
		return false
	}
}

func CreateIsDate(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.Date
	}
}

func CreateHasDate(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.Date {
				return true
			}
		}
		return false
	}
}

func CreateIsFile(t types.TypeRegistry) func(models.Field) bool {
	return func(f models.Field) bool {
		return t.GetType(f.Type) == types.File
	}
}

func CreateHasFile(t types.TypeRegistry) func(specification.Feature) bool {
	return func(f specification.Feature) bool {
		for _, field := range f.Fields {
			if t.GetType(field.Type) == types.File {
				return true
			}
		}
		return false
	}
}
