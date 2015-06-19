## Usage ##

```go
import "github.com/natd/location302"
```

## Then ##

```go
location302.GetLink(1, "secret", "http://google.com")
```

or you can make a Location and get a link

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

you can allocate memory with go new and use fluent api

```go
link := new(location302.Location).WithId(1).WithSecret("secret").WithUrl("http://google.com").GetLink()
```

or with New() method and use fluent api

```go
link := location302.New().WithId(1).WithSecret("secret").WithUrl("http://google.com").GetLink()
```

you can verify link, this method make HEAD request to builded already link and return bool and error

```go
link := location302.New().WithId(1).WithSecret("secret").WithUrl("http://google.com")
response, err := link.Verify()
if err != nil {
	fmt.Println(err) //Handle this error
}else {
	if response {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
```