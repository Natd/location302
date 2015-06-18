## Usage ##

```go
import "github.com/natd/location302"
```

## Then ##

```go
location302.GetLink(1, "secret", "http://google.com")
```

or you can build a Location Struct and get a link

```go
loc := location302.Location{}
loc.SetId(1)
loc.SetSecret("asdasdadww")
loc.SetUrl("http://mail.ru")
loc.GetLink()
```