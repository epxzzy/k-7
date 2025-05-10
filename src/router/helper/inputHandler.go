package helper;


import (
    "fmt"
)
func getTextInput() {
  var promptMessage = "";
  var defaultQuestion = "hey kevin what's the latest three.js update?";

  /*
  var readlineInterface = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });
  */

  defer func ()  {
    if err := recover();err != nil{
      panic("whopti do no inpooo")
    } 
  }();
  if (settings.defaultInputs) {
      fmt.Println("[TEXT] using Default Input...");
      promptMessage += defaultQuestion;
  } else {
      promptActive = true;
      sinkPrompt();
    fmt.Println("INPUT:")

      promptMessage = await new Promise((resolve) => {
        readlineInterface.question("INPUT:\n", (answer) => {
          promptMessage +=
            answer.toLowerCase() === "n" ? defaultQuestion : answer;
          readlineInterface.close();
          resolve(promptMessage);
        });
      });

      freezePromptSinking();
      promptActive = false;
    }
  } catch (error) {
    console.log(error);
  }

  return promptMessage;
}

