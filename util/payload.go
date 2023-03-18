package util

import "github.com/gin-gonic/gin"

func ResourcePayload(resource any) map[string]any {
	return gin.H{
		"data": resource,
	}
}

func ErrorsPayload(message string) map[string]any {
	return gin.H{
		"errors": []map[string]any{
			{"detail": message},
		},
	}
}
