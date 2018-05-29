package main

import (
        "fmt"
//        "io"
        "io/ioutil"
        "log"
        "net"
        "net/http"
        "net/url"
        "os"
        "strconv"
        "strings"
        "time"
)

func main() {
        if len(os.Args) != 3 {
                fmt.Fprintf(os.Stderr, "Usage: %s N URL\n", os.Args[0])
                fmt.Fprintf(os.Stderr, "N: Number of probes per second\n")
                fmt.Fprintf(os.Stderr, "URL: Where to probe\n")
                os.Exit(1)
        }

        n, _ := strconv.Atoi(os.Args[1])
        delay := int64(float32(1) / float32(n) * 1e9)
        fmt.Fprintf(os.Stderr, "number of probes is %d\n", n)
        fmt.Fprintf(os.Stderr, "delay between probes is %d ns\n", int(delay))

        for true {
            go tcp_probe(os.Args[2])
            time.Sleep(time.Duration(delay))

        }
}

func tcp_probe(ip string) {
    u, _ := url.Parse(ip)

    fmt.Println(u.Host)
    fmt.Println(u.Path)

    tcpAddr, _ := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:80", u.Host))

    conn, _ := net.DialTCP("tcp", nil, tcpAddr)

    start := time.Now()
    _, _ = conn.Write([]byte(fmt.Sprintf("GET %s HTTP/1.0\r\n\r\n", u.Path)))
    result, _ := ioutil.ReadAll(conn)
    stop := time.Now()

    elapsed := stop.Sub(start)

    fmt.Fprintf(os.Stdout, "\n**\nTTLB: %d \nBODY: %s\n", elapsed.Nanoseconds() / 1e6, string(result))
}

func probe(url string) {
        start := time.Now()
        response, err := http.Get(url)
        t := time.Now()
        elapsed := t.Sub(start)

        if err != nil {
                log.Fatal(err)
        } else {
                //defer response.Body.Close()

                responseData, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        log.Fatal(err)
                }
                
                //response.Body.Close()

                responseText := strings.TrimSuffix(string(responseData), "\n")

                fmt.Fprintf(os.Stdout, "STATUS: %s BODY: %s TTLB: %d ms\n", 
                           response.Status, string(responseText), elapsed.Nanoseconds() / 1e6)
        }

}
