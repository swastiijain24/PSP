package utils

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

func GenerateSortableTxnID() string {
	t := time.Now().Format("20060102150405")
	id := uuid.New().String()
	
	return fmt.Sprintf("%s-%s", t, id)
}