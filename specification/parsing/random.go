package parsing

import (
	"strings"

	"github.com/followthepattern/forgefy/specification"
	"github.com/followthepattern/forgefy/specification/defaults"
	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/types"
)

func CreateRandomRecordsFunc(t types.TypeRegistry) func(f specification.Feature) []models.Record {
	return func(f specification.Feature) (records []models.Record) {
		for i := 0; i < f.CountOfRandomValues; i++ {
			record := models.Record{}
			for _, field := range f.Fields {
				field.Value = field.RandomValue()
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
