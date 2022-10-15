package twitterscraper

import "fmt"

var bearerToken2 = "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

// GetTrends return list of trends.
func (s *Scraper) GetTrends() ([]string, error) {
	req, err := s.newRequest("GET", "https://twitter.com/i/api/2/guide.json")
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("candidate_source", "trends")
	q.Add("include_page_configuration", "false")
	q.Add("entity_tokens", "false")
	req.URL.RawQuery = q.Encode()

	var jsn timeline
	s.setBearerToken(bearerToken2)
	err = s.RequestAPI(req, &jsn)
	s.setBearerToken(bearerToken)
	if err != nil {
		return nil, err
	}

	if len(jsn.Timeline.Instructions[1].AddEntries.Entries) < 2 {
		return nil, fmt.Errorf("no trend entries found")
	}

	var trends []string
	for _, item := range jsn.Timeline.Instructions[1].AddEntries.Entries[1].Content.TimelineModule.Items {
		trends = append(trends, item.Item.ClientEventInfo.Details.GuideDetails.TransparentGuideDetails.TrendMetadata.TrendName)
	}

	return trends, nil
}

// Deprecated: GetTrends wrapper for default Scraper
func GetTrends() ([]string, error) {
	return defaultScraper.GetTrends()
}
