package productmap_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProjectMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Project Map Suite")
}
