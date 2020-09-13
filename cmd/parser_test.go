package cmd

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	testCases := []struct {
		desc     string
		filePath string
		want     map[string][]*FieldData
	}{
		{
			desc:     "single",
			filePath: "../test_data/single.go",
			want:     map[string][]*FieldData{"single": {{"fieldName", "string"}}},
		},
		{
			desc:     "multiple fields",
			filePath: "../test_data/multiple_fields.go",
			want: map[string][]*FieldData{
				"multiple": {
					{"field1", "string"},
					{"field2", "int"},
					{"field3", "*[]int"},
					{"field4", "[]int"},
				},
			},
		},
		// {
		// 	desc: "nested first order",
		// 	filePath: "../test_data_first_order_nest.go",
		// 	want: map[string][]*FieldData{
		// 		"root": {
		// 			"strcutField": {type: "struct", children: {"hello": {}}}
		// 		}
		// 	},
		// }
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := parseFile(tC.filePath)
			if !reflect.DeepEqual(v.s, tC.want) {
				t.Fatalf("want %+v got %+v", tC.want, v.s)
			}
		})
	}
}

func TestDir(t *testing.T) {
	testCases := []struct {
		desc    string
		dirPath string
		want    map[string][]*FieldData
	}{
		{
			desc:    "all",
			dirPath: "../test_data",
			want: map[string][]*FieldData{
				"single": {{"fieldName", "string"}},
				"multiple": {
					{"field1", "string"},
					{"field2", "int"},
					{"field3", "*[]int"},
					{"field4", "[]int"},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := parse(tC.dirPath)
			if !reflect.DeepEqual(v.s, tC.want) {
				t.Fatalf("want %+v got %+v", tC.want, v.s)
			}
		})
	}
}
