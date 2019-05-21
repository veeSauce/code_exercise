package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"time"
	"strings"
)

type nameInfo struct{
	FirstName string `json:"name"`
	Surname string `json:"surname"`
	Gender string `json:"gender"`
	Region string `json:"region"`
}

type joke struct{
	Type string	`json:"type"`
	Value struct{
		Id int	`json:"id"'`
		Joke string `json:"joke"`
		categories []string `json:"categories"`
}
}

func main(){

		fmt.Println("Starting the application...")

		r := mux.NewRouter()
		r.HandleFunc("/", HomeHandler).Methods("GET")
		r.HandleFunc("/health", Health).Methods("GET")
		http.Handle("/", r)
		log.Println("running on port 5000")

		log.Fatal(http.ListenAndServe(":5000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request){

	nameChannel := make(chan *nameInfo)
	jokeChannel := make(chan string)

	log.Println("Handler requested")
	retry := 3
	name := nameInfo{}
	joke := joke{}

	go name.GetName(retry, nameChannel)

	go joke.GetJoke(retry, nameChannel, w, jokeChannel)


	w.Write([]byte( <- jokeChannel))


}


func (n nameInfo)GetName(retry int, nameChannel chan *nameInfo) {

	log.Println("getName called")

			response, err := http.Get("http://uinames.com/api/")
			if err != nil && retry > 0{
				log.Printf("The HTTP request failed with error %s\n", err)
				log.Println("retry attempt: %v", retry)
				retry--
				n.GetName(retry, nameChannel)
			} else {
				data, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}

				if err := json.Unmarshal(data, &n); err != nil {
					log.Fatal(err)
				}
				log.Println("got name response")
				nameChannel <- &n

			}
}

func (j joke)GetJoke(retry int, randomName chan *nameInfo, w http.ResponseWriter, jokeChannel chan string) {

	for {

		name := <- randomName

		fmt.Println("channel gave this name: ", name.FirstName)

		url := "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]"
		log.Println("this is going to be the joke url: ", url)

		response, err := http.Get(url)
		if err != nil && retry > 0 {
			log.Printf("The HTTP request failed with error %s\n", err)
			log.Println("retry attempt: %v", retry)
			retry--
			j.GetJoke(retry, randomName, w, jokeChannel)
		} else {
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			if err := json.Unmarshal(data, &j); err != nil {
				log.Fatal(err)
			}

			log.Println(j.Value.Joke)

			JokePrint := strings.Replace(j.Value.Joke, "John", name.FirstName, -1)
			JokePrint = strings.Replace(JokePrint, "Doe", name.Surname, -1)

			log.Println(JokePrint)
			jokeChannel <- JokePrint

		}

	}
}

func Health(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Server", "Healthy")
	w.WriteHeader(200)
}
