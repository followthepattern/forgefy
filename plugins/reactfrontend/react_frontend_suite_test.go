package reactfrontend_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestReactFrontend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "React Front-end App Test Suite")
}
