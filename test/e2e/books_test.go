package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/igor-karpukhin/gh-labels-test/pkg/books"
)

var _ = Describe("Books test", Label("books"), func() {
	It("should create a new book", func() {
		Expect(books.NewBook("The Hobbit", "J.R.R. Tolkien")).NotTo(BeNil())
	})
})
