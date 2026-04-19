package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg" // JPEG画像のデコードに必要
	_ "image/png"  // PNG画像のデコードに必要
	"net/http"

	"github.com/fogleman/gg"
)

type MetObject struct {
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	PrimaryImage      string `json:"primaryImage"`
}

func main() {
	fmt.Println("Hello World")

	objectID := "436532"
	apiURL := fmt.Sprintf("https://collectionapi.metmuseum.org/public/collection/v1/objects/%s", objectID)

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var obj MetObject
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj.Title)
	fmt.Println(obj.PrimaryImage)
	imgResp, err := http.Get(obj.PrimaryImage)
	if err != nil {
		panic(err)
	}
	defer imgResp.Body.Close()

	img, _, err := image.Decode(imgResp.Body)
	const width = 1920
	const height = 1080
	dc := gg.NewContext(width, height)

	dc.DrawImageAnchored(img, width*3/4, height/2, 0.5, 0.5)

	outputFile := "wallpaper.png"

	if err := dc.SavePNG(outputFile); err != nil {
		panic(err)
	}

	fmt.Println("画像の保存が完了しました。")
}
