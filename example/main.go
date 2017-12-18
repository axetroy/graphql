package main

import (
  "github.com/axetroy/graphql"
  "fmt"
)

func main() {
  client := graphql.New("http://0.0.0.0:9066/api/tian/graphql")
  if res, err := client.Run(graphql.Request{
    OperationName: "posts",
    Query: `
query posts {
  public {
    posts(query: {limit: 10, page: 0}) {
      data {
        id
      }
    }
  }
}
`,
    Variables: graphql.Variables{
      "v": graphql.Variables{
        "adminName": "rootadmin",
        "password":  "rootadmin",
      },
    },
  }); err != nil {
    panic(err)
  } else {
    fmt.Println(res.Data)
    fmt.Println(res.Errors)
  }
}
