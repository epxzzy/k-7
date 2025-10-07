package main

import (
	"fmt"

	"github.com/epxzzy/k-7/src/router/helper"
)

func main(){
  spinner := helper.ShowSpinner("eyoo", 50, nil);
  lagger := helper.Logger{DisableSystemProgramLogs: false};
  lagger.Log(helper.SystemProgramLog, "rule")
  lagger.Log(helper.ActProgramLog, "brawl")
  lagger.Log(helper.LlmProgramLog, "fool")
  lagger.Log(helper.ScheduleProgramLog, "loose")
  inplut := helper.GetTextInput();
  fmt.Println("LADIES AND MEN WE HAVE  " + inplut)
	//shows up late for some reason
  spinner();

}
