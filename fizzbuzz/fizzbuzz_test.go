package fizzbuzz_test

import (
	"fizz-buzz/fizzbuzz"
	"reflect"
	"testing"
)

func TestSingleFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{
			name: "Default return n",
			n:    1,
			want: "1",
		},
		{
			name: "Divisible by 3 return Fizz",
			n:    3,
			want: "Fizz",
		},
		{
			name: "Divisible by 5 return buzz",
			n:    5,
			want: "Buzz",
		},
		{
			name: "Divisible by 3 and 5 return FizzBuzz",
			n:    15,
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz.SingleFizzBuzz(tt.n); got != tt.want {
				t.Errorf("SingleFizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeFizzBuzz(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "From 1 to 15",
			args: args{from: 1, to: 15},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name: "From 1 to 1",
			args: args{from: 1, to: 1},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz.RangeFizzBuzz(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeFizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
