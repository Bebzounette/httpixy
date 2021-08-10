package main 

import (
    //"net"
    //"net/http"
    //"net/url"
    //"crypto/tls"
    //"fmt"
    //"flag"
    //"log"
    //"os"
    //"bufio"
    //"time"
    //"strconv"
    //"github.com/fatih/color"
    "httpixy2/internal/runner"
    //"httpixy2/pkg/httpixy"
    //"httpixy2/pkg/request"


)

/*func payloadInject() {
    timeout := time.Duration(to * 1000000)
    g := color.New(color.FgGreen)
    r := color.New(color.FgRed)

    // open and iterate
    file, err := os.Open(pfile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    // baseline request - gauge normal response
    breq, err := http.Get(urlt)
    if err != nil {
        log.Fatal(err)
    }

    // loop through payload file and inject
    for scanner.Scan() {
        for _, header := range headers {
            req, err := http.NewRequest("GET", urlt, nil)
            req.Header.Set(header, scanner.Text())
            resp, err := client.Do(req)
            if err != nil {
                continue
            }
            if breq.ContentLength != resp.ContentLength {
                g.Println("[+] "+"["+urlt+"]"+" "+"["+header+": "+scanner.Text()+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            } else {
                r.Println("[-] "+"["+urlt+"]"+" "+"["+header+": "+scanner.Text()+"]"+" "+"[Code: "+strconv.Itoa(int(resp.StatusCode))+"] "+"[Size: "+ strconv.Itoa(int(resp.ContentLength))+"]")
            }
            defer resp.Body.Close()
        }
    }
}
*/


func main() {
    options := runner.ParseOptions()
    //fmt.Println(options.Concurrency)
    runner.New(options)
    /*fmt.Println(options)
    url := options.URL
    fmt.Println(url)
    fmt.Println(httpixy.Inject) 
    /*httpixy.headerInject(options.URL)*/
}
