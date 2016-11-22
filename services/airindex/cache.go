package airindex

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var urlMapper = map[string]string{
	"beijing":   "http://www.stateair.net/web/rss/1/1.xml",
	"chengdu":   "http://www.stateair.net/web/rss/1/2.xml",
	"guangzhou": "http://www.stateair.net/web/rss/1/3.xml",
	"shanghai":  "http://www.stateair.net/web/rss/1/4.xml",
	"shenyang":  "http://www.stateair.net/web/rss/1/5.xml",
}

var urlMapperDebug = map[string]string{
	"beijing":   "http://0.0.0.0:8000/1.xml",
	"chengdu":   "http://0.0.0.0:8000/2.xml",
	"guangzhou": "http://0.0.0.0:8000/3.xml",
	"shanghai":  "http://0.0.0.0:8000/4.xml",
	"shenyang":  "http://0.0.0.0:8000/5.xml",
}

// AqiData cache the result from rss feed
type AqiData struct {
	Result    map[string][][]string
	Age       int64
	status    string
	lock      *sync.Mutex
	pollValue int64
	pollFreq  time.Duration
}

func fetchData(url string) ([]Item, error) {
	log.Printf("start to fetch %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := Rss{}
	err = xml.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	log.Printf("finish fetching %s\n", url)
	return res.Channel.Item, nil
}

// NewAqiData is
func NewAqiData(dev bool) *AqiData {
	ret := &AqiData{
		Result:    make(map[string][][]string),
		lock:      &sync.Mutex{},
		pollValue: 3600,
		pollFreq:  1810 * time.Second,
	}
	if !dev {
		go ret.UpdateData()
	}
	return ret
}

func (a *AqiData) fetchURLs() {
	var wg sync.WaitGroup
	var lock = &sync.Mutex{}
	for n, u := range urlMapper {
		wg.Add(1)
		go func(n, u string) {
			defer wg.Done()
			res, err := fetchData(u)
			if err != nil {
				log.Println(err)
				lock.Lock()
				a.Result[n] = nil
				lock.Unlock()
			} else {
				lock.Lock()
				a.Result[n] = a.parseDataToDraw(res)
				lock.Unlock()
			}
		}(n, u)
	}
	wg.Wait()
}

// UpdateData is
func (a *AqiData) UpdateData() {
	for {
		log.Println("Start to polling")
		if a.Age == 0 || time.Now().Unix()-a.Age > a.pollValue {
			log.Println("start to update data")
			a.lock.Lock()
			a.fetchURLs()
			a.Age = time.Now().Unix()
			a.lock.Unlock()
		} else {
			log.Println("data is new, continue")
		}
		time.Sleep(a.pollFreq)
	}
}

func (a AqiData) parseDataToDraw(data []Item) [][]string {
	timeSeria := make([]string, 0, 12)
	concSeria := make([]string, 0, 12)
	pm25Seria := make([]string, 0, 12)
	descSeria := make([]string, 0, 12)
	res := make([][]string, 0, 4)

	for _, i := range data {
		tmp := strings.Split(i.Description, "; ")
		if len(tmp) == 5 {
			timeSeria = append(timeSeria, tmp[0])
			concSeria = append(concSeria, tmp[2])
			pm25Seria = append(pm25Seria, tmp[3])
			descSeria = append(descSeria, tmp[4])
		} else if len(tmp) == 3 {
			timeSeria = append(timeSeria, tmp[0])
			concSeria = append(concSeria, "0")
			pm25Seria = append(pm25Seria, "0")
			descSeria = append(descSeria, "No Data")
		} else {
			panic("New error found, please see me")
		}

	}
	res = append(res, timeSeria, concSeria, pm25Seria, descSeria)
	return res
}

// CurrentData is
func (a *AqiData) CurrentData(id string) [][]string {
	return a.Result[id]
}
