package repository_testing

import "log"

// logging error
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
