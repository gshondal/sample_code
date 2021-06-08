package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sts"
)

func main(){

  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))

  svc := sts.New(sess)
  result, err := svc.GetCallerIdentity(nil)
    if err != nil {
      fmt.Printf("Error getting caller identity for %s: %v\n", err)
    }
  fmt.Println(result) 
}