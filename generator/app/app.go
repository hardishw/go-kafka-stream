package main

import (
  "../generator"
  "../kafka"
  "flag"
  "time"
)

func main()  {
  transferType := flag.String("type", "", "batch | stream")
  numGrps := flag.Int("number-of-groups",0,"How many event groups to generate")
  duration := flag.Int("time",0,"How long in seconds to generate data")
  //batch := flag.Int("batch-size",1000,"How many events per file")
  //interval := flag.Int("interval",1,"How long in seconds to wait inbtween generating files")
  //outputDir := flag.String("output-directory","/data/","Where to write batch files")
  topic := flag.String("topic","/data/","Where to write batch files")
  end := false

  flag.Parse()

  if *transferType == "" {
    panic("type flag must be populated with either 'batch' or 'stream'")
  }

  if *numGrps != 0 && *duration != 0 {
    panic("Only on of these flags can be populated: number-of-groups,time")
  }

  switch *transferType {
  case "stream":
    if *duration == 0 {
      kafka.SendMsgs(*topic,generator.CreateEvents(*numGrps))
    }else {
      time.AfterFunc(time.Second*time.Duration(*duration),func ()  {
        end = true
      })
      for !end {
        kafka.SendMsgs(*topic,generator.CreateEvents(100))
        time.Sleep(time.Second / 100)
      }
    }

  }

}
