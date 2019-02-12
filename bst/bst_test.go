package bst

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binary Search Tree")
}

var _ = Describe("Binary Search Tree", func() {

	Context("when Binary Search Tree is used", func() {

		It("should correctly insert and return the data with Insert() and CollectData() respectively", func() {
			var t BST
			var expectedPreorder, expectedPostorder, expectedInorder []int

			t.Insert(100, 110, 150, 50, 25, 50, 75)

			expectedPreorder = append(expectedPreorder, 100, 50, 25, 50, 75, 110, 150)
			expectedPostorder = append(expectedPostorder, 50, 25, 75, 50, 150, 110, 100)
			expectedInorder = append(expectedInorder, 25, 50, 50, 75, 100, 110, 150)

			actualPreorder := t.CollectData(Preorder)
			actualPostorder := t.CollectData(Postorder)
			actualInorder := t.CollectData(Inorder)

			Expect(actualPreorder).To(Equal(expectedPreorder))
			Expect(actualPostorder).To(Equal(expectedPostorder))
			Expect(actualInorder).To(Equal(expectedInorder))
		})
	})
})
