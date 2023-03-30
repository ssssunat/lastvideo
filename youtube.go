package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"errors"
)

const  YOTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const YOUTUBE_API_TOKEN = "AIzaSyCWsPZSOzIqvF_i4bxIQUaCU5ndfTMgFXA"
const  YOUTUBE_VIDEO_URL = "https://www.youtube.com/watch?v="


//возвращение видео 
func GetLastVideo(channelUrl string) (string, error) {
	items, err := retrieveVideos(channelUrl)
	if err != nil {
		return "", err
	}
	if len(items) < 1 {
		return "", errors.New("No video founded ")
	}
	return YOUTUBE_VIDEO_URL + items[0].Id.VideoId, nil
}


//возвращение видео согласно фильтрам
func retrieveVideos(channelUrl string) ([]Item, error) {
	req, err := makeRequest(channelUrl, 1)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Items, nil
}


func makeRequest(channelUrl string, maxResults int) (*http.Request, error) {
	lastSlashIndex := strings.LastIndex(channelUrl, "/")
	channelId := channelUrl[lastSlashIndex + 1 :]
	req, err := http.NewRequest("GET", YOTUBE_SEARCH_URL, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("part","id")
	query.Add("channelId", channelId)
	query.Add("maxResults", strconv.Itoa(maxResults))
	query.Add("ordet","date")
	query.Add("key",YOUTUBE_API_TOKEN)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())	
	return req, nil
}