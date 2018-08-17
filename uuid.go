package killbill

import "github.com/gofrs/uuid"

func RandomUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
