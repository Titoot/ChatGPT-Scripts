package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "net/url"
    "io/ioutil"
    "sync"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <sites_file> <num_threads>")
        return
    }

    sitesFile := os.Args[1]

    // Read the number of threads from the command line argument
    numThreads, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }

    // Read the list of sites from the specified file
    file, err := os.Open(sitesFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create the files for the different status code groups
    f200, err := os.Create("200.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f200.Close()

    f300, err := os.Create("300.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f300.Close()

    f400, err := os.Create("400.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f400.Close()

    // Use a wait group to wait for all goroutines to finish
    var wg sync.WaitGroup
    wg.Add(numThreads)

    // Make a request to each site and print the status code and content length
    scanner := bufio.NewScanner(file)
    for i := 0; i < numThreads; i++ {
        go func() {
            defer wg.Done()

            for scanner.Scan() {
                site := scanner.Text()

                if !strings.Contains(site, "http") {
                    site = "http://" + site
                }
                // Parse the URL and check for any errors
                u, err := url.Parse(site)
                if err != nil {
                    fmt.Println(err)
                    continue
                }

                resp, err := http.Get(u.String())
                if err != nil {
                    fmt.Println(err)
                    continue
                }
                defer resp.Body.Close()

                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    fmt.Println(err)
                    continue
                }

                // Print the status code and content length
                fmt.Printf("%d %d\n", resp.StatusCode, len(body))

                // Write the site to the appropriate file based on the status code group
                if resp.StatusCode >= 200 && resp.StatusCode < 300 {
                    f200.WriteString(fmt.Sprintf("%s [%d]\n", site, len(body)))
                } else if resp.StatusCode >= 300 && resp.StatusCode < 400 {
                    f300.WriteString(fmt.Sprintf("%s [%d]\n", site, len(body)))
                } else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
                    f400.WriteString(fmt.Sprintf("%s [%d]\n", site, len(body)))
                }
            }
        }()
    }

    wg.Wait()

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }
}