## Simple GraphQL Client for Go

![License](https://img.shields.io/badge/license-Apache-green.svg)


## Usage

```bash
go get -v github.com/axetroy/graphql
```


```go
package main

import (
  "github.com/axetroy/graphql"
  "fmt"
)

func main() {
  client := graphql.New("http://0.0.0.0:9066/api/graphql")
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
        "key": "value",
      },
    },
  }); err != nil {
    panic(err)
  } else {
    fmt.Println(res.Data)
    fmt.Println(res.Errors)
  }
}

```

## Contributing

[Contributing Guid](https://github.com/axetroy/graphql/blob/master/CONTRIBUTING.md)

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars1.githubusercontent.com/u/9758711?v=3" width="100px;"/><br /><sub>Axetroy</sub>](http://axetroy.github.io)<br />[💻](https://github.com/axetroy/graphql/commits?author=axetroy) [🐛](https://github.com/axetroy/graphql/issues?q=author%3Aaxetroy) 🎨 |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fgraphql.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fgraphql?ref=badge_large)
