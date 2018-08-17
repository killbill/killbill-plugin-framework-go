package killbill

import "github.com/gofrs/uuid"

func randomUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
