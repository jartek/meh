package models

import (
	"log"
	"reflect"
	"regexp"
	"strings"
)

func Attributes(m interface{}) []string {
	typ := reflect.TypeOf(m)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	attrs := make([]string, typ.NumField())

	if typ.Kind() != reflect.Struct {
		log.Fatalln("Not a structure")
	}

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		attrs[i] = p.Name
	}

	return attrs
}

func SnakeCaseAndJoin(str []string) string {
	camelCaseRe := regexp.MustCompile(`[\p{Lu}][\p{Ll}]*`)
	for i := 0; i < len(str); i++ {
		words := camelCaseRe.FindAllString(str[i], -1)
		for i := 0; i < len(words); i++ {
			words[i] = strings.ToLower(words[i])
		}
		str[i] = strings.Join(words, "_")
	}
	return strings.Join(str, ", ")
}

func ModelName(m interface{}) interface{} {
	typ := reflect.TypeOf(m)
	start := strings.Index(typ.String(), ".") + 1
	end := len(typ.String())
	if start == -1 {
		log.Fatalln("Invalid Model Name")
		return -1
	}
	return typ.String()[start:end]
}

func GetAll(m interface{}) interface{} {
	// typ := reflect.TypeOf(m)
	// model_name := ModelName(m)

	attributes := SnakeCaseAndJoin(Attributes(m))
	return attributes
}
