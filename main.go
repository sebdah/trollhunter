package main

import (
	goLogging "github.com/op/go-logging"
	"github.com/sebdah/trollhunter/languages"
	"github.com/sebdah/trollhunter/logging"
	"github.com/sebdah/trollhunter/normalizers"
)

var log goLogging.Logger

func main() {
	// Configure logging
	logging.Setup()

	log.Info("Starting Trollhunter")

	text := "In the above program, we have only defined a method or behavior for Car. Since we then defined Car as an anonymous field in Ferrari, the latter class automatically can call on all the visible behaviors/methods of the anonymous field type. So here, we have not subclassed a parent class, but composed it. But the effect is the very same - you have all the behaviors of the parent with none of the frills of object oriented programming. C’mon, you have to agree with me that that is cool! Let’s bring in the Aston Martin also now, and this time add some individual behavior in addition to that inherited."

	english := languages.NewEnglish()
	normalizer := normalizers.NewSimple()
	_ = normalizer.Normalize(text, english)
}
