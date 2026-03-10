package main

import "testing"

func Test_mergeJSONStrings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		jsonStr1 string
		jsonStr2 string
		want     string
		wantErr  bool
	}{
		{
			name: "1",
			jsonStr1: `{
	"status":-1,
	"message":"success",
	"time":1343489
}`,
			jsonStr2: `{
	"status":0,
	"message":"success",
	"url":"www.xxyy.com",
	"data":{}
}`,
			want: `{
	"status":-1,
	"message":"success",
	"time":1343489,
	"url":"www.xxyy.com",
	"data":{}
}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := mergeJSONStrings(tt.jsonStr1, tt.jsonStr2)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("mergeJSONStrings() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("mergeJSONStrings() succeeded unexpectedly")
			}
			if false {
				t.Errorf("mergeJSONStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		map1 map[string]interface{}
		map2 map[string]interface{}
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeJSON(tt.map1, tt.map2)
			// TODO: update the condition below to compare got with tt.want.
			if false {
				t.Errorf("mergeJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
