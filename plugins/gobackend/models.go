package gobackend

import "github.com/followthepattern/forgefy/specification"

type AppTemplate struct {
	App
	specification.Product
}

type FeatureTemplateModel struct {
	Feature
	App
}
