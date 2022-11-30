package main 

import "strconv"

func NewData(args []string) (*Data, error) {
  var err error
  var data Data

  data.PhilosCount, err = strconv.Atoi(args[1])
  if err != nil {
    return nil, err
  }

  data.TimeToDie, err = strconv.Atoi(args[2])
  if err != nil {
    return nil, err
  }

  data.TimeToEat, err = strconv.Atoi((args[3]))
  if err != nil {
    return nil, err
  }

  data.TimeToSleep, err = strconv.Atoi(args[4])
  if err != nil {
    return nil, err
  }

  return &data, nil
}
