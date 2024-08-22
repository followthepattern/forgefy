package parsing

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/datagenerator"
	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/defaults"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
	"github.com/google/uuid"
)

func RandomValue(t types.TypeRegistry, f models.Field) string {
	switch t.GetType(f.Type) {
	case types.Boolean:
		return fmt.Sprintf("%v", datagenerator.RandomBool())
	case types.Number:
		return fmt.Sprintf("%v", datagenerator.RandomInt())
	case types.ID:
		return uuid.NewString()
	case types.String:
		return datagenerator.String(10)
	case types.UUID:
		return uuid.NewString()
	case types.Undefined:
		return types.UNDEFINED_PLACEHOLDER
	case types.Date,
		types.DateTime:
		return datagenerator.RandomTime().Format("2006-01-02 15:04:05")
	}

	return datagenerator.String(10)
}

func CreateRandomRecordsFunc(t types.TypeRegistry) func(f specification.Feature) []models.Record {
	return func(f specification.Feature) (records []models.Record) {
		for i := 0; i < f.CountOfRandomValues; i++ {
			record := models.Record{}
			for _, field := range f.Fields {
				field.Value = RandomValue(t, field)
				record.Fields = append(record.Fields, field)
			}
			record.Fields = append(record.Fields, defaults.DefaultUserlog(defaults.AdminUser()).ToFields()...)
			records = append(records, record)
		}
		return
	}
}

func UserDefinedRecords(f specification.Feature) (records []models.Record) {
	for _, givenRecord := range f.DefinedRecords {
		record := models.Record{}
		values := strings.Split(givenRecord, ",")
		for i, value := range values {
			f.Fields[i].Value = value
			record.Fields = append(record.Fields, f.Fields[i])
		}
		record.Fields = append(record.Fields, defaults.DefaultUserlog(defaults.AdminUser()).ToFields()...)
		records = append(records, record)
	}

	return records
}

func CreateRecordsFunc(t types.TypeRegistry) func(f specification.Feature) []models.Record {
	randomRecordsFunc := CreateRandomRecordsFunc(t)

	return func(f specification.Feature) []models.Record {
		return append(UserDefinedRecords(f), randomRecordsFunc(f)...)
	}
}
