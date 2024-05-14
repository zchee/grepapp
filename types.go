// Copyright 2024 The grepapp Authors
// SPDX-License-Identifier: BSD-3-Clause

package main

type Grepapp struct {
	Facets  *Facets `json:"facets,omitempty"`
	Hits    *Hits   `json:"hits,omitempty"`
	Partial bool    `json:"partial,omitempty"`
	Time    int     `json:"time,omitempty"`
}

type Facets struct {
	Count int   `json:"count,omitempty"`
	Lang  *Lang `json:"lang,omitempty"`
	Path  *Path `json:"path,omitempty"`
	Repo  *Repo `json:"repo,omitempty"`
}

type Lang struct {
	Buckets []Bucket `json:"buckets,omitempty"`
}

type Bucket struct {
	Count int    `json:"count,omitempty"`
	Val   string `json:"val,omitempty"`
}

type Path struct {
	Buckets []Bucket `json:"buckets,omitempty"`
}

type Repo struct {
	Buckets []RepoBucket `json:"buckets,omitempty"`
}

type RepoBucket struct {
	Count   int    `json:"count,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
	Val     string `json:"val,omitempty"`
}

type Hits struct {
	Hits  []Hit `json:"hits,omitempty"`
	Total int   `json:"total,omitempty"`
}

type Hit struct {
	Branch       *Branch       `json:"branch,omitempty"`
	Content      *Content      `json:"content,omitempty"`
	ID           *ID           `json:"id,omitempty"`
	OwnerID      *OwnerID      `json:"owner_id,omitempty"`
	Path         *HitPath      `json:"path,omitempty"`
	Repo         *HitRepo      `json:"repo,omitempty"`
	TotalMatches *TotalMatches `json:"total_matches,omitempty"`
}

type Branch struct {
	Raw string `json:"raw,omitempty"`
}

type Content struct {
	Snippet string `json:"snippet,omitempty"`
}

type ID struct {
	Raw string `json:"raw,omitempty"`
}

type OwnerID struct {
	Raw string `json:"raw,omitempty"`
}

type HitPath struct {
	Raw string `json:"raw,omitempty"`
}

type HitRepo struct {
	Raw string `json:"raw,omitempty"`
}

type TotalMatches struct {
	Raw string `json:"raw,omitempty"`
}
