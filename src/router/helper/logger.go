package helper 

import (
  "fmt"
)

type LogSettings struct {
  DisableSystemProgramLogs bool;
}

type Logger struct {
  DisableSystemProgramLogs bool;
  //LogSettings
}

func (this Logger) Log(logtype ProgramLogType, message string) {
  switch logtype {
  case ProgramLogType(SystemProgramLog): this.logSystem(message);
  case ProgramLogType(ActProgramLog): this.logAction(message);
  case ProgramLogType(ScheduleProgramLog): this.logScheduled(message);
  case ProgramLogType(LlmProgramLog): this.logLLM(message); 
  }
}

func (this Logger) logAction(message string){
    fmt.Println("\U000026A1 "+message);
}

func (this Logger) logSystem(message string){
  if(!this.DisableSystemProgramLogs){
    fmt.Println("\U0001F5A5 "+message);
  }
}

func (this Logger) logLLM(message string){
    fmt.Println("\U0001F9E0 "+message);
}

func (this Logger) logScheduled(message string){
    fmt.Println("\U0001F4C5 "+message);
}

