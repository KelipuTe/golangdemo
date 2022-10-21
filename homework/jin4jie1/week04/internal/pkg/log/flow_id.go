package log

import (
  "fmt"
  "time"
)

const FLOW_ID_PREFIX string = "week04"

func MakeFlowId() string {
  return fmt.Sprintf("%s-%d", FLOW_ID_PREFIX, time.Now().UnixMilli())
}
