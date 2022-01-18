package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultChannel := make(chan result)	// NEW

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}	// NEW
		}(url)
	}

	for i := 0; i < len(urls); i++ {  	// NEW
		r := <-resultChannel			// NEW
		results[r.string] = r.bool		// NEW
	}

	return results
}