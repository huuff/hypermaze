package util

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetAndCheckEtag(c *gin.Context, data any) bool {
  etagValue := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%+v", data))))

  c.Header("ETag", etagValue)

  if strings.Contains(c.GetHeader("If-None-Match"), etagValue) {
    c.Status(http.StatusNotModified) 
    return true
  } else {
    return false
  }
}

const maxAgeSeconds int = 600
func AddDefaultCacheHeaders(c *gin.Context) {
  c.Header("Cache-Control", fmt.Sprintf("max-age=%d", maxAgeSeconds))
}
