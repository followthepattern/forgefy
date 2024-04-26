package forgeio_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestForgeIO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ForgeIO Suite")
}
