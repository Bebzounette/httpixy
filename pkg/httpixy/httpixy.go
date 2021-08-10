package httpixy

import (
    "net/http"
    "log"
    "fmt"
    "strconv"
    //"strings"
    "github.com/fatih/color"
    "httpixy2/pkg/request"


)

/*func headerInject(urlt string) {
    //timeout := time.Duration(to * 1000000)
    client := request.Client()
    g := color.New(color.FgGreen)
    r := color.New(color.FgRed)

    // baseline request - gauge normal response
    breq, err := http.Get(urlt)
    if err != nil {
        log.Fatal(err)
    }

    // loop through default payloads and inject
    for _, header := range Headers {
        for _, i := range Inject {
            req, err := http.NewRequest("GET", urlt, nil)
            req.Header.Set(header, i)
            resp, err := client.Do(req)
            if err != nil {
                continue
            }
            if breq.ContentLength != resp.ContentLength {
                g.Println("[+] "+"["+urlt+"]"+" "+"["+header+": "+i+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            } else {
                r.Println("[-] "+"["+urlt+"]"+" "+"["+header+": "+i+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            }
            defer resp.Body.Close()
        }
    }
}*/


func Scan(urlt string, method string, /*data string,*/ proxy string) {
	client := request.Client(proxy)
	//timeout := time.Duration(to * 1000000)
    g := color.New(color.FgGreen)
    r := color.New(color.FgRed)

    // baseline request - gauge normal response
    breq, err := http.Get(urlt)
    if err != nil {
    	fmt.Println("error")
        log.Fatal(err)
    }

    // loop through default payloads and inject
    for _, header := range Headers {
        for _, i := range Inject {
            req, err := http.NewRequest(method, urlt, nil/*,strings.NewReader(data)*/)
            req.Header.Set(header, i)
            resp, err := client.Do(req)
            if err != nil {
                continue
            }
            if breq.ContentLength != resp.ContentLength {
                g.Println("[+] "+"["+urlt+"]"+" "+"["+header+": "+i+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            } else {
                r.Println("[-] "+"["+urlt+"]"+" "+"["+header+": "+i+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            }
            defer resp.Body.Close()
        }
    }
}
