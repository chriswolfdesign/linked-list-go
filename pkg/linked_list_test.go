package pkg_test

import (
	. "linked-list/pkg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	var (
		list *LinkedList
	)
	BeforeEach(func() {
		list = &LinkedList{}
	})

	Describe("Add tests", func() {
		It("Adds item to empty list", func() {
			Expect(list.Size).To(Equal(0))
			list.Add(1)
			Expect(list.Size).To(Equal(1))
		})

		It("Adds item to list with one element", func() {
			list.Add(1)
			Expect(list.Size).To(Equal(1))
			list.Add(2)
			Expect(list.Size).To(Equal(2))
		})

		It("Adds item to list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			Expect(list.Size).To(Equal(2))
			list.Add(3)
			Expect(list.Size).To(Equal(3))
		})
	})

	Describe("Get tests", func() {
		It("Returns error in empty list", func() {
			item, err := list.Get(0)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Returns only item in a single item list", func() {
			list.Add(1)
			item, err := list.Get(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
		})

		It("Returns first item in list with multiple items", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Get(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
		})

		It("Returns last item in list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Get(2)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(3))
		})

		It("Returns middle item in list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Get(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(2))
		})

		It("Returns error when given negative index", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Get(-2)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Returns error when index is same as list size", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Get(3)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})
	})

	Describe("AddFirst tests", func() {
		It("Adds to beginning of empty list", func() {
			Expect(list.Size).To(Equal(0))
			list.AddFirst(1)
			Expect(list.Size).To(Equal(1))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
		})

		It("Adds to beginning of list with one element", func() {
			list.Add(2)
			Expect(list.Size).To(Equal(1))
			list.AddFirst(1)
			Expect(list.Size).To(Equal(2))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})

		It("Adds to beginning of list with multiple elements", func() {
			list.Add(2)
			list.Add(3)
			Expect(list.Size).To(Equal(2))
			list.AddFirst(1)
			Expect(list.Size).To(Equal(3))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})
	})

	Describe("Peek tests", func() {
		It("Errors when peeking at empty list", func() {
			item, err := list.Peek()
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Gets only item in a list with one element", func() {
			list.Add(1)
			item, err := list.Peek()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
		})

		It("Gets first items in list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.Peek()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
		})
	})

	Describe("RemoveFirst tests", func() {
		It("Errors when removing the first item from an empty list", func() {
			item, err := list.RemoveFirst()
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Removes and returns only item from list with one element", func() {
			list.Add(1)
			item, err := list.RemoveFirst()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
			Expect(list.Size).To(Equal(0))
		})

		It("Removes and returns the first element from list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.RemoveFirst()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
			item, _ = list.Get(0)
			Expect(item).To(Equal(2))
			item, _ = list.Get(1)
			Expect(item).To(Equal(3))
		})
	})

	Describe("RemoveLast tests", func() {
		It("Errors when removing last item from an empty list", func() {
			item, err := list.RemoveLast()
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Removes only item from list with one element", func() {
			list.Add(1)
			item, err := list.RemoveLast()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
			Expect(list.Size).To(Equal(0))
		})

		It("Removes last item from list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.RemoveLast()
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(3))
			Expect(list.Size).To(Equal(2))
			item, _ = list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})
	})

	Describe("AddAtIndex tests", func() {
		It("Adding to empty list adds to the list", func() {
			err := list.AddAtIndex(1, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(1))
			Expect(list.Peek()).To(Equal(1))
		})

		It("Adds to beginning of list with single element", func() {
			list.Add(2)
			err := list.AddAtIndex(1, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(2))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})

		It("Adds to end of list with single element", func() {
			list.Add(1)
			err := list.AddAtIndex(2, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(2))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})

		It("Adds to beginning of list with multiple elements", func() {
			list.Add(2)
			list.Add(3)
			err := list.AddAtIndex(1, 0)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(3))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
			item, _ = list.Get(2)
			Expect(item).To(Equal(3))
		})

		It("Adds to the middle of list with multiple elements", func() {
			list.Add(1)
			list.Add(3)
			err := list.AddAtIndex(2, 1)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(3))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
			item, _ = list.Get(2)
			Expect(item).To(Equal(3))
		})

		It("Adds to the end of list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			err := list.AddAtIndex(3, 2)
			Expect(err).ToNot(HaveOccurred())
			Expect(list.Size).To(Equal(3))
			item, _ := list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
			item, _ = list.Get(2)
			Expect(item).To(Equal(3))
		})

		It("Throws error when adding to negative index", func() {
			err := list.AddAtIndex(1, -1)
			Expect(err).To(HaveOccurred())
		})

		It("Throws error when adding to index greater than size", func() {
			err := list.AddAtIndex(1, 1)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("RemoveAtIndex tests", func() {
		It("Errors when attempting to remove from empty list", func() {
			item, err := list.RemoveAtIndex(0)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Removes only item from list with single element", func() {
			list.Add(1)
			item, err := list.RemoveAtIndex(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
			Expect(list.Size).To(Equal(0))
		})

		It("Removes first item from list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.RemoveAtIndex(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(1))
			item, _ = list.Get(0)
			Expect(item).To(Equal(2))
			item, _ = list.Get(1)
			Expect(item).To(Equal(3))
		})

		It("Removes item from middle of list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.RemoveAtIndex(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(2))
			item, _ = list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(3))
		})

		It("Remove item from end of list with multiple elements", func() {
			list.Add(1)
			list.Add(2)
			list.Add(3)
			item, err := list.RemoveAtIndex(2)
			Expect(err).ToNot(HaveOccurred())
			Expect(item).To(Equal(3))
			item, _ = list.Get(0)
			Expect(item).To(Equal(1))
			item, _ = list.Get(1)
			Expect(item).To(Equal(2))
		})

		It("Throws error when removing from negative index", func() {
			list.Add(1)
			item, err := list.RemoveAtIndex(-1)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Throws error when removing from index equal to size", func() {
			list.Add(1)
			item, err := list.RemoveAtIndex(1)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})

		It("Throws error when removing from index greather than size", func() {
			list.Add(1)
			item, err := list.RemoveAtIndex(2)
			Expect(err).To(HaveOccurred())
			Expect(item).To(BeNil())
		})
	})
})
