```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Add a mutex for synchronization
        
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Lock the mutex before accessing the counter
                        count++
                        mu.Unlock() // Unlock the mutex after accessing the counter
                }()
        }
        
        wg.Wait()
        fmt.Println(count) // this will print 1000
}
```