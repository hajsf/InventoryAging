package methods

import "log"

// Identifier should be upper case letter to be exported, regardless it is a type, field, function, ..
func FailOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
		panic(err)
	}
}
