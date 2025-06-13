package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var errorStatusMap = map[error]int{
	//errs.InvalidRequest:      http.StatusBadRequest,
	//errs.ErrLocationNotFound: http.StatusNotFound,
	//errs.ExternalApiError:    http.StatusServiceUnavailable,
	//errs.StorageServiceError: http.StatusServiceUnavailable,
	//errs.InvalidID:           http.StatusBadRequest,
}

func HandleError(c *gin.Context, err error) {
	log.Println(err.Error())
	for e, code := range errorStatusMap {
		if errors.Is(err, e) {
			c.JSON(code, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
