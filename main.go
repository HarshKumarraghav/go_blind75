package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/valyala/fastjson"
)

const url = "https://leetcode.com/graphql"

func main() {
	w := csv.NewWriter(os.Stdout)
	header := []string{"slug", "link", "difficulty"}
	if err := w.Write(header); err != nil {
		panic(err)
	}

	for _, question := range questions {
		query := getQuery(question)
		queryString, err := json.Marshal(query)
		if err != nil {
			panic(err)
		}

		queryBytes := []byte(queryString)
		rw := ioutil.NopCloser(bytes.NewBuffer(queryBytes))
		req, err := http.NewRequest(http.MethodGet, url, rw)
		req.Header.Add("Content-Type", "application/json")
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			fmt.Printf("failed for question %q with response %q", question, body)
			continue
		}

		row := []string{
			fastjson.GetString(body, "data", "question", "titleSlug"),
			"https://leetcode.com/problems/" + fastjson.GetString(body, "data", "question", "titleSlug"),
			fastjson.GetString(body, "data", "question", "difficulty"),
		}
		if err := w.Write(row); err != nil {
			panic(err)
		}
	}

	w.Flush()
}

func getQuery(question string) map[string]interface{} {
	return map[string]interface{}{
		"operationName": "questionData",
		"variables": map[string]interface{}{
			"titleSlug": question,
		},
		"query": "query questionData($titleSlug: String!) {question(titleSlug: $titleSlug) {    questionId    questionFrontendId   title    titleSlug   isPaidOnly    difficulty  }}",
	}
}

var questions = []string{
	"two-sum",
	"longest-substring-without-repeating-characters",
	"longest-palindromic-substring",
	"container-with-most-water",
	"3sum",
	"remove-nth-node-from-end-of-list",
	"valid-parentheses",
	"merge-two-sorted-lists",
	"merge-k-sorted-lists",
	"search-in-rotated-sorted-array",
	"combination-sum",
	"rotate-image",
	"group-anagrams",
	"maximum-subarray",
	"spiral-matrix",
	"jump-game",
	"merge-intervals",
	"insert-interval",
	"unique-paths",
	"climbing-stairs",
	"set-matrix-zeroes",
	"minimum-window-substring",
	"word-search",
	"decode-ways",
	"validate-binary-search-tree",
	"same-tree",
	"binary-tree-level-order-traversal",
	"maximum-depth-of-binary-tree",
	"construct-binary-tree-from-preorder-and-inorder-traversal",
	"best-time-to-buy-and-sell-stock",
	"binary-tree-maximum-path-sum",
	"valid-palindrome",
	"longest-consecutive-sequence",
	"clone-graph",
	"word-break",
	"linked-list-cycle",
	"reorder-list",
	"maximum-product-subarray",
	"find-minimum-in-rotated-sorted-array",
	"reverse-bits",
	"number-of-1-bits",
	"house-robber",
	"number-of-islands",
	"reverse-linked-list",
	"course-schedule",
	"implement-trie-prefix-tree",
	"design-add-and-search-words-data-structure",
	"word-search-ii",
	"house-robber-ii",
	"contains-duplicate",
	"invert-binary-tree",
	"kth-smallest-element-in-a-bst",
	"lowest-common-ancestor-of-a-binary-search-tree",
	"lowest-common-ancestor-of-a-binary-tree",
	"product-of-array-except-self",
	"valid-anagram",
	"meeting-rooms",
	"meeting-rooms-ii",
	"graph-valid-tree",
	"missing-number",
	"alien-dictionary",
	"encode-and-decode-strings",
	"find-median-from-data-stream",
	"longest-increasing-subsequence",
	"coin-change",
	"number-of-connected-components-in-an-undirected-graph",
	"counting-bits",
	"top-k-frequent-elements",
	"sum-of-two-integers",
	"pacific-atlantic-water-flow",
	"longest-repeating-character-replacement",
	"non-overlapping-intervals",
	"serialize-and-deserialize-bst",
	"subtree-of-another-tree",
	"palindromic-substrings",
	"longest-common-subsequence",
}
