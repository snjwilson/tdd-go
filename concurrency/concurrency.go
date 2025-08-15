package concurrency

type WebsiteChecker func(string) bool
type Result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan Result)

	for _, url := range urls {
		go func() {
			resultChan <- Result{url, wc(url)}
		}()
	}

	for range len(urls) {
		result := <-resultChan
		results[result.string] = result.bool
	}

	return results
}