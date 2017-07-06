package cmd

// import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"os"
//	"testing"

//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/suite"
// )

// const outputPath = "../fixtures/fomatted_output.json"

// type DownloadTestSuite struct {
//	suite.Suite
// }

// func (suite *DownloadTestSuite) TearDownTest() {
//	os.RemoveAll("../fixtures/output")
//	os.RemoveAll("../fixtures/download")
// }

// func (suite *DownloadTestSuite) TestDownloadWithFileNames() {
//	defer os.Remove("../fixtures/project/assets/hello.txt")
//	client, server := newClientAndTestServer(func(w http.ResponseWriter, r *http.Request) {
//		if "asset[key]=assets/hello.txt" == r.URL.RawQuery {
//			fmt.Fprintf(w, jsonFixture("responses/asset"))
//		} else {
//			w.WriteHeader(404)
//			fmt.Fprintf(w, "404")
//		}
//	})
//	defer server.Close()

//	err := download(client, []string{"assets/hello.txt"})
//	assert.Nil(suite.T(), err)
// }

// func (suite *DownloadTestSuite) TestDownloadWithReadOnly() {
//	defer os.Remove("../fixtures/project/assets/hello.txt")
//	client, server := newClientAndTestServer(func(w http.ResponseWriter, r *http.Request) {
//		if "asset[key]=assets/hello.txt" == r.URL.RawQuery {
//			fmt.Fprintf(w, jsonFixture("responses/asset"))
//		} else {
//			w.WriteHeader(404)
//			fmt.Fprintf(w, "404")
//		}
//	})
//	defer server.Close()
//	client.Config.ReadOnly = true
//	err := download(client, []string{"output/nope.txt"})
//	assert.Nil(suite.T(), err)
// }

// func (suite *DownloadTestSuite) TestDownloadAll() {
//	client, server := newClientAndTestServer(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, jsonFixture("responses/asset"))
//	})
//	defer server.Close()

//	client.Config.Directory = "../fixtures/download"
//	os.MkdirAll(client.Config.Directory, 7777)
//	defer os.Remove(client.Config.Directory)

//	assert.Nil(suite.T(), download(client, []string{}))
// }

// func (suite *DownloadTestSuite) TestExpandWildcards() {
//	requestCount := make(chan int, 100)
//	client, server := newClientAndTestServer(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, jsonFixture("responses/assets"))
//		requestCount <- 1
//	})
//	defer server.Close()

//	filenames, err := expandWildcards(client, []string{"assets/hello.txt"})
//	assert.Nil(suite.T(), err)
//	assert.Equal(suite.T(), len(requestCount), 0)
//	assert.Equal(suite.T(), filenames, []string{"assets/hello.txt"})

//	filenames, err = expandWildcards(client, []string{"assets/*"})
//	assert.Nil(suite.T(), err)
//	assert.Equal(suite.T(), len(requestCount), 1)
//	assert.Equal(suite.T(), filenames, []string{"assets/goodbye.txt", "assets/hello.txt"})
// }

// func TestDownloadTestSuite(t *testing.T) {
//	suite.Run(t, new(DownloadTestSuite))
// }

// func fileFixture(name string) *os.File {
//	path := fmt.Sprintf("../fixtures/%s.json", name)
//	file, _ := os.Open(path)
//	return file
// }

// func jsonFixture(name string) string {
//	bytes, err := ioutil.ReadAll(fileFixture(name))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return string(bytes)
// }
