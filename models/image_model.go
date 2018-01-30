package models

import ()

type Image struct {
	Title 			string 			`json:"title"`
	Url 			string 			`json:"url"`
	ThumbnailUrl	string 			`json:"thumbnail_url"`
	PreviewUrl		string 			`json:"preview_url"`
}

func (i *Image) Parse(image_payload map[string]interface{}) error {

	for k, v := range image_payload {
		if k == "Title" && v != nil {
			i.Title = v.(string)
		} else if k == "URL" {
			i.Url = v.(string)
		} else if k == "ThumbnailURL" {
			i.ThumbnailUrl = v.(string)
		} else if k == "PreviewURL" {
			i.PreviewUrl = v.(string)
		}
	}

	return nil
}
