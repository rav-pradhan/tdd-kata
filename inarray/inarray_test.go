package inarray

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	t.Run("substrings array containing 'arp', but strings array is empty, returns an empty array", func(t *testing.T) {
		got := InArray([]string{"arp"}, []string{})
		if !reflect.DeepEqual(got, []string{}) {
			t.Errorf("got %v, wanted %v", got, []string{})
		}
	})

	t.Run("substrings array containing 'arp' and 'mat', strings array containing 'harp', 'sharp', 'grasp' returns 'arp'", func(t *testing.T) {
		got := InArray([]string{"arp", "mat"}, []string{"harp", "sharp", "grasp"})
		want := []string{"arp"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("returned substrings array should be in lexicographical order", func(t *testing.T) {
		got := InArray([]string{"mat", "arp", "hash"}, []string{"format", "cry", "harp"})
		want := []string{"arp", "mat"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
