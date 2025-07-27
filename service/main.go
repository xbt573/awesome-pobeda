// это ужасный код, лучше не смотрите
package main

import (
	"html/template"
	"net"
	"net/http"
	"os"
	"unicode"

	"github.com/goccy/go-yaml"
	"golang.org/x/net/idna"
)

type Config struct {
	Categories []Category `yaml:"categories"`
}

type Category struct {
	Name  string   `yaml:"name"`
	ID    string   `yaml:"id"`
	Items []Domain `yaml:"items"`
}

type Domain struct {
	Name  string   `yaml:"name"`
	Items []string `yaml:"items"`
}

type Availability struct {
	Resolved   bool
	Reachable  bool
	Successful bool
}

type Template struct {
	Config       Config
	Availability map[string]Availability
}

func isASCII(s string) bool {
	for x := range len(s) {
		if s[x] > unicode.MaxASCII {
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
	defer f.Close()

	var config Config

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	availability := map[string]Availability{}

	for _, category := range config.Categories {
		for _, domain := range category.Items {
			for _, url := range domain.Items {
				local := Availability{}

				ascii := url

				if !isASCII(ascii) { // BRUH
					conv, err := idna.ToASCII(ascii)
					if err != nil {
						panic(err)
					}
					ascii = conv
				}

				_, err := net.LookupIP(ascii)
				if err != nil {
					continue // failed to lookup, Resolved == false
				}

				local.Resolved = true
				availability[url] = local // ugly hack

				resp, err := http.Get("https://" + url)
				if err != nil {
					continue // failed to get, Reachable == false
				}
				resp.Body.Close()

				local.Reachable = true
				availability[url] = local

				if resp.StatusCode == 200 { // Successful == false unless
					local.Successful = true
					availability[url] = local
				}
			}
		}
	}

	input := Template{
		Config:       config,
		Availability: availability,
	}

	tmplfile, err := os.ReadFile("template.md")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("generate").Parse(string(tmplfile))
	if err != nil {
		panic(err)
	}

	if err = tmpl.Execute(os.Stdout, input); err != nil {
		panic(err)
	}
}
