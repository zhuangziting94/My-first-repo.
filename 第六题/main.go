package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const M = 2000
const N = 10
const K = 200

var ch = make(chan struct{})
var wait sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	sum := 0
	for n := 0; ; n++ {
		for i := 0; i < N; i++ {
			r := rand.Intn(K + 1)
			wait.Add(1)
			go func(t int) {
				ch <- struct{}{}
				fmt.Printf("the Car--%d Moved %d\n", t, r)
				sum += r
				wait.Done()
			}(10*n + i)
			<-ch
			wait.Wait()
		}

		if sum >= M {
			break
		}
	}
}
/*输出：（本人不大会实时输出）
the Car--0 Moved 147
the Car--1 Moved 142
the Car--2 Moved 155
the Car--3 Moved 148
the Car--4 Moved 125
the Car--5 Moved 48
the Car--6 Moved 10
the Car--7 Moved 197
the Car--8 Moved 44
the Car--9 Moved 124
the Car--10 Moved 43
the Car--11 Moved 84
the Car--12 Moved 21
the Car--13 Moved 113
the Car--14 Moved 178
the Car--15 Moved 137
the Car--16 Moved 48
the Car--17 Moved 77
the Car--18 Moved 66
the Car--19 Moved 55
the Car--20 Moved 45
the Car--21 Moved 25
the Car--22 Moved 113
the Car--23 Moved 146
the Car--24 Moved 170
the Car--25 Moved 115
the Car--26 Moved 33
the Car--27 Moved 177
the Car--28 Moved 66
the Car--29 Moved 9
  */
