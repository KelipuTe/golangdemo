package typehkn

import "testing"

func Test_f8TypeAssertion(p7s6t *testing.T) {
	value := 1
	f8TypeAssertion(value)
	value2 := "1"
	f8TypeAssertionV2(value2)
}

func Test_f8ForcedConversion(p7s6t *testing.T) {
	value := testStruct{1}
	f8ForcedConversion(value)
	value2 := testStructV2{"1"}
	f8ForcedConversionV2(value2)
}
