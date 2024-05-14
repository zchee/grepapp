// Copyright 2024 The grepapp Authors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
)

const endpoint = "https://grep.app/api/search"

var (
	lang string
)

func init() {
	spew.Config = spew.ConfigState{
		Indent:           " ",
		ContinueOnMethod: true,
		SortKeys:         true,
	}

	flag.StringVar(&lang, "lang", "", "filter language")
}

func main() {
	flag.Parse()
	query := flag.Arg(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v := make(url.Values)
	v.Set("q", query)
	v.Set("page", "1")
	if lang != "" {
		v.Set("lang", lang)
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	u.RawQuery = v.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	var m Grepapp
	if err := sonic.ConfigFastest.Unmarshal(buf, &m); err != nil {
		log.Fatal(err)
	}
	spew.Dump(m)
}
