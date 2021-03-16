package utils

import "testing"

func TestToJSON(t *testing.T) {
	type args struct {
		v       interface{}
		indents []bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				v: struct {
					A int32 `json:"a"`
				}{
					A: 99,
				},
			},
			want: `{"a":99}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJSON(tt.args.v, tt.args.indents...); got != tt.want {
				t.Errorf("ToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
