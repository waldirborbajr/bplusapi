package error

import "log"

func Catch(err error) {
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
		panic(err)
	}
}
