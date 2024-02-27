package somepackage

import "reproduction/logger"

var log = logger.NewLogger()

func SomeFunc() {
	log.Warnf("some text %v") // this line is supposed to be reported by govet printf linter
}
