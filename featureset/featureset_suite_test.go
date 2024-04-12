package featureset_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFeatureSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feature Set Suite")
}
