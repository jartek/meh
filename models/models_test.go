package models_test

import (
	"reflect"
	. "github.com/jartek/worldcup/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Models", func() {
	Describe("#PluralizedModelName", func() {
		It("should return the pluralized structure name", func() {
			plural, err := PluralizedModelName(Team{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(plural).To(Equal("teams"))
		})
	})

	Describe("#Attributes", func() {
		It("Should return a list of attributes", func() {
			type Temp struct {
				attr int
			}
			attrs := make(map[string]reflect.Type)
			attrs["attr"] = reflect.ValueOf(&Temp{}).Elem().Field(0).Type()
			Expect(Attributes(Temp{})).To(Equal(attrs))
		})
	})
})
