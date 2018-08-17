package killbill

import (
	pbp "github.com/killbill/killbill-rpc/go/api/plugin/payment"

	"testing"
)

func TestToKbTransactionType(t *testing.T) {
	scenarii := []struct {
		Input          string
		ExpectedResult pbp.PaymentTransactionInfoPlugin_TransactionType
	}{
		{"", -1},
		{"AUTHORIZE", pbp.PaymentTransactionInfoPlugin_AUTHORIZE},
		{"PURCHASE", pbp.PaymentTransactionInfoPlugin_PURCHASE},
	}

	for _, s := range scenarii {
		result := ToKbTransactionType(s.Input)
		if result != s.ExpectedResult {
			t.Fatalf("invalid. input: %s, expected: %v, got: %s", s.Input, result, s.ExpectedResult)
		}
	}
}
