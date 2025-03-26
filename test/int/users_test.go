package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/igor-karpukhin/gh-labels-test/pkg/users"
)

var _ = Describe("Users test", Label("test/int/users"), func() {
	It("should create a new user", func() {
		Expect(users.NewUser("test", 10)).NotTo(BeNil())
	})
})
