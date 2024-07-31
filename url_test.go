package gotenberg

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/v8/test"
)

func TestURL(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req := NewURLRequest("http://example.com")
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestURLComplete(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req := NewURLRequest("http://example.com")
	req.SetBasicAuth("foo", "bar")
	header, err := NewDocumentFromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.Nil(t, err)
	req.Header(header)
	footer, err := NewDocumentFromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	req.Footer(footer)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	req.AddRemoteURLHTTPHeader("A-Header", "Foo")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestURLPageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req := NewURLRequest("http://example.com")
	req.SetBasicAuth("foo", "bar")
	req.PageRanges("1-1")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestURLWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req := NewURLRequest("http://example.com")
	req.SetBasicAuth("foo", "bar")
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestURLScreenshot(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	req := NewURLRequest("https://example.com")
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
