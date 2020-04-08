# TLDR:
Application weatherllo provides HTTP endpoint with GraphQl and timed caching.
It fetches https://openweathermap.org using API TOKEN.


# running with docker
```bash
docker run --rm -p 8081:8080 -e TOKEN="<TOKEN VALUE>" phob0spl/weatherllo:latest
```

Then visit following address for web interface :)
```bash
http://127.0.0.1:8081
```



# manual compilation 
```bash
go get -u github.com/99designs/gqlgen
make
```

# Todo
 - Improve test coverage
 - cache is not distinguishing units and language
 - documentation
 - dockerfile improvements: security hardening: run as user, use smaller base image, keeping in mind ca-certificates must be installed
 - discuss if error only should be returned when one of cities provided fails or data + error
 - provide metrics with API calls counter


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