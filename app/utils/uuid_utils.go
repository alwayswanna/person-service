package utils

import "github.com/google/uuid"

func IsNullableUUID(id *uuid.UUID) bool {
	for x := 0; x < 16; x++ {
		if id[x] != 0 {
			return false
		}
	}
	return true
}
