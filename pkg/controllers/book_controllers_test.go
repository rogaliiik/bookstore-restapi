package controllers_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
)

// TODO: write unit tests
var _ = Describe("Book Controllers Test:", func() {

	BeforeEach(func() {
		ctl := gomock.NewController(GinkgoT())

		DeferCleanup(func() {
			ctl.Finish()
		})
	})

	Context("Positive test cases", func() {

	})

	Context("Negative test cases", func() {

	})
})
