package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/igor-karpukhin/gh-labels-test/pkg/vehicles"
)

var _ = Describe("Users test", Label("test-vehicles"), func() {
	It("should create a new user", func() {
		Expect(vehicles.NewCar("BMW", "525")).NotTo(BeNil())
	})
})
