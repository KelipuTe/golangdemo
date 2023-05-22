package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type S6StructWithTag struct {
	Id   int    `orm:"fieldWant=id"`
	Name string `orm:"fieldWant=testName"`
	Age  int8   `orm:"fieldWant=age"`
	Sex  int8   `orm:"fieldWant=sex"`
}

func TestStructToMap(p7tt *testing.T) {
	s5case := []struct {
		name    string
		input   any
		mapWant map[string]any
		errWant error
	}{
		{
			name: "struct",
			input: S6StructWithTag{
				Id:   1,
				Name: "aa",
				Age:  11,
				Sex:  1,
			},
			mapWant: map[string]any{
				"id":       1,
				"testName": "aa",
				"age":      int8(11),
				"sex":      int8(1),
			},
			errWant: nil,
		},
		{
			name: "pointer",
			input: &S6StructWithTag{
				Id:   1,
				Name: "aa",
				Age:  11,
				Sex:  1,
			},
			mapWant: map[string]any{
				"id":       1,
				"testName": "aa",
				"age":      int8(11),
				"sex":      int8(1),
			},
			errWant: nil,
		},
	}

	for _, t4case := range s5case {
		p7tt.Run(t4case.name, func(p7tt *testing.T) {
			m3res, err := F4StructToMap(t4case.input)
			assert.Equal(p7tt, t4case.errWant, err)
			if err != nil {
				return
			}
			assert.Equal(p7tt, t4case.mapWant, m3res)
		})
	}
}
