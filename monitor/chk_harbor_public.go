package main

//  Harbor 公开仓库检查，null则为正常，目前没有公开的仓库，非null则有
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"os"
)



//var url = "https://harbor.xxx.net/api/projects?project_name=guest"

func main() {
	response, err := http.Get(os.Args[1]+"/api/projects?project_name=guest")
	//response, err := http.Get(url)
	if err != nil {
		fmt.Println("	Your Domain is:", len(os.Args[1]))
		fmt.Println("	Domain error ！ \n" +
			"	sample : go run chk_harbor_public.go https://harbor.***.com")
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

