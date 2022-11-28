package errorHandler

import "log"

func Handle(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}
