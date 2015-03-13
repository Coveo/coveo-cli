package main

type Query struct {
	Q               string            `json:"q,omitempty"`
	NumberOfResults int               `json:"numberOfResults"`
	FirstResult     int               `json:"firstResult"`
	GroupBy         []*GroupByRequest `json:"groupBy"`
}

type GroupByRequest struct {
	Field                 string `json:"field"`
	MaximumNumberOfValues int    `json:"maximumNumberOfValues"`
	SortCriteria          string `json:"sortCriteria"`
	InjectionDepth        int    `json:"injectionDepth"`
}

func (q *Query) AddGroupByRequest(g *GroupByRequest) {
	q.GroupBy = append(q.GroupBy, g)
}

type QueryResponse struct {
	TotalCount         int             `json:"totalCount"`
	TotalCountFiltered int             `json:"totalCountFiltered"`
	Duration           int             `json:"duration"`
	IndexDuration      int             `json:"indexDuration"`
	RequestDuration    int             `json:"requestDuration"`
	SearchUID          string          `json:"searchUid"`
	GroupByResults     []GroupByResult `json:"groupByResults"`
	//Results []map[string]interface{} // Naive method
	Results []struct {
		Title          string                 `json:"title"`
		URI            string                 `json:"uri"`
		Excerpt        string                 `json:"excerpt"`
		FirstSentences string                 `json:"firstSentences"`
		Score          int                    `json:"score"`
		PercentScore   float32                `json:"percentScore"`
		Raw            map[string]interface{} `json:"raw"`
	} `json:"results"`
}

type GroupByResult struct {
	Field  string `json:"field"`
	Values []struct {
		Value           string `json:"value"`
		NumberOfResults int    `json:"numberOfResults"`
		Score           int    `json:"score"`
		ValueType       string `json:"valueType"`
	} `json:"values"`
}
