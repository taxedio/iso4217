package iso4217

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	Nil       *string
	NilStruct *ISOEntry

	EUR       string = "EUR"
	EURStruct        = ISOEntry(isocodes[EUR])

	USD       string = "USD"
	USDStruct        = ISOEntry(isocodes[USD])

	AWG       string = "AWG"
	AWGStruct        = ISOEntry(isocodes[AWG])

	VES       string = "VES"
	VESStruct        = ISOEntry(isocodes[VES])

)

func TestAlphaMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"EUR Test", args{s: "EUR"}, &EUR},
		{"EURO Test", args{s: "euro"}, &EUR},
		{"USD lower Test", args{s: "usd"}, &USD},
		{"Nothing Test", args{s: ""}, Nil},
		{"AWG Test", args{s: "Åruban Florin"}, &AWG},
		{"VES Test", args{s: "Bolívar Soberano"}, &VES},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlphaMatch(tt.args.s)
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alpha3Match() = %#v, want %#v", *got, *tt.want)
			}
		})
	}
}

func TestStructMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ISOEntry
	}{
		{"EUR Test", args{s: EUR}, &EURStruct},
		{"EURO Test", args{s: "euro"}, &EURStruct},
		{"USD upper Test", args{s: "usd"}, &USDStruct},
		{"Nothing Test", args{s: ""}, NilStruct},
		{"AWG Test", args{s: "Åruban Florin"}, &AWGStruct},
		{"VES Test", args{s: "Bolívar Soberano"}, &VESStruct},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
