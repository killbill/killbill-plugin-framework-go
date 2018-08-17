package killbill

import (
	pbp "github.com/killbill/killbill-rpc/go/api/plugin/payment"

	"testing"
)

func TestFindPluginProperty(t *testing.T) {
	key := "keyToFind"

	scenarii := []struct {
		Input          []pbp.PluginProperty
		ExpectedResult string
	}{
		{[]pbp.PluginProperty{
		}, ""},
		{[]pbp.PluginProperty{
			{
				Key:   "otherKey",
				Value: "someValue",
			},
		}, ""},
		{[]pbp.PluginProperty{
			{
				Key:   "keyToFind",
				Value: "valueToFind",
			},
			{
				Key:   "otherKey",
				Value: "someValue",
			},
		}, "valueToFind"},
		{[]pbp.PluginProperty{
			{
				Key:   "keyToFind",
				Value: "valueToFind",
			},
			{
				Key:   "keyToFind",
				Value: "byConventionWeReturnTheFirstOne",
			},
		}, "valueToFind"},
	}

	for _, s := range scenarii {
		result := FindPluginProperty(s.Input, key)
		if result != s.ExpectedResult {
			t.Fatalf("invalid. input: %v, expected: %s, got: %s", s.Input, result, s.ExpectedResult)
		}
	}
}
