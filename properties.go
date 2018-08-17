package killbill

import (
	pbp "github.com/killbill/killbill-rpc/go/api/plugin/payment"
)

func FindPluginProperty(properties []pbp.PluginProperty, key string) string {
	for _, prop := range properties {
		if key == prop.Key {
			return prop.GetValue()
		}
	}
	return ""
}
