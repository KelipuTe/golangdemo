package orm_metadata

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type S6TestModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func TestRegistry_Get(p7tt *testing.T) {
	s5s6case := []struct {
		name    string
		input   any
		resWant *S6OrmModel
		errWant error
	}{
		{
			// 指针
			name:  "pointer",
			input: &S6TestModel{},
			resWant: &S6OrmModel{
				TableName: "test_model",
				M3StructToField: map[string]*S6ModelField{
					"Id": {
						StructName: "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						FieldName:  "id",
					},
					"FirstName": {
						StructName: "FirstName",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						FieldName:  "first_name",
					},
					"Age": {
						StructName: "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						FieldName:  "age",
					},
					"LastName": {
						StructName: "LastName",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     32,
						FieldName:  "last_name",
					},
				},
				M3FieldToStruct: map[string]*S6ModelField{
					"id": {
						StructName: "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						FieldName:  "id",
					},
					"first_name": {
						StructName: "FirstName",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						FieldName:  "first_name",
					},
					"age": {
						StructName: "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						FieldName:  "age",
					},
					"last_name": {
						StructName: "LastName",
						I9Type:     reflect.TypeOf(int8(0)),
						FieldName:  "last_name",
						Offset:     32,
					},
				},
			},
		},
	}

	i9registry := NewI9Registry()
	for _, s6case := range s5s6case {
		p7tt.Run(s6case.name, func(p7tt *testing.T) {
			model, err := i9registry.F4Get(s6case.input)
			assert.Equal(p7tt, s6case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, s6case.resWant, model)
		})
	}
}
