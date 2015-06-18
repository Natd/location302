## Usage ##

```go
import "github.com/natd/location302"
```

## Then ##

```go
location302.GetLink(1, "secret", "http://google.com")
```

or you can build a Location and get a link

```go
loc := location302.Location{}
loc.SetId(1)
loc.SetSecret("secret")
loc.SetUrl("http://google.com")
loc.GetLink()
```

or in one line with fluent api

```go
location302.NewLocation(1, "secret", "http://google.com").GetLink()
```

more fluent api

```go
link := new(location302.Location).WithId(1).WithSecret("secret").WithUrl("http://google.com").GetLink()
```