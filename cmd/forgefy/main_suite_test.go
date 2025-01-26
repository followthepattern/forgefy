package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestForgeIOSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forge CLI Suite")
}
