package main

import (
	//"fmt"
	"github.com/epxzzy/k-7/src/router/helper"
)

func main(){
	//spinner := helper.ShowSpinner("eyoo", 50, nil);
	lagger := helper.Logger{DisableSystemProgramLogs: false};
	//lagger.Log(helper.SystemProgramLog, "rule")
	//lagger.Log(helper.ActProgramLog, "brawl")
	//lagger.Log(helper.LlmProgramLog, "fool")
	//lagger.Log(helper.ScheduleProgramLog, "loose")
	go helper.InitWsServer()

	for {
		if helper.RecieveQueue.IsEmpty() == false {
			for {
				if helper.RecieveQueue.GetLength() == 0 {
					break
				}

				lagger.Log(helper.LlmProgramLog, helper.RecieveQueue.Dequeue())
			}
		}

		inplut := helper.GetTextInput();
		//fmt.Println("LADIES AND MEN WE HAVE  " + inplut)
		//shows up late for some reason
		//spinner();


		helper.SendToClient(inplut)
		lagger.Log(helper.ActProgramLog, "sending: " +inplut)

	}

}
