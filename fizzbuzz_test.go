package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/fizzbuzz"
)

var _ = Describe("Fizzbuzz", func() {
	Context("knows when a number is", func() {
		It("is divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleByThree(3)).To(BeTrue())
		})
		It("is not divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleByThree(1)).To(BeFalse())
		})
		It("is divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleByFive(5)).To(BeTrue())
		})
		It("is not divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleByFive(4)).To(BeFalse())
		})
	})
})
