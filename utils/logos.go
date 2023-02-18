package utils

import "fmt"

var ubuntuLogo string = `
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
               '-'    
`

var alpineLogo string = `
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
`

func GetUbuntuLogo(header, divider, os, kernel, cpu, memory string) string {
	return fmt.Sprintf(ubuntuLogo, header, divider, os, kernel, cpu, memory)
}

func GetAlpineLogo(header, divider, os, kernel, cpu, memory string) string {
	return fmt.Sprintf(alpineLogo, header, divider, os, kernel, cpu, memory)
}

// EOF
