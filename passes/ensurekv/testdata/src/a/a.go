package a

import "go.uber.org/zap"

func f() {
	delogger, _ := zap.NewProduction()
	sugared := delogger.Sugar()
	sugared.Errorw("test", "ok", "ok")            // OK
	sugared.Errorw("test")                        // OK
	sugared.Errorw("test", "one")                 // want `Errorw needs to be called with a message, key and value, missing some args`
	sugared.Errorw("test", "one", "two", "three") // want `Errorw needs to be called with a message, key and value, missing some args`
}
