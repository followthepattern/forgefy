package featureset

import "gopkg.in/yaml.v3"

type FeatureSet struct {
	Version     string `yaml:"version"`
	ProductName string `yaml:"product_name"`
}

func (f FeatureSet) Unmarshal(data []byte) error {
	return yaml.Unmarshal(data, &f)
}

func UnmarshalYaml(data []byte) (fs FeatureSet, err error) {
	err = yaml.Unmarshal(data, &fs)
	return
}
