/*
 * PageSpeed, (C) 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	pagespeedonline "google.golang.org/api/pagespeedonline/v2"
)

// ResultRow  - results row.
type ResultRow struct {
	Strategy string `json:"strategy"`
	URL      string `json:"url"`
	Score    string `json:"score"`
}

// AnalyzeParam - analyze param.
type AnalyzeParam struct {
	target, strategy string
}

const (
	urlsFilePath   = "./urls.txt"
	resultFilePath = "./result.json"
	strategyMOBILE = "mobile"
	strategyPC     = "desktop"
)

func main() {

	fmt.Println("--- start ---")

	file, err := os.Open(urlsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		param := AnalyzeParam{target: scanner.Text(), strategy: strategyPC}
		writeJSON(analyze(param))
		param.strategy = strategyMOBILE
		writeJSON(analyze(param))
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- end ---")
}

func analyze(param AnalyzeParam) ResultRow {
	pso, err := pagespeedonline.New(&http.Client{
		Timeout: time.Duration(60) * time.Second,
	})
	if err != nil {
		panic(err)
	}

	r, err := pso.Pagespeedapi.Runpagespeed(param.target).Strategy(param.strategy).Do()
	if err != nil {
		panic(err)
	}

	return ResultRow{
		Strategy: param.strategy,
		URL:      r.Id,
		Score:    strconv.FormatInt(r.RuleGroups["SPEED"].Score, 10),
	}
}

func writeJSON(data ResultRow) {
	file, err := os.OpenFile(resultFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	e := json.NewEncoder(file)
	if err = e.Encode(&data); err != nil {
		log.Fatal(err)
	}
}
