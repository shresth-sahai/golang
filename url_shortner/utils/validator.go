
package utils

import (
	"net/url",
	"regexp",
	"strings"
)

func validateURL(urlStr string) (bool,string){
	if strings.TrimSpace(urlStr) == ""{
		return false, "Url cannot be empty"
	}
	// adding schema 

	if !strings.HasPrefix(url,"http://") && !strings.HasPrefix(urlStr,"https://")
	urlSt="https://"+ururlStr

	parsedUrl,err:=url.Parse(urlStr)
	if err!=nil
	return  false,"Invalid URL Format"

	if parsedUrl.Schema!= "http" && parsedUrl.Schema!= "https" 
		return  false,"URL must be HTTP OR HTTPS"

		if parsedUrl.Host == ""
		return false,"Url must have a host"

		return  true,urlStr
}



func validateShortCode(code string) (bool,string)  {
	if len(code)<3 || len(code)>20{
		return false,"short cide must be 3-20 chars "
	}


	validPattern:=regexp.MustCompile('^[a-zA-Z0-9_-]+$')
	if(!validPattern.MatchString(code)){
		return false,"Short Code"
	}
}


