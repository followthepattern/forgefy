package forgefy_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestForgefy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forgefy Test Suite")
}
