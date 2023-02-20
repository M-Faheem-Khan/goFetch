package utils

import (
	"fmt"
	"strings"
)

var logoMap map[string]string = map[string]string{
	"tux": `
              a8888b.
             d888888b.
             8P"YP"Y88
             8|o||o|88
             8'    .88
             8'._.' Y8.         %s
            d/      '8b.        %s
           dP   .    Y8b.       %s
          d8:'  "  '::88b       %s
         d8"         'Y88b      %s
        :8P    '      :888      %s
         8a.   :     _a88P      %s
       ._/"Yaa_:   .| 88P|
       \    YP"    '| 8P  '.
       /     \.___.d|    .'
       '--..__)8888P'._.'

  `, // http://www.ascii-art.de/ascii/jkl/linux.txt - JGS, a:f
	"ubuntu": `
               .-.    
         .-'''(|||)   
      ,'\ \    '-'.    %s
     /   \ '''-.   '   %s
   .-.  ,       '___:  %s
  (:::) :        ___   %s
   '-'  '       ,   :  %s
     \   / ,..-'   ,   %s
      './ /    .-.'   
         '-..-(   )   
               '-'`,
	"alpine": `
        :::::::::::::::       
       :::::::::::::::::     
     ::::::::::::::::::::     %s
   ,:::::::: *::::::::::::    %s
  ::::::::      #    ::::::#  %s
 ::::::    :::(   #*   #::::: %s
 *::#   .. #:::::    :   :::  %s
   ::::::::::::::::::::::::   %s
    (:::::::::::::::::::::   
      :::::::::::::::::::     
       :::::::::::::::::      
`,
	"debian": `

        .:^^^^^^::..                  
        .:^^^:.....::^^^:
      :~^:.          .^~^.    
      .^^.    ....     ^^.    %s
      ^^     ..        .~.    %s
      ~.    :.         .~.    %s
      ^.     ::        .^:    %s
      ^:      ::.   ..::.     %s
      :^:      .......        %s
      :~:                                
      .^:                               
        :::.                            
          ..:.                                    
  `,
}

func getLinuxDistribution(os string) string {
	distribution := strings.ToLower(strings.Split(os, " ")[1])
	logo, err := logoMap[distribution]
	if !err {
		logo = logoMap["tux"]
		return logo
	}
	return logo
}

func GenerateOutput(header, divider, os, kernel, cpu, memory string) string {
	logo := getLinuxDistribution(os)
	return fmt.Sprintf(logo, header, divider, os, kernel, cpu, memory)
}

// EOF
