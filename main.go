package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Data struct {
  PhilosCount int
  TimeToDie   int
  TimeToEat   int
  TimeToSleep int
}

type Philo struct {
  LastEat     int
  index       int
  isDead      bool
}

func main(){

  argv := os.Args

  // check args validity 
  if len(argv) != 5 && len(argv) != 6 {
    log.Fatal("Invalid Args !")
  }

  data, err := NewData(argv)
  if err != nil {
    log.Fatal("Invalid Data !", err) 
  }


  philos := make([]Philo, data.PhilosCount)
  m := make([]sync.Mutex, data.PhilosCount)
  wg.Add(data.PhilosCount)

  for i := 0; i < data.PhilosCount; i++ {
    go Living(data, &philos[i], m)
  }

  wg.Wait()
  fmt.Println("args : ", argv)

}

func Living(data *Data, philo *Philo, m []sync.Mutex){
  for {
    m[philo.index].Lock()
    fmt.Printf("[time] %d has taken fork\n", philo.index)
    
    m[(philo.index + 1) % data.PhilosCount].Lock()
    fmt.Printf("[time] %d has taken fork\n", philo.index)
    
    fmt.Printf("[time] %d is eating\n", philo.index)
    time.Sleep(time.Duration(data.TimeToEat) * time.Millisecond)

    fmt.Printf("[time] %d is sleeping\n", philo.index)
    time.Sleep(time.Duration(data.TimeToSleep) * time.Millisecond)

    fmt.Printf("[time] %d is thinking\n", philo.index)

    m[philo.index].Lock()
    m[(philo.index + 1) % data.PhilosCount].Lock()
  }
  wg.Done()
}

func timeInMilliseconds() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}

