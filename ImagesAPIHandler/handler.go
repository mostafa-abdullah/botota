package ImagesAPIHandler

import (
	"botota/models"
	"botota/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	endpoint    = "https://api.gettyimages.com/v3/search/images/creative?phrase="
	queryString = "&fields=preview,title&exclude_nudity=true"
	numImages   = 3
)

//GetImages gets searches for images using a certain phrase
func GetImages(ph string) []models.Image {

	url := endpoint + ph + queryString
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	utils.Check(err)

	req.Header.Set("Api-Key", os.Getenv("GETTY_API_KEY"))
	res, err := client.Do(req)
	defer res.Body.Close()
	utils.Check(err)

	body, err := ioutil.ReadAll(res.Body)
	utils.Check(err)

	var js map[string]interface{}
	json.Unmarshal(body, &js)

	var imgs []models.Image

	for i := 0; i < numImages; i++ {
		// Extract the URI and the Caption from the deep Json response body
		iurl := js["images"].([]interface{})[i].(map[string]interface{})["display_sizes"].([]interface{})[0].(map[string]interface{})["uri"].(string)
		icap := js["images"].([]interface{})[i].(map[string]interface{})["title"].(string)

		img := models.Image{Url: iurl, Caption: icap}

		imgs = append(imgs, img)
	}

	return imgs
}

func GetImagesMessages(destination string) []models.Message {
	imgs := GetImages(destination + "+landmarks")
	var res []models.Message

	for _, img := range imgs {
		msg := models.Message{Value: img.Caption, Image: img.Url}
		res = append(res, msg)
	}

	// Set the highlight for the first image
	res[0].Highlight = "We brought you some of the nicest images for " + destination
	return res
}
