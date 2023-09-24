package util

import "github.com/gin-gonic/gin"

/**
 * Whether this request was made by HTMX, which indicates an HTML fragment
 * is requested, instead of a full page
 */
func IsHxRequest(c *gin.Context) bool {
  return c.GetHeader("HX-Request") != ""
}
