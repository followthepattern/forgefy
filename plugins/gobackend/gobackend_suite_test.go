package gobackend_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoBackendApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Go Back-end App Test Suite")
}
