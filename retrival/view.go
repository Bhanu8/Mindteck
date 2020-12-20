package view

import (
	"net/http"
	"io/ioutil"
	// "time"
	"sort"
	"mindteck/handlers"
	"mindteck/upload"
	"mindteck/errors"
	"encoding/json"
	"strconv"
)

func RetrieveImages(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		items, _ := ioutil.ReadDir("images")
		//soting files based on timestamp desc
		sort.Slice(items, func(i,j int) bool{
			return items[i].ModTime().After(items[j].ModTime())
		})
		//getting limit to implement pagination
		limit, ok := r.URL.Query()["limit"]
		if !ok || len(limit[0]) < 1 {
			return
		}
		pagelimit,_ := strconv.Atoi(limit[0])
		if len(items) > pagelimit{		
		}else{
			pagelimit = len(items)
		}
		var images []handlers.RetrieveImages
		for i := 0; i<pagelimit; i++ {
			images = append(images, handlers.RetrieveImages{
				ImageName: items[i].Name(),
				ImagePath: "/images/"+items[i].Name(),
			})
		}
		json.NewEncoder(w).Encode(images)
		return
	}else{
		var err []errors.Error
		err = upload.GetErrorCodeMessages(401,"Method Not Allowed",w)
		json.NewEncoder(w).Encode(err)
		return
	}
}