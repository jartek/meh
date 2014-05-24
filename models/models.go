package models

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
)

func Attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	attrs := make(map[string]reflect.Type)

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		attrs[p.Name] = p.Type
	}
	return attrs
}

func BuildStruct(m interface{}, values []interface{}) (interface{}, error) {
	typ := reflect.TypeOf(m)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	obj := reflect.New(typ)
	attrs := Attributes(m)
	ctr := 0
	for k, v := range attrs {
		field := obj.Elem().FieldByName(k)
		if v.Kind() == reflect.Int {
			field.SetInt(values[ctr].(int64))
		} else if v.Kind() == reflect.String {
			field.SetString(values[ctr].(string))
		} else if v.Kind() == reflect.Bool {
			field.SetBool(values[ctr].(bool))
		}
		ctr = ctr + 1
	}
	return obj.Interface(), nil
}

func PluralizedModelName(m interface{}) (string, error) {
	typ := reflect.TypeOf(m)
	start := strings.Index(typ.String(), ".") + 1
	end := len(typ.String())
	return inflector.Pluralize(strings.ToLower(typ.String()[start:end])), nil
}

func GetAll(db *sql.DB, m interface{}) interface{} {
	model_name, err := PluralizedModelName(m)
	if err != nil {
		log.Fatalln(err)
	}
	query := "SELECT * FROM " + model_name
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	typ := reflect.TypeOf(m)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	ctr := 0
	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatalln(err)
		}
		for i, _ := range columns {
			b, err := values[i].([]byte)

			if err {
				values[i] = string(b)
			}
		}
		a, err := BuildStruct(m, values)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(a)
		// collection = append(collection, unit)
		ctr = ctr + 1
	}
	return query
}
