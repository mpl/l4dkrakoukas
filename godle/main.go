package godle

import (
	"appengine"
//	"appengine/blobstore"
	"appengine/datastore"
//	"appengine/user"
//	"crypto/md5"
	"fmt"
//	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

// TODO: do it with bitfields?
const (
	monday = iota
	tuesday = iota
	wednesday = iota
	thursday = iota
	friday = iota
	saturday = iota
	sunday = iota
)

const (
	Asticot = iota
	ChuckMaurice = iota
	Posi = iota
	Lagoule = iota
)

const defaultMaxMemory = 32 << 20 // 32 MB

var (
	weekdays = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	allPlayers = []string{"Asticot", "ChuckMaurice", "Posi", "Lagoule"}
	allFalse = map[int]bool{
		monday: false,
		tuesday: false,
		wednesday: false,
		thursday: false,
		friday: false,
		saturday: false,
		sunday: false,
	}
)

func toString(day int) string {
	return weekdays[day]
}

func prettyDate(date string) string {
	return date[0:4] + ", " + date[4:6]
}


func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/week/", serveWeek)
	http.HandleFunc("/newweek", newWeek)
/*
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
*/
}

func serveError(c appengine.Context, w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Internal Server Error")
	c.Errorf("%v", err)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// TODO: the Schedule in this one could probably be a big JSON string
// derived from the Schedule in tplWeek
type Week struct {
	Date string
	Schedule []string
}

type tplWeek struct {
	Date string
	Schedule map[string]map[int]bool
}

func newWeek(w http.ResponseWriter, r *http.Request) {
	year, weeknb := time.Now().ISOWeek()
	weekId := fmt.Sprintf("%d%d", year, weeknb)
	emptySchedule := []string{}
	avail := map[string]map[int]bool{}
	for _, p := range allPlayers {
		emptySchedule = append(emptySchedule, "")
		avail[p] = allFalse
	}
	week := Week{Date: weekId, Schedule: emptySchedule}
	tWeek := tplWeek{weekId, avail}

	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "week", weekId, 0, nil)
	_, err := datastore.Put(c, key, &week)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := weekTemplate.Execute(w, tWeek); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveWeek(w http.ResponseWriter, r *http.Request) {
	_, weekId := path.Split(r.URL.String())
//	println(weekId)
	if weekId == "" {
		http.Error(w, "No week Id", http.StatusInternalServerError)
		return
	}
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "week", weekId, 0, nil)
	week := Week{}
	if err := datastore.Get(c, key, &week); err != nil {
		// TODO: log this?
//		http.Error(w, "Getting from the datastore: "+err.Error(), http.StatusInternalServerError)
		http.Error(w, "Invalid week Id", http.StatusInternalServerError)
		return
	}
	for k, v := range allPlayers {
		if r.Form == nil {
			r.ParseMultipartForm(defaultMaxMemory)
		}
		days := r.Form[v+"days"]
		if len(days) > 0 {
			week.Schedule[k] = strings.Join(days, " ")
		}
	}
	_, err := datastore.Put(c, key, &week)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	avail := map[string]map[int]bool{}
	for k, p := range allPlayers {
		days := strings.Fields(week.Schedule[k])
		okDays := map[int]bool{}
		for i:=0; i<7; i++ {
			okDays[i] = false
		}
		for _, sd := range days {
			dd, err := strconv.Atoi(sd)
			if err != nil {
				panic(err)
			}
			okDays[dd] = true
		}
		avail[p] = okDays
	}
	tWeek := tplWeek{weekId, avail}
	w.Header().Set("Content-Type", "text/html")
	if err := weekTemplate.Execute(w, tWeek); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
func isAnon(c appengine.Context) bool {
	u := user.Current(c)
	return (u == nil) 
}

func login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
//	c.Debugf(r.URL.Path)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	if r.URL.Path == "/login" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
//	c.Debugf(r.URL.Path)
	u := user.Current(c)
	if u != nil {
		url, err := user.LogoutURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	if r.URL.Path == "/logout" {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

*/
