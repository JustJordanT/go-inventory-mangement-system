package main

import "testing"

func Test_addItem(t *testing.T) {
	type args struct {
		product    Product
		mapProduct map[int]Product
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addItem(tt.args.product, tt.args.mapProduct)
		})
	}
}
