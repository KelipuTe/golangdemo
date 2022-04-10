package response

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func Set200Response(p1gc *gin.Context, mapdata map[string]interface{}) {
  p1gc.JSON(
    http.StatusOK,
    gin.H{
      "res_code": http.StatusOK,
      "res_msg":  "",
      "res_data": mapdata,
    })
}

func Set500Response(p1gc *gin.Context, msg string) {
  p1gc.JSON(
    http.StatusBadRequest,
    gin.H{
      "res_code": http.StatusInternalServerError,
      "res_msg":  msg,
      "res_data": map[string]interface{}{},
    })
}
