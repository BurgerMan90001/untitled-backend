package util

import "log"

func FatalCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
/* Panics when error exists */
func PanicCheck(err error) {
	if err != nil {
		panic(err)
	}
}