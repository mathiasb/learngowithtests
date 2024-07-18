package concurrency

type WebsiteChecker func(string) bool
type result struct {
	u  string
	ok bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			//results[u] = wc(u)
			resultChannel <- result{u: u, ok: wc(u)}
		}(url)
	}

	// receive and map out the resultchannel into the local results
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.u] = r.ok
	}
	return results
}
