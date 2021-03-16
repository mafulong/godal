package utils

import "testing"

func TestToCamelFirstLower(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				str: "AaBbCc",
			},
			want: "aaBbCc",
		},
		{
			args: args{
				str: "aa_bb_cc",
			},
			want: "aaBbCc",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelFirstLower(tt.args.str); got != tt.want {
				t.Errorf("ToCamelFirstLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCamelFirstUpper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				str: "aa_bb_cc",
			},
			want: "AaBbCc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelFirstUpper(tt.args.str); got != tt.want {
				t.Errorf("ToCamelFirstUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				str: "AaBbCc",
			},
			want: "aa_bb_cc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnakeCase(tt.args.str); got != tt.want {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
