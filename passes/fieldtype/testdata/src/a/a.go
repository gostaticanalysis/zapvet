package a

import "go.uber.org/zap"

func f() {
	zap.String("id", "100")
	zap.Int("id", 100) // want `"id" conflict type Int vs String`
	zap.String("xxx", "100") // OK
}
