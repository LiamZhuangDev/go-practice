package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Page struct {
	URL  string
	Body []byte
	Err  error
}

type WebCrawler struct {
	maxConcurrency int
	timeout        time.Duration
}

func NewWebCrawler(maxConcurrency int, timeout time.Duration) *WebCrawler {
	return &WebCrawler{
		maxConcurrency: maxConcurrency,
		timeout:        timeout,
	}
}

func (wc *WebCrawler) fetchURL(url string) (*Page, error) {
	client := &http.Client{
		Timeout: wc.timeout,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body := make([]byte, 4096)
	n, err := res.Body.Read(body)
	if err != nil {
		return nil, err
	}

	return &Page{
		URL:  url,
		Body: body[:n],
	}, nil
}

// ┌───────────────┐     ┌─────────────────────┐     ┌─────────────────────────┐     ┌───────────────────┐      ┌─────────────────────┐
// │ urls slice    │     │ jobs channel        │     │ worker goroutines       │     │ pagesChan channel │      │ main goroutine      │
// │ ["u1","u2"]   │ ──▶ │ chan string         │ ──▶ │ fetchURL + build Page   │ ──▶ │ chan *Page        │ ──▶  │ range pagesChan     │
// │               │     │ (buffered)          │     │ (N workers)             │     │ (buffered)        │      │ collect results     │
// └───────────────┘     └─────────────────────┘     └─────────────────────────┘     └───────────────────┘      └─────────────────────┘
func (wc *WebCrawler) fetchURLs(urls []string) []*Page {
	jobs := make(chan string, len(urls))
	pagesChan := make(chan *Page, len(urls))
	var wg sync.WaitGroup

	// Start workers
	for range wc.maxConcurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				page, err := wc.fetchURL(url)
				if err != nil {
					page = &Page{
						URL: url,
						Err: err,
					}
				}
				pagesChan <- page
			}
		}()
	}

	// Submit jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Wait all workers complete
	// Why need to put wg.Wait() and close() inside a closer goroutine? Important!
	// If pagesChan fills up, workers block, wg.Done() is never reached, so wg.Wait() blocks forever
	// This is classic producer-consumer deadlock.
	//
	// You might think the buffer is big enough, so sends won't block. This is NOT guaranteed.
	// Correctness should not depends on buffer size.
	go func() {
		wg.Wait()
		close(pagesChan)
	}()

	// Return pages
	var pages []*Page
	for p := range pagesChan {
		pages = append(pages, p)
	}

	return pages
}

func WebCrawlerTest() {
	crawler := NewWebCrawler(3, 5*time.Second)

	urls := []string{
		"https://go.cyub.vip/gmp/gmp-model/https://go.cyub.vip/gmp/gmp-model/https://go.cyub.vip/gmp/gmp-model/",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://cloud.google.com/application/web3/faucet/ethereum/sepolia",
	}

	start := time.Now()
	pages := crawler.fetchURLs(urls)
	duration := time.Since(start)

	fmt.Printf("Fetch took %v\n", duration)
	for _, p := range pages {
		if p.Err != nil {
			fmt.Printf("Fetch %s failed, error: %s\n", p.URL, p.Err)
		} else {
			fmt.Printf("Fetch %s succeeded, content: %v\n", p.URL, p.Body)
		}
	}
}
