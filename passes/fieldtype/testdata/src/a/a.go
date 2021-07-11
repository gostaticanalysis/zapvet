package a

import "go.uber.org/zap"

func f() {
	zap.String("id", "100")
	zap.Int("id", 100) // want `"id" conflict type Int vs String`
	zap.Any("id", "100") // OK - ignore
	zap.Reflect("id", "100") // OK - ignore
	zap.String("xxx", "100") // OK
}
