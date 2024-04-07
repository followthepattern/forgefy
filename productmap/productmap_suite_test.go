package productmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product Map Suite")
}
