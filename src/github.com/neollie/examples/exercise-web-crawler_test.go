package examples

import (
	"testing"
	"os"
	"io/ioutil"
	"fmt"
)

type  testData struct {
	site string
	destDir string // filepath to directory
	filenames []string
}

// TODO
func TestCrawl(t *testing.T) {
	data := testData{"test1/page1.html", "tmp/page1", []string{"page1.html","test.html"}}
	os.Remove(data.destDir)
	os.Mkdir(data.destDir,0777)
	os.Remove(data.destDir)
	Crawl(data.site, data.destDir)

	// Check existence of the directory

	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	fmt.Print("test")
}