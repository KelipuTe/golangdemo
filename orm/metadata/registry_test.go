package metadata

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type S6APPUserModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func TestI9RegistryF8Get(p7tt *testing.T) {
	s5s6case := []struct {
		name    string
		input   any
		wantRes *S6Model
		wantErr error
	}{
		{
			// 指针
			name:  "pointer",
			input: &S6APPUserModel{},
			wantRes: &S6Model{
				TableName: "s6_test_model",
				M3FieldToColumn: map[string]*S6ModelField{
					"Id": {
						FieldName:  "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						ColumnName: "id",
					},
					"s6Column": {
						FieldName:  "s6Column",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						ColumnName: "name",
					},
					"Age": {
						FieldName:  "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						ColumnName: "age",
					},
					"Sex": {
						FieldName:  "Sex",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     25,
						ColumnName: "sex",
					},
				},
				M3ColumnToField: map[string]*S6ModelField{
					"id": {
						FieldName:  "Id",
						I9Type:     reflect.TypeOf(int(0)),
						Offset:     0,
						ColumnName: "id",
					},
					"name": {
						FieldName:  "s6Column",
						I9Type:     reflect.TypeOf(""),
						Offset:     8,
						ColumnName: "name",
					},
					"age": {
						FieldName:  "Age",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     24,
						ColumnName: "age",
					},
					"sex": {
						FieldName:  "Sex",
						I9Type:     reflect.TypeOf(int8(0)),
						Offset:     25,
						ColumnName: "sex",
					},
				},
			},
		},
	}

	i9registry := F8NewI9Registry()
	for _, s6case := range s5s6case {
		p7tt.Run(s6case.name, func(p7tt *testing.T) {
			model, err := i9registry.F8Get(s6case.input)
			assert.Equal(p7tt, s6case.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, s6case.wantRes, model)
		})
	}
}
