package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
)

func main() {
	redisClient := initRedisClient()
	redisClient.Set("/urlshort-godoc", "https://godoc.org/github.com/gophercises/urlshort", 0)
	redisClient.Set("/yaml-godoc", "https://godoc.org/gopkg.in/yaml.v2", 0)
	redisClient.Set("/urlshort", "https://github.com/gophercises/urlshort", 0)
	redisClient.Set("/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution", 0)
	redisClient.Set("/google", "https://www.google.com", 0)
	redisClient.Set("/go", "https://golang.org/", 0)

	mux := defaultMux()
	/*pathsToUrls := map[string]string{
			"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
			"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		}
		mapHandler := handleMaps(pathsToUrls, mux)

		yamlString := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution`
		yamlHandler, err := yamlHandler([]byte(yamlString), mapHandler)
		if err != nil {
			panic(err)
		}

		jsonString := `[
				{"Path":"/google", "URL":"https://www.google.com"},
				{"Path":"/go", "URL":"https://golang.org/"}
				]`

		jsonHandler, err := jsonHandler([]byte(jsonString), yamlHandler)

		if err != nil {
			panic(err)
		}
	*/
	redisHandler := handleRedis(redisClient, mux)
	fmt.Println("Starting server on 8080")
	http.ListenAndServe(":8080", redisHandler)
}

func homepage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		fmt.Fprint(res, "Invalid URL")
		return
	}
	fmt.Fprint(res, "Welcome to homepage")
}

func handleMaps(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if newURL, ok := pathToUrls[req.URL.Path]; ok {
			http.Redirect(resp, req, newURL, http.StatusFound)
			return
		}
		fallback.ServeHTTP(resp, req)
	}
}

func handleRedis(client *redis.Client, fallback http.Handler) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if newURL, err := client.Get(req.URL.Path).Result(); err == nil {
			http.Redirect(resp, req, string(newURL), http.StatusFound)
			return
		}
		fallback.ServeHTTP(resp, req)
	}
}

type pathYAML struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func yamlHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := yamlParser(yamlBytes)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathURLs)
	return handleMaps(pathMap, fallback), nil
}

func yamlParser(yamlBytes []byte) ([]pathYAML, error) {
	var pathURLs []pathYAML
	err := yaml.Unmarshal(yamlBytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func jsonHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathURLs []pathYAML
	err := json.Unmarshal(jsonBytes, &pathURLs)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(pathURLs)
	handleMaps(pathMap, fallback)
	return handleMaps(pathMap, fallback), nil
}

func buildMap(pathURLs []pathYAML) map[string]string {
	pathMap := make(map[string]string)
	for _, py := range pathURLs {
		pathMap[py.Path] = py.URL
		println(py.Path + ":" + py.URL)
	}
	return pathMap
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func initRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	//err = client.Set("name", "Elliot", 0).Err()
	return client
}
