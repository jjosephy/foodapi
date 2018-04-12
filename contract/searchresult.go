package contract

import (
	"strconv"

	"github.com/jjosephy/foodapi/logging"
)

// SearchResults list of SearchResult
type SearchResults struct {
	Results []SearchResult
}

// SearchResult for public API
type SearchResult struct {
	Title string `json:"name"`
	NDbno int    `json:"ndbno"`
}

// NewSearchResults creates a list of Search Results from a passed map
func NewSearchResults(m []interface{}) *SearchResults {
	if m == nil || len(m) == 0 {
		logging.WriteMessage("WARNING", "input param to NewSearchResult is nil")
		return &SearchResults{}
	}

	results := &SearchResults{
		Results: make([]SearchResult, len(m)),
	}

	for x := 0; x < len(m); x++ {
		n := m[x].(map[string]interface{})
		if n == nil {
			continue
		}

		res := new(SearchResult)
		res.Title = n["name"].(string)
		i, err := strconv.Atoi(n["ndbno"].(string))
		if err != nil {
			logging.WriteMessage("ERROR", "failed converting int on nbdno field")
		}
		res.NDbno = i

		results.Results[x] = *res
	}

	return results
}
