package controllers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBookControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book Controllers Test Suite")
}
