package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"unicode"

	"github.com/goccy/go-yaml"
	"github.com/xbt573/awesome-pobeda/service/generate"
	"golang.org/x/net/idna"
)

type Config struct {
	Categories []generate.Category `yaml:"categories"`
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err = yaml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	availability := map[string]generate.Availability{}

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}

	for _, category := range config.Categories {
		for _, website := range category.Items {
			for _, domain := range website.Items {
				wg.Add(1)
				go func() {
					defer wg.Done()

					status := generate.Availability{}

					defer func() {
						mux.Lock()
						defer mux.Unlock()

						availability[domain.Domain] = status
						fmt.Fprintf(os.Stderr, "%v: %#+v\n", domain.Domain, status)
					}()

					ascii := domain.Domain
					if !isASCII(ascii) {
						conv, err := idna.Lookup.ToASCII(ascii)
						if err != nil {
							panic(err)
						}

						ascii = conv
					}

					_, err := net.LookupIP(ascii)
					if err != nil {
						return
					}
					status.Resolved = true

					resp, err := http.Get("http://" + domain.Domain)
					if err != nil {
						return
					}
					resp.Body.Close()
					status.Reachable = true

					if resp.StatusCode != 200 {
						return
					}
					status.Successful = true
				}()
			}
		}
	}

	wg.Wait()

	fmt.Println(generate.GeneratePage(config.Categories, availability))
}
