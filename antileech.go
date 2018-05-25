package antileech

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gobwas/glob"
	"github.com/penggy/EasyGoLib/utils"

	"github.com/gin-gonic/gin"
)

func AntiLeech() gin.HandlerFunc {
	return func(c *gin.Context) {
		allows := []string{}
		_allows := utils.Conf().Section("").Key("allows").MustString("")
		if _allows != "" {
			allows = strings.Split(_allows, ",")
		}
		if len(allows) == 0 {
			return
		}
		referer := c.Request.Header.Get("Referer")
		if referer == "" {
			// log.Println("referer not found")
			return
		}
		bMatch := false
		if refererUrl, err := url.Parse(referer); err == nil {
			for _, allow := range allows {
				if g, err := glob.Compile(allow); err == nil {
					if g.Match(refererUrl.Hostname()) {
						bMatch = true
						break
					}
				}
			}
		}
		if !bMatch {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
