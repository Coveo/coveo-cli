package main

type Query struct {
	Q               string `json:"q,omitempty"`
	NumberOfResults int    `json:"numberOfResults"`
	FirstResult     int    `json:"firstResult"`
}

type QueryResponse struct {
	TotalCount         int    `json:"totalCount"`
	TotalCountFiltered int    `json:"totalCountFiltered"`
	Duration           int    `json:"duration"`
	IndexDuration      int    `json:"indexDuration"`
	RequestDuration    int    `json:"requestDuration"`
	SearchUID          string `json:"searchUid"`
	GroupByResults     []struct {
		Field string `json:"field"`
	} `json:"groupByResults"`
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
