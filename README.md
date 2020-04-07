# Todo
 - Improve test coverage
 - cache is not distinguishing units and language
 - documentation

# running with docker
```bash
docker run --rm -p 8081:8080 -e TOKEN="<TOKEN VALUE>" weatherllo:latest
```

Then visit following address for 
```bash
http://127.0.0.1:8081
```

# manual compilation 
```bash
go get -u github.com/99designs/gqlgen
make
```


# Example call

```graphql
query {
  default: openweathermap {
    cities(input: ["Warsaw","Bytom"]) {
      city
      general {
        humidity
        pressure
        temparature_feel
      }
    }
  }
  
  polish: openweathermap(input: {
    language: POLISH,
    units: METRIC,
  }) {
    cities(input: ["New York","London"]) {
      city
      description {
        main
        description
      }
      general {
        temparature_feel
      }
      wind {
        speed
      }
    }
  }
  
  empty: openweathermap(input: {
    language: POLISH,
    units: METRIC,
  }) {
    cities(input: []) {
      city
    }
  }
}
```