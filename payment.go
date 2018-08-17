package killbill

import (
	pbp "github.com/killbill/killbill-rpc/go/api/plugin/payment"
)

func ToKbTransactionType(kbTransactionType string) pbp.PaymentTransactionInfoPlugin_TransactionType {
	kbTransactionTypeInt, ok := pbp.PaymentTransactionInfoPlugin_TransactionType_value[kbTransactionType]
	if !ok {
		return -1
	}
	return pbp.PaymentTransactionInfoPlugin_TransactionType(kbTransactionTypeInt)
}
