package one

import (
	"testing"
)

func TestHashStringLen(t *testing.T) {
	type args struct {
		id     string
		maxLen int
	}
	tests := []struct {
		args args
		want string
	}{
		{args{"hi", 0}, "qby7kjpp5hjtbptf7haaxhjp9bbj6eee"},
		{args{"hj", 0}, "qbzqkjpp5hjtbptf7haaxhjp9bbj6eee"},
		{args{"hello world", 0}, "qbw025dsfb5086xppxmb5dh3v6anfbjltaf3v5j2kl9aeeee"},
		{args{"hello world!", 0}, "qbw025dsfb5086xpptt7kjpp5hjtbptf7haaxhjp9bbj6eee"},
		{args{"hi", 100}, "qby7kjpp5hjtbptf7haaxhjp9bbj6eee"},
		{args{"hi", 7}, "qby7kjp"},
		{args{"hi", 1}, "q"},
		{args{"", 0}, "4tr230psac3al4pabhpr38ccs2eeeeee"},
		{args{"a", 2}, "pj"},
		{args{"b", 2}, "pn"},
		{args{"c", 2}, "ps"},
		{args{"d", 2}, "px"},
		{args{"e", 2}, "p1"},
		{args{"f", 2}, "p5"},
		{args{"g", 2}, "p9"},
		{args{"h", 2}, "qd"},
		{args{"i", 2}, "qj"},
		{args{"j", 2}, "qn"},
		{args{"k", 2}, "qs"},
		{args{"l", 2}, "qx"},
		{args{"m", 2}, "q1"},
		{args{"n", 2}, "q5"},
		{args{"o", 2}, "q9"},
		{args{"p", 2}, "rd"},
		{args{"q", 2}, "rj"},
		{args{"r", 2}, "rn"},
		{args{"s", 2}, "rs"},
		{args{"t", 2}, "rx"},
		{args{"u", 2}, "r1"},
		{args{"v", 2}, "r5"},
		{args{"w", 2}, "r9"},
		{args{"x", 2}, "sd"},
		{args{"y", 2}, "sj"},
		{args{"z", 2}, "sn"},
	}
	for _, tt := range tests {
		if got := HashStringLen(tt.args.id, tt.args.maxLen); got != tt.want {
			t.Errorf("HashLen() = %v, want %v", got, tt.want)
		}
	}
}