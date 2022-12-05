package reflectkn

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
)

func interface_test() {
	a := sql.NullString{}
	atype := reflect.TypeOf(a)
	i := (*driver.Valuer)(nil)
	btype := reflect.TypeOf(i).Elem()
	akind := atype.Kind()
	fmt.Println(akind)
	ok := atype.Implements(btype)
	fmt.Println(ok)
}
