package specification

import (
	"fmt"
	"strings"

	"github.com/followthepattern/forgefy/specification/models"
	"github.com/followthepattern/forgefy/specification/naming"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Feature struct {
	App
	FeatureName    string         `yaml:"name"`
	Fields         []models.Field `yaml:"fields"`
	DefinedRecords []string       `yaml:"records"`
}

func (f Feature) FeatureNameCamelCase() string {
	return naming.ToLowerCamelCase(f.FeatureName)
}

func (f Feature) FeatureNameCamelCaseUpper() string {
	return naming.ToUpperCamelCase(f.FeatureName)
}

func (f Feature) FeatureHumanReadableName() string {
	return naming.CapitalizeFirst(f.FeatureName)
}

func (f Feature) IDField() models.Field {
	return f.Fields[0]
}

func (f Feature) FeatureNameDir() string {
	return strings.ToLower(naming.ToSnakeCase(f.FeatureName))
}

func (f Feature) Validate() error {
	for _, data := range f.DefinedRecords {
		line := strings.Split(data, ",")
		if len(line) != len(f.Fields) {
			return fmt.Errorf("given record (%s) is invalid, the length of the list doesn't equal to the number of fields", data)
		}
	}

	return validation.ValidateStruct(&f,
		validation.Field(&f.FeatureName, validation.Required),
		validation.Field(&f.Fields, validation.Required),
	)
}

func (f Feature) RandomRecords() (records []models.Record) {
	for i := 0; i < f.CountOfRandomValues; i++ {
		record := models.Record{}
		for _, field := range f.Fields {
			field.Value = field.RandomValue()
			record.Fields = append(record.Fields, field)
		}
		records = append(records, record)

	}

	return records
}

func (f Feature) UserDefinedRecords() (records []models.Record) {
	for _, givenRecord := range f.DefinedRecords {
		record := models.Record{}
		values := strings.Split(givenRecord, ",")
		for i, value := range values {
			f.Fields[i].Value = value
			record.Fields = append(record.Fields, f.Fields[i])
		}
		records = append(records, record)
	}

	return records
}

func (f Feature) Records() []models.Record {
	return append(f.UserDefinedRecords(), f.RandomRecords()...)
}
