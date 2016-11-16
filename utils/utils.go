package utils

func Check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
