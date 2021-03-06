package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

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
	for _, v := range attrs {
		field := obj.Elem().Field(ctr)
		if v.Kind() == reflect.Int64 {
			if reflect.TypeOf(values[ctr]) == reflect.TypeOf(nil) {
				field.SetInt(-1)
			} else if reflect.TypeOf(values[ctr]) == reflect.TypeOf(time.Now()) {
				field.SetInt(values[ctr].(time.Time).Unix())
			} else if reflect.TypeOf(values[ctr]).Kind() == reflect.Int64 {
				field.SetInt(values[ctr].(int64))
			}
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

func GetAll(db *sql.DB, m interface{}) ([]interface{}, error) {
	model_name, err := PluralizedModelName(m)
	if err != nil {
		log.Fatalln(err)
	}
	query := fmt.Sprintf("SELECT * FROM %s", model_name)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	collection := make([]interface{}, 0)

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
		unit, err := BuildStruct(m, values)
		if err != nil {
			log.Fatalln(err)
		}
		collection = append(collection, unit)
		ctr = ctr + 1
	}
	return collection, nil
}

func GetOne(db *sql.DB, m interface{}, id int) (interface{}, error) {
	model_name, err := PluralizedModelName(m)
	if err != nil {
		log.Fatalln(err)
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", model_name, id)
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

	if rows.Next() {
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
		unit, err := BuildStruct(m, values)
		if err != nil {
			log.Fatalln(err)
		}
		return unit, nil
	}
	return nil, errors.New("RecordNotFound")
}
