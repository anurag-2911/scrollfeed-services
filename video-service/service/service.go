package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"video-service/config"
	"video-service/model"
)

var cfg = config.Load()

// Legacy function for fetching categories - kept for compatibility
func FetchCategories(region string) ([]model.CategoryResponse, error) {
	log.Printf("[INFO] Fetching categories for region: %s", region)

	// Return predefined categories since we're now using MongoDB
	categories := []model.CategoryResponse{
		{ID: "10", Title: "Music"},
		{ID: "24", Title: "Entertainment"},
		{ID: "25", Title: "News & Politics"},
		{ID: "17", Title: "Sports"},
		{ID: "20", Title: "Gaming"},
		{ID: "23", Title: "Comedy"},
		{ID: "26", Title: "Howto & Style"},
		{ID: "27", Title: "Education"},
		{ID: "28", Title: "Science & Technology"},
	}

	log.Printf("[INFO] Retrieved %d predefined categories", len(categories))
	return categories, nil
}

// Legacy function for fetching regions - kept for compatibility
func FetchRegions() ([]string, error) {
	log.Printf("[INFO] Fetching supported regions")

	regions := []string{"US", "IN", "DE", "GB", "CA"}
	log.Printf("[INFO] Retrieved %d supported regions", len(regions))
	return regions, nil
}

// Still needed for comments functionality
func FetchComments(videoID string, maxResults int) (interface{}, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/commentThreads?part=snippet&videoId=%s&maxResults=%d&key=%s",
		videoID, maxResults, cfg.YouTubeAPIKey)

	log.Printf("[INFO] Fetching comments for video: %s", videoID)
	log.Printf("[DEBUG] Request URL: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch comments: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("[ERROR] YouTube API returned status: %d, body: %s", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("YouTube API error: %d", resp.StatusCode)
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("[ERROR] Failed to decode comments response: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully fetched comments for video: %s", videoID)
	return result, nil
}

// Still needed for video statistics functionality
func FetchVideoStatistics(videoID string) (interface{}, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=statistics&id=%s&key=%s",
		videoID, cfg.YouTubeAPIKey)

	log.Printf("[INFO] Fetching statistics for video: %s", videoID)
	log.Printf("[DEBUG] Request URL: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch video statistics: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("[ERROR] YouTube API returned status: %d, body: %s", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("YouTube API error: %d", resp.StatusCode)
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("[ERROR] Failed to decode statistics response: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully fetched statistics for video: %s", videoID)
	return result, nil
}

// SearchYouTubeVideos searches YouTube directly using the API
func SearchYouTubeVideos(query, region string, maxResults int) (interface{}, error) {
	apiURL := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&regionCode=%s&maxResults=%d&key=%s",
		url.QueryEscape(query), region, maxResults, cfg.YouTubeAPIKey)

	log.Printf("[INFO] Searching YouTube for query: %s, region: %s", query, region)
	log.Printf("[DEBUG] Request URL: %s", apiURL)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("[ERROR] Failed to search YouTube: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("[ERROR] YouTube API returned status: %d, body: %s", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("YouTube API error: %d", resp.StatusCode)
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("[ERROR] Failed to decode search response: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully searched YouTube for query: %s", query)
	return result, nil
}
