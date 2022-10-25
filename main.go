package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"webapiserver/query"

	"github.com/rs/cors"
)

type Read_Json struct {
	Live struct {
		Ssd struct {
			Util   []int `json:"UTIL"`
			CPU    []int `json:"CPU"`
			Memory []int `json:"MEMORY"`
			Power  []int `json:"POWER"`
		} `json:"SSD"`
		Csd struct {
			Util   []int `json:"UTIL"`
			CPU    []int `json:"CPU"`
			Memory []int `json:"MEMORY"`
			Power  []int `json:"POWER"`
		} `json:"CSD"`
	} `json:"Live"`
	Livecsd struct {
		Csd1 []int `json:"CSD1"`
		Csd2 []int `json:"CSD2"`
		Csd3 []int `json:"CSD3"`
		Csd4 []int `json:"CSD4"`
		Csd5 []int `json:"CSD5"`
		Csd6 []int `json:"CSD6"`
		Csd7 []int `json:"CSD7"`
		Csd8 []int `json:"CSD8"`
	} `json:"LIVECSD"`
	Query struct {
		H1 struct {
			Ssd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"Server"`
			Csd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"CSD"`
		} `json:"H1"`
		H5 struct {
			Ssd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"SSD"`
			Csd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"CSD"`
		} `json:"H5"`
		H11 struct {
			Ssd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"SSD"`
			Csd []struct {
				CPU     int `json:"CPU"`
				Memory  int `json:"MEMORY"`
				Power   int `json:"POWER"`
				Io      int `json:"IO"`
				Network int `json:"NETWORK"`
				Exetime int `json:"EXETIME"`
			} `json:"CSD"`
		} `json:"H11"`
	} `json:"QUERY"`
	Run struct {
		H1 []struct {
			Workid int    `json:"WORKID"`
			Smw    string `json:"SMW"`
			CPU    int    `json:"CPU"`
			Memory int    `json:"MEMORY"`
			Power  int    `json:"POWER"`
		} `json:"H1"`
		H4 []struct {
			Workid int    `json:"WORKID"`
			Smw    string `json:"SMW"`
			CPU    int    `json:"CPU"`
			Memory int    `json:"MEMORY"`
			Power  int    `json:"POWER"`
		} `json:"H4"`
		H5 []struct {
			Workid int    `json:"WORKID"`
			Smw    string `json:"SMW"`
			CPU    int    `json:"CPU"`
			Memory int    `json:"MEMORY"`
			Power  int    `json:"POWER"`
		} `json:"H5"`
		H11 []struct {
			Workid int    `json:"WORKID"`
			Smw    string `json:"SMW"`
			CPU    int    `json:"CPU"`
			Memory int    `json:"MEMORY"`
			Power  int    `json:"POWER"`
		} `json:"H11"`
	} `json:"RUN"`
}
type Return_Grid_Data struct {
	Result bool `json:"result"`
	Data   struct {
		Contents []struct {
			Querynum int    `json:"querynum"`
			Worknum  string `json:"worknum"`
			Smw      string `json:"smw"`
			CPU      string `json:"cpu"`
			Memory   string `json:"memory"`
			Power    string `json:"power"`
		} `json:"contents"`
		Pagination struct {
			Page       int `json:"page"`
			TotalCount int `json:"totalCount"`
		} `json:"pagination"`
	} `json:"data"`
}

type Return_Live_Data struct {
	Csd struct {
		Util   []int `json:"UTIL"`
		CPU    []int `json:"cpu"`
		Memory []int `json:"memory"`
		Power  []int `json:"power"`
		Num1   []int `json:"1"`
		Num2   []int `json:"2"`
		Num3   []int `json:"3"`
		Num4   []int `json:"4"`
		Num5   []int `json:"5"`
		Num6   []int `json:"6"`
		Num7   []int `json:"7"`
		Num8   []int `json:"8"`
	} `json:"csd"`
	Ssd struct {
		Util   []int `json:"UTIL"`
		CPU    []int `json:"cpu"`
		Memory []int `json:"memory"`
		Power  []int `json:"power"`
	} `json:"ssd"`
	Time struct {
		Timearray []string `json:"timearray"`
	} `json:"Time"`
}

type Return_DB_Data struct {
	Csd struct {
		CPU     []int `json:"cpu"`
		Memory  []int `json:"memory"`
		Power   []int `json:"power"`
		Network []int `json:"network"`
		IO      []int `json:"io"`
		Exetime []int `json:"exe"`
	} `json:"csd"`
	Ssd struct {
		CPU     []int `json:"cpu"`
		Memory  []int `json:"memory"`
		Power   []int `json:"power"`
		Network []int `json:"network"`
		IO      []int `json:"io"`
		Exetime []int `json:"exe"`
	} `json:"ssd"`
}

type CSD_Data struct {
	CPU     int `json:"cpu"`
	Memory  int `json:"memory"`
	Power   int `json:"power"`
	Network int `json:"network"`
	IO      int `json:"io"`
}

type Query struct {
	Name string `json:"query"`
}

type ServerLabels struct {
	Name []string `json:"Labels"`
}

type ResultData struct {
	Result string `json:"Result"`
}

var readdata Read_Json

var timearray []string
var resultflag = false
var queryresult string

func main() {

	mux := http.NewServeMux()
	file, _ := ioutil.ReadFile("data.json")
	// var readdata Read_Json
	json.Unmarshal(file, &readdata)
	// fmt.Println(readdata)
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	// fmt.Println(time.Now())

	mux.HandleFunc("/api/runData", func(w http.ResponseWriter, r *http.Request) {
		// data := r.URL.Query()["Query"]
		// fmt.Println(data[0])
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		for {
			if resultflag {
				break
			}
		}
		resultdata := ResultData{queryresult}

		retdata, _ := json.Marshal(resultdata)
		w.Write([]byte(retdata))
		resultflag = false
		// fmt.Fprintln(w, "Hello there!")

	})

	mux.HandleFunc("/api/readData", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// time.Sleep(5 * time.Second)
		// fmt.Println(r.URL.Query())
		data := r.URL.Query()["page[selectbox]"]
		server := r.URL.Query()["page[serverbox]"]
		// fmt.Println(server)
		switch server[0] {
		case "Simul":
			switch data[0] {
			case "H1":
				//append Simulator Code
			}
		case "SSD":
			switch data[0] {
			case "H4":
				query := query.Get_SSD_Query("4")
				querydata := Query{query}
				pbyte, _ := json.Marshal(querydata)
				buff := bytes.NewBuffer(pbyte)
				resp, err := http.Post("http://10.0.5.119:34568", "application/json", buff)
				if err != nil {
					fmt.Println(err)
				}
				defer resp.Body.Close()
				resultdata, _ := io.ReadAll(resp.Body)
				// fmt.Println(resultdata)
				queryresult = string(resultdata)
				resultflag = true
			}
		case "CSD":
			switch data[0] {
			case "H4":
				querydata := Query{"TPC-H_04"}
				pbyte, _ := json.Marshal(querydata)
				buff := bytes.NewBuffer(pbyte)
				// fmt.Println(string(buff.Bytes()))
				resp, err := http.Post("http://10.0.5.119:34568", "application/json", buff)
				if err != nil {
					fmt.Println(err)
				}
				defer resp.Body.Close()
				resultdata, _ := io.ReadAll(resp.Body)
				// fmt.Println(resultdata)
				queryresult = string(resultdata)
				resultflag = true
			}
		}
		// var jsondata string
		var jsondata Return_Grid_Data
		jsondata.Result = true
		jsondata.Data.Pagination.Page = 1
		jsondata.Data.Pagination.TotalCount = 100
		// curl -X GET -H "Content-Type: application/json" -d '{"query":"TPC-H_06"}' http://10.0.5.121:34568/
		// querys := url.Values{}
		// querys := make(map[string]string)
		// querys["query"] = "TPC-H_06"
		// reqbody, _ := json.Marshal(querys)
		// body := strings.NewReader(querys.Encode())
		// fmt.Println(body)
		// req, err := http.NewRequest("POST", "http://10.0.5.119:34568", querys)
		// req
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// resp, err := http.DefaultClient.Do(req)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer resp.Body.Close()

		// fmt.Println(resp.Body)

		// resp, err := http.PostForm("http://10.0.5.119:34568", url.Values{"query": {"TPC-H_04"}})
		// querydata := Query{"TPC-H_04"}
		// pbyte, _ := json.Marshal(querydata)
		// buff := bytes.NewBuffer(pbyte)
		// resp, err := http.Post("http://10.0.5.119:34568", "application/json", buff)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// defer resp.Body.Close()
		// respbody, _ := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(respbody))

		switch data[0] {
		case "H1":
			//QUERY1
			for i := 0; i < len(readdata.Run.H1); i++ {
				jsondata.Data.Contents = append(jsondata.Data.Contents, struct {
					Querynum int    "json:\"querynum\""
					Worknum  string "json:\"worknum\""
					Smw      string "json:\"smw\""
					CPU      string "json:\"cpu\""
					Memory   string "json:\"memory\""
					Power    string "json:\"power\""
				}{1, strconv.Itoa(readdata.Run.H1[i].Workid), readdata.Run.H1[i].Smw, strconv.Itoa(readdata.Run.H1[i].CPU), strconv.Itoa(readdata.Run.H1[i].Memory), strconv.Itoa(readdata.Run.H1[i].Power)})
			}
		case "H4":
			//QUERY1
			for i := 0; i < len(readdata.Run.H4); i++ {
				jsondata.Data.Contents = append(jsondata.Data.Contents, struct {
					Querynum int    "json:\"querynum\""
					Worknum  string "json:\"worknum\""
					Smw      string "json:\"smw\""
					CPU      string "json:\"cpu\""
					Memory   string "json:\"memory\""
					Power    string "json:\"power\""
				}{4, strconv.Itoa(readdata.Run.H4[i].Workid), readdata.Run.H4[i].Smw, strconv.Itoa(readdata.Run.H4[i].CPU), strconv.Itoa(readdata.Run.H4[i].Memory), strconv.Itoa(readdata.Run.H4[i].Power)})
			}
		case "H5":
			for i := 0; i < len(readdata.Run.H5); i++ {
				jsondata.Data.Contents = append(jsondata.Data.Contents, struct {
					Querynum int    "json:\"querynum\""
					Worknum  string "json:\"worknum\""
					Smw      string "json:\"smw\""
					CPU      string "json:\"cpu\""
					Memory   string "json:\"memory\""
					Power    string "json:\"power\""
				}{5, strconv.Itoa(readdata.Run.H5[i].Workid), readdata.Run.H5[i].Smw, strconv.Itoa(readdata.Run.H5[i].CPU), strconv.Itoa(readdata.Run.H5[i].Memory), strconv.Itoa(readdata.Run.H5[i].Power)})
			}
		case "H11":
			for i := 0; i < len(readdata.Run.H11); i++ {
				jsondata.Data.Contents = append(jsondata.Data.Contents, struct {
					Querynum int    "json:\"querynum\""
					Worknum  string "json:\"worknum\""
					Smw      string "json:\"smw\""
					CPU      string "json:\"cpu\""
					Memory   string "json:\"memory\""
					Power    string "json:\"power\""
				}{11, strconv.Itoa(readdata.Run.H11[i].Workid), readdata.Run.H11[i].Smw, strconv.Itoa(readdata.Run.H11[i].CPU), strconv.Itoa(readdata.Run.H11[i].Memory), strconv.Itoa(readdata.Run.H11[i].Power)})
			}
		}
		// tt, _ := url.Parse(data)
		// fmt.Println(tt.)
		// fmt.Println(data.String())
		retdata, _ := json.Marshal(jsondata)
		w.Write([]byte(retdata))
	})

	mux.HandleFunc("/api/DBData", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.URL.Query()["Query"])
		data := r.URL.Query()["Query"]
		// server := r.URL.Query()["Server"]
		// server := r.URL.Query()["Server"]
		// fmt.Println(data[0])
		// fmt.Println(server)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var dbdata Return_DB_Data
		// switch server[0] {
		// 	case "CSD1"
		// }
		switch data[0] {
		case "H1":
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[0].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[1].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[2].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[3].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[4].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[5].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[6].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H1.Csd[7].CPU)

			dbdata.Ssd.CPU = append(dbdata.Ssd.CPU, readdata.Query.H1.Ssd[0].CPU)

			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[0].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[1].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[2].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[3].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[4].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[5].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[6].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H1.Csd[7].Memory)

			dbdata.Ssd.Memory = append(dbdata.Ssd.Memory, readdata.Query.H1.Ssd[0].Memory)

			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[0].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[1].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[2].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[3].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[4].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[5].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[6].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H1.Csd[7].Power)

			dbdata.Ssd.Power = append(dbdata.Ssd.Power, readdata.Query.H1.Ssd[0].Power)

			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[0].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[1].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[2].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[3].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[4].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[5].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[6].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H1.Csd[7].Network)

			dbdata.Ssd.Network = append(dbdata.Ssd.Network, readdata.Query.H1.Ssd[0].Network)

			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[0].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[1].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[2].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[3].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[4].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[5].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[6].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H1.Csd[7].Io)

			dbdata.Ssd.IO = append(dbdata.Ssd.IO, readdata.Query.H1.Ssd[0].Io)

			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[0].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[1].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[2].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[3].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[4].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[5].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[6].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H1.Csd[7].Exetime)

			dbdata.Ssd.Exetime = append(dbdata.Ssd.Exetime, readdata.Query.H1.Ssd[0].Exetime)
		case "H5":
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[0].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[1].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[2].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[3].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[4].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[5].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[6].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H5.Csd[7].CPU)

			dbdata.Ssd.CPU = append(dbdata.Ssd.CPU, readdata.Query.H5.Ssd[0].CPU)

			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[0].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[1].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[2].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[3].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[4].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[5].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[6].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H5.Csd[7].Memory)

			dbdata.Ssd.Memory = append(dbdata.Ssd.Memory, readdata.Query.H5.Ssd[0].Memory)

			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[0].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[1].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[2].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[3].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[4].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[5].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[6].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H5.Csd[7].Power)

			dbdata.Ssd.Power = append(dbdata.Ssd.Power, readdata.Query.H5.Ssd[0].Power)

			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[0].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[1].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[2].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[3].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[4].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[5].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[6].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H5.Csd[7].Network)

			dbdata.Ssd.Network = append(dbdata.Ssd.Network, readdata.Query.H5.Ssd[0].Network)

			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[0].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[1].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[2].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[3].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[4].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[5].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[6].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H5.Csd[7].Io)

			dbdata.Ssd.IO = append(dbdata.Ssd.IO, readdata.Query.H5.Ssd[0].Io)

			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[0].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[1].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[2].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[3].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[4].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[5].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[6].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H5.Csd[7].Exetime)

			dbdata.Ssd.Exetime = append(dbdata.Ssd.Exetime, readdata.Query.H5.Ssd[0].Exetime)
		case "H11":
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[0].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[1].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[2].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[3].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[4].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[5].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[6].CPU)
			dbdata.Csd.CPU = append(dbdata.Csd.CPU, readdata.Query.H11.Csd[7].CPU)

			dbdata.Ssd.CPU = append(dbdata.Ssd.CPU, readdata.Query.H11.Ssd[0].CPU)

			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[0].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[1].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[2].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[3].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[4].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[5].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[6].Memory)
			dbdata.Csd.Memory = append(dbdata.Csd.Memory, readdata.Query.H11.Csd[7].Memory)

			dbdata.Ssd.Memory = append(dbdata.Ssd.Memory, readdata.Query.H11.Ssd[0].Memory)

			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[0].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[1].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[2].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[3].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[4].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[5].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[6].Power)
			dbdata.Csd.Power = append(dbdata.Csd.Power, readdata.Query.H11.Csd[7].Power)

			dbdata.Ssd.Power = append(dbdata.Ssd.Power, readdata.Query.H11.Ssd[0].Power)

			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[0].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[1].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[2].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[3].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[4].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[5].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[6].Network)
			dbdata.Csd.Network = append(dbdata.Csd.Network, readdata.Query.H11.Csd[7].Network)

			dbdata.Ssd.Network = append(dbdata.Ssd.Network, readdata.Query.H11.Ssd[0].Network)

			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[0].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[1].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[2].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[3].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[4].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[5].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[6].Io)
			dbdata.Csd.IO = append(dbdata.Csd.IO, readdata.Query.H11.Csd[7].Io)

			dbdata.Ssd.IO = append(dbdata.Ssd.IO, readdata.Query.H11.Ssd[0].Io)

			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[0].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[1].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[2].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[3].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[4].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[5].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[6].Exetime)
			dbdata.Csd.Exetime = append(dbdata.Csd.Exetime, readdata.Query.H11.Csd[7].Exetime)

			dbdata.Ssd.Exetime = append(dbdata.Ssd.Exetime, readdata.Query.H11.Ssd[0].Exetime)

		}
		retdata, _ := json.Marshal(dbdata)
		w.Write([]byte(retdata))
		// fmt.Fprintln(w, "Hello there!")

	})

	mux.HandleFunc("/api/ChartData", func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println(1)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var retdata Return_Live_Data
		// for i := 0; i < 9; i++ {
		// 	retdata.Csd.CPU = append(retdata.Csd.CPU, readdata.Live.CPU...)
		// }
		retdata.Csd.CPU = append(retdata.Csd.CPU, readdata.Live.Csd.CPU...)
		retdata.Csd.Memory = append(retdata.Csd.Memory, readdata.Live.Csd.Memory...)
		retdata.Csd.Power = append(retdata.Csd.Power, readdata.Live.Csd.Memory...)
		retdata.Ssd.CPU = append(retdata.Ssd.CPU, readdata.Live.Ssd.CPU...)
		retdata.Ssd.Memory = append(retdata.Ssd.Memory, readdata.Live.Ssd.Memory...)
		retdata.Ssd.Power = append(retdata.Ssd.Power, readdata.Live.Ssd.Power...)
		retdata.Csd.Util = readdata.Live.Csd.Util
		retdata.Ssd.Util = readdata.Live.Ssd.Util
		retdata.Csd.Num1 = append(retdata.Csd.Num1, readdata.Livecsd.Csd1...)
		retdata.Csd.Num2 = append(retdata.Csd.Num2, readdata.Livecsd.Csd2...)
		retdata.Csd.Num3 = append(retdata.Csd.Num3, readdata.Livecsd.Csd3...)
		retdata.Csd.Num4 = append(retdata.Csd.Num4, readdata.Livecsd.Csd4...)
		retdata.Csd.Num5 = append(retdata.Csd.Num5, readdata.Livecsd.Csd5...)
		retdata.Csd.Num6 = append(retdata.Csd.Num6, readdata.Livecsd.Csd6...)
		retdata.Csd.Num7 = append(retdata.Csd.Num7, readdata.Livecsd.Csd7...)
		retdata.Csd.Num8 = append(retdata.Csd.Num8, readdata.Livecsd.Csd8...)
		retdata.Time.Timearray = append(retdata.Time.Timearray, timearray...)
		// fmt.Fprintln(w, "Hello there!")
		// update()
		chartdata, _ := json.Marshal(retdata)
		w.Write([]byte(chartdata))

	})

	mux.HandleFunc("/api/CSDData", func(w http.ResponseWriter, r *http.Request) {
		// data := r.URL.Query()["Query"]
		// fmt.Println(data[0])
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// data := r.Body
		// json.Unmarshal()
		var dbdata Return_DB_Data
		// switch data[0] {
		// case "1":

		// case "2":

		// case "3":

		// case "4":

		// case "5":

		// case "6":

		// case "7":

		// case "8":

		// }
		retdata, _ := json.Marshal(dbdata)
		w.Write([]byte(retdata))
		// fmt.Fprintln(w, "Hello there!")

	})

	mux.HandleFunc("/api/ServerData", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query()["Server"]
		// fmt.Println(data[0])
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// data := r.Body
		// json.Unmarshal()
		var dbdata ServerLabels
		var tmplables []string
		switch data[0] {
		case "CSD1":
			tmplables = append(tmplables, "CSD1")
			tmplables = append(tmplables, "CSD2")
			tmplables = append(tmplables, "CSD3")
			tmplables = append(tmplables, "CSD4")
			tmplables = append(tmplables, "CSD5")
			tmplables = append(tmplables, "CSD6")
			tmplables = append(tmplables, "CSD7")
			tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "CSD Server")
		case "CSD2":
			tmplables = append(tmplables, "CSD1")
			tmplables = append(tmplables, "CSD2")
			tmplables = append(tmplables, "CSD3")
			tmplables = append(tmplables, "CSD4")
			tmplables = append(tmplables, "CSD5")
			tmplables = append(tmplables, "CSD6")
			tmplables = append(tmplables, "CSD7")
			tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "CSD Server")
		case "CSD3":
			tmplables = append(tmplables, "CSD1")
			tmplables = append(tmplables, "CSD2")
			tmplables = append(tmplables, "CSD3")
			tmplables = append(tmplables, "CSD4")
			tmplables = append(tmplables, "CSD5")
			tmplables = append(tmplables, "CSD6")
			tmplables = append(tmplables, "CSD7")
			tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "CSD Server")
		case "SSD1":
			// tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "SSD Server")
		case "SSD2":
			// tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "SSD Server")
		case "SSD3":
			// tmplables = append(tmplables, "CSD8")
			tmplables = append(tmplables, "SSD Server")
		}
		dbdata = ServerLabels{tmplables}
		retdata, _ := json.Marshal(dbdata)
		w.Write([]byte(retdata))
		// fmt.Fprintln(w, "Hello there!")

	})

	handler := cors.Handler(mux)
	go update()
	http.ListenAndServe("10.0.6.1:10111", handler)
}

func update() {
	for {
		time.Sleep(10 * time.Second)
		t := time.Now()
		// fmt.Println()
		// tmptime := strings.Split(t.Format("15:04:05"), ":")
		// tmptime1 := tmptime[1] + " " + tmptime[2]
		timearray = append(timearray, t.Format("15:04:05"))
		if len(timearray) > 9 {
			timearray = timearray[1:]
		}
		var tmpcsd []int
		var tmpssd []int
		csd := rand.Intn(10)
		ssd := rand.Intn(20)
		tmpcsd = append(tmpcsd, 100-csd)
		tmpssd = append(tmpssd, 100-ssd)
		tmpcsd = append(tmpcsd, csd)
		tmpssd = append(tmpssd, ssd)
		readdata.Live.Csd.Util = tmpcsd
		readdata.Live.Ssd.Util = tmpssd

		readdata.Live.Csd.CPU = readdata.Live.Csd.CPU[1:]
		readdata.Live.Csd.CPU = append(readdata.Live.Csd.CPU, rand.Intn(30))

		readdata.Live.Csd.Memory = readdata.Live.Csd.Memory[1:]
		readdata.Live.Csd.Memory = append(readdata.Live.Csd.Memory, rand.Intn(30))

		readdata.Live.Csd.Power = readdata.Live.Csd.Power[1:]
		readdata.Live.Csd.Power = append(readdata.Live.Csd.Power, rand.Intn(30))

		readdata.Live.Ssd.CPU = readdata.Live.Ssd.CPU[1:]
		readdata.Live.Ssd.CPU = append(readdata.Live.Ssd.CPU, rand.Intn(50))

		readdata.Live.Ssd.Memory = readdata.Live.Ssd.Memory[1:]
		readdata.Live.Ssd.Memory = append(readdata.Live.Ssd.Memory, rand.Intn(50))

		readdata.Live.Ssd.Power = readdata.Live.Ssd.Power[1:]
		readdata.Live.Ssd.Power = append(readdata.Live.Ssd.Power, rand.Intn(50))

		readdata.Livecsd.Csd1 = readdata.Livecsd.Csd1[1:]
		readdata.Livecsd.Csd1 = append(readdata.Livecsd.Csd1, rand.Intn(50))

		readdata.Livecsd.Csd2 = readdata.Livecsd.Csd2[1:]
		readdata.Livecsd.Csd2 = append(readdata.Livecsd.Csd2, rand.Intn(50))

		readdata.Livecsd.Csd3 = readdata.Livecsd.Csd3[1:]
		readdata.Livecsd.Csd3 = append(readdata.Livecsd.Csd3, rand.Intn(50))

		readdata.Livecsd.Csd4 = readdata.Livecsd.Csd4[1:]
		readdata.Livecsd.Csd4 = append(readdata.Livecsd.Csd4, rand.Intn(50))

		readdata.Livecsd.Csd5 = readdata.Livecsd.Csd5[1:]
		readdata.Livecsd.Csd5 = append(readdata.Livecsd.Csd5, rand.Intn(50))

		readdata.Livecsd.Csd6 = readdata.Livecsd.Csd6[1:]
		readdata.Livecsd.Csd6 = append(readdata.Livecsd.Csd6, rand.Intn(50))

		readdata.Livecsd.Csd7 = readdata.Livecsd.Csd7[1:]
		readdata.Livecsd.Csd7 = append(readdata.Livecsd.Csd7, rand.Intn(50))

		readdata.Livecsd.Csd8 = readdata.Livecsd.Csd8[1:]
		readdata.Livecsd.Csd8 = append(readdata.Livecsd.Csd8, rand.Intn(50))
	}
}
