package helper 

import (
)

type inputType string;

const (
  text inputType = "TextUserInput"
  audio inputType = "AudioUserInput"
  visual inputType = "VisualUserInput"
)

type LogType string;

type ProgramLogType LogType;
type UserLogType LogType;

const (
  ActProgramLog ProgramLogType = "action"
  LlmProgramLog ProgramLogType = "llm"
  ScheduleProgramLog ProgramLogType = "schedule"
  SystemProgramLog ProgramLogType = "system"
)
var ProgramLogTypeType = map[ProgramLogType]string{
  ActProgramLog: "action",
  LlmProgramLog: "llm",
  ScheduleProgramLog: "schedule",
  SystemProgramLog: "system",

}

const (
  inputUserLog ProgramLogType = "system"
  sudoUserLog ProgramLogType = "system"
)

