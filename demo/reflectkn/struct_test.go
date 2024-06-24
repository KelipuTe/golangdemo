package reflectkn

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type CaseUser struct {
	PublicV1  int
	privateV1 int
}

func (t *CaseUser) GetPublicV1() int {
	return t.PublicV1
}

func (t *CaseUser) SetPublicV1(in int) int {
	t.PublicV1 = in
	return t.PublicV1
}

func (t *CaseUser) setPrivateV1(in int) {
	t.privateV1 = in
}

type CaseUserV2 struct {
	PublicV1 int `tag1:"tag1" tag2:"tag2"`
}

func (t CaseUserV2) GetPublicV1() int {
	return t.PublicV1
}

func TestVisitStructField(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes map[string]any
		wantErr error
	}{
		{
			name:    "非法nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustStructOrPointer,
		},
		{
			name: "非法int指针",
			input: func() *int {
				i := 1
				return &i
			}(),
			wantRes: nil,
			wantErr: ErrMustStructOrPointer,
		},
		{
			name:    "普通结构体",
			input:   CaseUser{PublicV1: 10, privateV1: 20},
			wantRes: map[string]any{"PublicV1": 10, "privateV1": 0},
			wantErr: nil,
		},
		{
			name:    "一级指针",
			input:   &CaseUser{PublicV1: 10, privateV1: 20},
			wantRes: map[string]any{"PublicV1": 10, "privateV1": 0},
			wantErr: nil,
		},
		{
			name: "二级指针",
			input: func() **CaseUser {
				u := &CaseUser{PublicV1: 10, privateV1: 20}
				return &u
			}(),
			wantRes: map[string]any{"PublicV1": 10, "privateV1": 0},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitStructField(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}

func TestVisitStructTag(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes map[string]string
		wantErr error
	}{
		{
			name:    "普通结构体",
			input:   CaseUserV2{PublicV1: 10},
			wantRes: map[string]string{"PublicV1": `tag1:"tag1" tag2:"tag2"`},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitStructTag(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}

func TestEditStructField(t *testing.T) {
	caseList := []struct {
		name      string
		input     any
		wantField string
		wantValue any
		wantErr   error
	}{
		{
			name:      "非法nil",
			input:     nil,
			wantField: "PublicV1",
			wantValue: 11,
			wantErr:   ErrMustStructPointer,
		},
		{
			name:      "非法普通结构体",
			input:     CaseUser{PublicV1: 10, privateV1: 20},
			wantField: "PublicV1",
			wantValue: 11,
			wantErr:   ErrMustStructPointer,
		},
		{
			name: "非法二级指针",
			input: func() **CaseUser {
				u := &CaseUser{PublicV1: 10, privateV1: 20}
				return &u
			}(),
			wantField: "PublicV1",
			wantValue: 11,
			wantErr:   ErrMustStructPointer,
		},
		{
			name:      "一级指针",
			input:     &CaseUser{PublicV1: 10, privateV1: 20},
			wantField: "PublicV1",
			wantValue: 11,
			wantErr:   nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			err := editStructField(v.input, v.wantField, v.wantValue)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
		})
	}
}

func TestVisitStructFunc(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes map[string]*FuncInfo
		wantErr error
	}{
		{
			name:    "非法nil",
			input:   nil,
			wantRes: nil,
			wantErr: ErrMustStructOrPointer,
		},
		{
			name:    "普通结构体",
			input:   CaseUser{},
			wantRes: map[string]*FuncInfo{},
			wantErr: nil,
		},
		{
			name:  "普通结构体V2",
			input: CaseUserV2{},
			wantRes: map[string]*FuncInfo{
				"GetPublicV1": {
					Name:         "GetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(CaseUserV2{})},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: nil,
				},
			},
			wantErr: nil,
		},
		{
			name:  "一级指针",
			input: &CaseUser{PublicV1: 10, privateV1: 20},
			wantRes: map[string]*FuncInfo{
				"GetPublicV1": {
					Name:         "GetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(&CaseUser{})},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: nil,
				},
				"SetPublicV2": {
					Name:         "SetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(&CaseUser{}), reflect.TypeOf(10)},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: nil,
				},
			},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := visitStructFunc(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}

func TestCallStructFunc(t *testing.T) {
	caseList := []struct {
		name    string
		input   any
		wantRes map[string]*FuncInfo
		wantErr error
	}{
		{
			name:  "普通结构体V2",
			input: CaseUserV2{PublicV1: 10},
			wantRes: map[string]*FuncInfo{
				"GetPublicV1": {
					Name:         "GetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(CaseUserV2{})},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: []any{10},
				},
			},
			wantErr: nil,
		},
		{
			name:  "一级指针",
			input: &CaseUser{PublicV1: 10, privateV1: 20},
			wantRes: map[string]*FuncInfo{
				"GetPublicV1": {
					Name:         "GetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(&CaseUser{})},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: []any{10},
				},
				"SetPublicV1": {
					Name:         "SetPublicV1",
					InTypeList:   []reflect.Type{reflect.TypeOf(&CaseUser{}), reflect.TypeOf(10)},
					OutTypeList:  []reflect.Type{reflect.TypeOf(10)},
					OutValueList: []any{0},
				},
			},
			wantErr: nil,
		},
	}

	for _, v := range caseList {
		t.Run(v.name, func(t *testing.T) {
			res, err := callStructFunc(v.input)
			assert.Equal(t, v.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, v.wantRes, res)
		})
	}
}
