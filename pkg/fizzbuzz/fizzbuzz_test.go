package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nikimanoledaki/fizzbuzz/pkg/fizzbuzz"
)

var _ = Describe("Fizzbuzz", func() {
	Context("knows when a number is", func() {
		It("is divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(3, 3)).To(BeTrue())
		})
		It("is not divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 3)).To(BeFalse())
		})
		It("is divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(5, 5)).To(BeTrue())
		})
		It("is not divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(4, 5)).To(BeFalse())
		})
		It("is divisible by 15", func() {
			Expect(fizzbuzz.IsDivisibleBy(15, 15)).To(BeTrue())
		})
		It("is not divisible by 15", func() {
			Expect(fizzbuzz.IsDivisibleBy(14, 15)).To(BeFalse())
		})
	})

	Context("when playing the game, Fizzbuzz says...", func() {
		It("'fizz' when a number is divisible by 3", func() {
			Expect(fizzbuzz.Says(3)).To(Equal("fizz"))
		})
		It("'buzz' when a number is divisible by 5", func() {
			Expect(fizzbuzz.Says(5)).To(Equal("buzz"))
		})
		It("'fizzbuzz' when a number is divisible by 15", func() {
			Expect(fizzbuzz.Says(15)).To(Equal("fizzbuzz"))
		})
		It("returns number that is not a multiple of 3, 5, or 15", func() {
			Expect(fizzbuzz.Says(14)).To(Equal("14"))
		})
	})
})
