package compress

import (
	"compress/gzip"
	"net/http"
	"strings"

	vo "github.com/alexbirbirdev/2026_ya.practicum_sprint-1_metrics/internal/domain/value_object"
)

type compressWriter struct {
	w           http.ResponseWriter
	zw          *gzip.Writer
	writeHeader bool
	compress    bool
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}
func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(p []byte) (int, error) {
	if !c.writeHeader {
		c.WriteHeader(http.StatusOK)
	}
	if c.compress {
		return c.zw.Write(p)
	}
	return c.w.Write(p)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	if c.writeHeader {
		return
	}
	c.writeHeader = true

	ct := c.w.Header().Get("Content-Type")
	c.compress = statusCode < 300 &&
		(strings.Contains(ct, vo.ContentTypeApplicationJSON) || strings.Contains(ct, vo.ContentTypeTextHTML))
	if c.compress {
		c.w.Header().Set("Content-Encoding", "gzip")
		c.w.Header().Del("Content-Length")
	}

	c.w.WriteHeader(statusCode)
}

func (c *compressWriter) Close() error {
	if !c.compress {
		return nil
	}
	return c.zw.Close()
}
