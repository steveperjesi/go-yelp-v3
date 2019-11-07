package yelp

import (
    "bytes"
    "errors"
    "fmt"
    "strings"
)

type SearchOptions struct {
    Term       string
    Location   string
    Limit      int
    Latitude   float64
    Longitude  float64
    Radius     int
    Categories string
    Locale     string
    Offset     int
    SortBy     string
    Price      string
    OpenNow    bool
    OpenAt     int
    Attributes string
}

type SearchResponse struct {
    Total      int        `json:"total"`
    Businesses []Business `json:"businesses"`
    Region     Region     `json:"region"`
}

type ErrorResponse struct {
    Error ErrorData `json:"error"`
}

type ErrorData struct {
    Code        string      `json:"code"`
    Description string      `json:"description"`
    Field       string      `json:"field"`
    Instance    interface{} `json:"instance"`
}

// See https://www.yelp.com/developers/documentation/v3/business_search
func (client *Client) Search(request SearchOptions) (SearchResponse, error) {
    method := "GET"

    var endpoint bytes.Buffer
    endpoint.WriteString("/businesses/search?")

    if request.Term != "" {
        term := strings.Join(strings.Split(request.Term, " "), "+")
        endpoint.WriteString(fmt.Sprintf("term=%s", term))
    }

    if request.Location == "" {
        if request.Latitude == 0 || request.Longitude == 0 {
            err := errors.New("Search LOCATION or LATITUDE/LONGITUDE cannot be empty")
            return SearchResponse{}, err
        } else {
            lat := request.Latitude
            lon := request.Longitude
            endpoint.WriteString(fmt.Sprintf("&latitude=%.7f&longitude=%.7f", lat, lon))
        }
    } else {
        location := strings.Join(strings.Split(request.Location, " "), "+")
        endpoint.WriteString(fmt.Sprintf("&location=%s", location))
    }

    if request.Limit != 0 {
        max := 50
        limit := 0
        if request.Limit > 50 {
            limit = max
        } else {
            limit = request.Limit
        }

        endpoint.WriteString(fmt.Sprintf("&limit=%d", limit))
    }

    if request.Offset > 0 {
        endpoint.WriteString(fmt.Sprintf("&offset=%d", request.Offset))
    }

    if request.Radius > 0 {
        endpoint.WriteString(fmt.Sprintf("&radius=%d", request.Radius))
    }

    if request.SortBy != "" {
        endpoint.WriteString(fmt.Sprintf("&sort_by=%s", request.SortBy))
    }

    if request.OpenNow {
        endpoint.WriteString("&open_now=true")
    } else {
        if request.OpenAt > 0 {
            endpoint.WriteString(fmt.Sprintf("&open_at=%d", request.OpenAt))
        }
    }

    if request.Categories != "" {
        endpoint.WriteString(fmt.Sprintf("&categories=%s", request.Categories))
    }

    if request.Locale != "" {
        endpoint.WriteString(fmt.Sprintf("&locale=%s", request.Locale))
    }

    if request.Price != "" {
        endpoint.WriteString(fmt.Sprintf("&price=%s", request.Price))
    }

    if request.Attributes != "" {
        endpoint.WriteString(fmt.Sprintf("&attributes=%s", request.Attributes))
    }

    params := make(map[string]interface{})

    response := SearchResponse{}
    err := client.request(method, endpoint.String(), params, &response)
    if err != nil {
        return response, err
    }
    return response, nil
}
