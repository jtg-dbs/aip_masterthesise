package error_log

import "log"

// Check --> checks if err is null and logs err with message and log prefix ste to domain
func Check(e error, message string, domain string) {
	log.SetPrefix(domain)
	if e != nil {
		log.Println(message)
		log.Fatal(e)
	}
}

// Log --> logs a message with the log prefix
func Log(prefix string, message string) {
	log.SetFlags(0)
	log.SetPrefix(prefix + ": ")
	log.Print(message)
}
