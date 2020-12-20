package view

import (
	"net/http"
	"io/ioutil"
	"strings"
	"mindteck/handlers"
	"mindteck/upload"
	"encoding/json"
	"mindteck/errors"
)

func SearchByID(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		id := r.FormValue("id")
		items, _ := ioutil.ReadDir("images")
		var image handlers.RetrieveImages
		var resp handlers.ImageResponse
		for _,item := range items {
			getID := strings.Split(item.Name(), ".")[0]
			if getID == id{
				image.ImageName = item.Name()
				image.ImagePath = "/images/"+item.Name()
				json.NewEncoder(w).Encode(image)
				return
			}else{
				resp.ID = id
				resp.Message = "Image Not Found for following requested ID"			
			}		
		}
		json.NewEncoder(w).Encode(resp)
	}else{
		var err []errors.Error
		err = upload.GetErrorCodeMessages(401,"Method Not Allowed",w)
		json.NewEncoder(w).Encode(err)
		return
	}
}