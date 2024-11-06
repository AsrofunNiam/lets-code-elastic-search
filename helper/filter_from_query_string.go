package helper

import "github.com/gin-gonic/gin"

func FilterFromQueryString(c *gin.Context, allowFilters ...string) map[string]string {
	filters := map[string]string{}
	for _, allowFilter := range allowFilters {
		filter := c.Query(allowFilter)
		if filter != "" {
			filters[allowFilter] = filter
		}
	}
	return filters
}
