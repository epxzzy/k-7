package main 

import (
  "github.com/epxzzy/k-7/src/router/helper"
)

func main(){
  lagger := helper.Logger{DisableSystemProgramLogs: false};
  lagger.Log(helper.SystemProgramLog, "rule")
  lagger.Log(helper.ActProgramLog, "brawl")
  lagger.Log(helper.LlmProgramLog, "fool")
  lagger.Log(helper.ScheduleProgramLog, "loose")

}
