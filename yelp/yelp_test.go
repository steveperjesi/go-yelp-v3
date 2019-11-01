package yelp

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

const (
    goodLat           = 36.0813328
    goodLon           = -115.3161651
    badLat            = -389.012456
    badLon            = -0.32801
    locationCity      = "denver"
    locationCityState = "las vegas, nv"
)

func TestYelpBadKey(t *testing.T) {
    assert := assert.New(t)

    authOptions, err := GetAuthOptions()
    assert.NoError(err)

    authOptions.ApiKey = "XYZ"
    client := NewClient(authOptions)
    assert.NotNil(client)
    _, err = client.Search(SearchOptions{
        Term:       "restaurants",
        Location:   locationCityState,
        Categories: "localservices",
        SortBy:     "distance",
        Limit:      50,
    })
    assert.Error(err)
}

func TestYelpLocation(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "child care",
        Location:   locationCityState,
        Categories: "localservices",
        SortBy:     "distance",
        Limit:      50,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpLatLong(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "plumbers",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Categories: "localservices",
        Limit:      50,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpRadius(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "auto repair",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Radius:     40000,
        Categories: "localservices",
        Limit:      50,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpNoResults(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "salamander",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Categories: "localservices",
        Limit:      150,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total == 0)
    assert.True(len(results.Businesses) == 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpOpenNow(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:      "hvac",
        Latitude:  goodLat,
        Longitude: goodLon,
        OpenNow:   true,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpMissingTerm(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Categories: "localservices",
        OpenNow:    false,
    })
    assert.EqualError(err, "Search term cannot be empty")
    assert.NotNil(results)
}

func TestYelpOpenAt(t *testing.T) {
    now := int(int32(time.Now().Unix()))
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "lawn care",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Categories: "localservices",
        OpenAt:     now,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpBadLat(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "septic",
        Latitude:   badLat,
        Categories: "localservices",
        OpenNow:    false,
    })
    assert.Error(err)
    assert.NotNil(results)
}

func TestYelpBadLon(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "hotel",
        Longitude:  badLon,
        Categories: "localservices",
        OpenNow:    false,
    })
    assert.Error(err)
    assert.NotNil(results)
}

func TestYelpMissingLocation(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "roofer",
        Categories: "localservices",
    })
    assert.Error(err)
    assert.NotNil(results)
}

func TestYelpLocale(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:      "ear nose throat",
        Latitude:  goodLat,
        Longitude: goodLon,
        Locale:    "en_US",
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpOffset(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:      "taxi",
        Latitude:  goodLat,
        Longitude: goodLon,
        Offset:    1,
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpPrice(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:      "contractor",
        Latitude:  goodLat,
        Longitude: goodLon,
        Price:     "1,2,3,4",
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}

func TestYelpAttributes(t *testing.T) {
    assert := assert.New(t)
    authOptions, err := GetAuthOptions()
    assert.NoError(err)
    client := NewClient(authOptions)
    assert.NotNil(client)
    results, err := client.Search(SearchOptions{
        Term:       "dentist",
        Latitude:   goodLat,
        Longitude:  goodLon,
        Attributes: "wheelchair_accessible",
    })
    assert.NoError(err)
    assert.NotNil(results)
    assert.True(results.Total > 0)
    assert.True(len(results.Businesses) > 0)
    assert.True(len(results.Businesses) <= 50)
}
