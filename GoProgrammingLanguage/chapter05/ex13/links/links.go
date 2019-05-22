package links

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func writeFileIfOrigHost(b []byte, rawurl string, origHosts map[string]bool) error {
	parsed, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("url.Parse(%s) error [%v]", rawurl, err)
	}
	if origHosts[parsed.Host] {
		return nil
	}
	path := "./got" + parsed.Path
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("os.MkdirAll(%s) error %v", path, err)
	}

	if err := ioutil.WriteFile(path+"/index.html", b, 0755); err != nil {
		return fmt.Errorf("ioutil.WriteFile error %v", err)
	}
	return nil
}

func Extract(rawurl string, origHosts map[string]bool) ([]string, error) {
	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", rawurl, resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error %v", err)
	}
	doc, err := html.Parse(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", rawurl, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return links, writeFileIfOrigHost(b, rawurl, origHosts)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
