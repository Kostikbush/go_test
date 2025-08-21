package main

import "fmt"

var mes = "Here's my spammy page: https://youth-elixir.com dfjsjs"

func main() {
	resultString := maskUrlInMessage(mes);

	fmt.Println(resultString)
}

func maskUrlInMessage (message string) string {
	byteMessage := []byte(message)
	var startIndex int;

  for i := range byteMessage {
    if  i+4 < len(byteMessage) && byteMessage[i] == 'h' &&
			 byteMessage[i+1] == 't' &&
			 byteMessage[i+2] == 't' &&
			 byteMessage[i+3] == 'p' && 
			 (byteMessage[i+4] == 's' || byteMessage[i+4] == ':'){
				if(byteMessage[i+4] == 's'){
					startIndex = i + 8
				}else {
					startIndex = i+7
				}

			continue
		}

		if(startIndex != 0 && byteMessage[i] != ' ' && startIndex <= i){
			byteMessage[i] = '*' 
			
			continue
		}

		if(byteMessage[i] == ' ') {
			startIndex = 0;

			
			continue
		}
  }

	return string(byteMessage)
}