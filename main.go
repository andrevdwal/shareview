package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/tabwriter"
	"time"
)

var (
	codeStr        = flag.String("c", "NFSWIX:SJ,DBXWD:SJ", "bloomberg instrument codes")
	updateInterval = flag.Duration("i", 1*time.Minute, "")
)

type quoteResult struct {
	SecurityType string     `json:"securityType"`
	BasicQuote   basicQuote `json:"basicQuote"`
}

type basicQuote struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	OpenPrice         int32     `json:"openPrice"`
	Price             int32     `json:"price"`
	PercentChange1Day float32   `json:"percentChange1Day"`
	LastUpdateISO     time.Time `json:"lastUpdateISO"`
}

func main() {
	flag.Parse()

	codes := strings.Split(*codeStr, ",")

	for {
		clearScreen()
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%s\t%s", "ID", "OPEN", "NOW", "CHANGE %", "DATE"))

		for _, c := range codes {
			res, err := query(c)
			if err != nil {
				fmt.Fprintln(w, fmt.Sprintf("%s\t%d\t%d\t%d\t%s", "", 0, 0, 0, err))
			} else {
				timeHere := res.BasicQuote.LastUpdateISO.In(time.Local)

				fmt.Fprintln(w, fmt.Sprintf("%s\t%d\t%d\t%f\t%s", res.BasicQuote.ID, res.BasicQuote.OpenPrice, res.BasicQuote.Price, res.BasicQuote.PercentChange1Day, timeHere))
			}
		}

		w.Flush()

		<-time.After(*updateInterval)
	}
}

func query(code string) (*quoteResult, error) {
	url := fmt.Sprintf("http://www.bloomberg.com/markets/api/quote-page/%s", code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res *quoteResult
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func clearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
