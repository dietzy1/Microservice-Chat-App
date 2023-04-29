package domain

import "time"

func formatTimestamp() string {
	loc, err := time.LoadLocation("Europe/Copenhagen")
	if err != nil {
		panic(err)
	}

	t := time.Now().In(loc)

	return t.Format("02-01-2006 15:04")
}
