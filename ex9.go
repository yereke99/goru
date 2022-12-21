package main

/*
type Movie struct {
	Title string `json:"title"`
}

type CachedPopularMovie struct {
	lock   sync.RWMutex
	Movies []Movie
}

func getPopularMoviesFromDB() []Movie {
	time.Sleep(5 * time.Second)
	return []Movie{{Title: "Avatar"}, {Title: "I Am Legend"}, {Title: "The Wolf of Wall Street"}}
}
func main() {
	ctx := context.Background()

	cache := CachedPopularMovie{}
	cache.Movies = getPopularMoviesFromDB()

	go func() {
		timer := time.NewTicker(1 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				movies := getPopularMoviesFromDB()

				// updating
				cache.lock.Lock()
				cache.Movies = movies
				cache.lock.Unlock()
			// app is terminating
			case <-ctx.Done():
				break
			}
		}
	}()

	http.HandleFunc("/getPopularMovies", func(w http.ResponseWriter, r *http.Request) {
		cache.lock.Lock()
		movies := cache.Movies
		cache.lock.Unlock()

		bytes_, _ := json.Marshal(movies)
		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes_)
	})

	_ = http.ListenAndServe(":8080", nil)

}
*/
