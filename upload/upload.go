package upload

import (
	"net/http"
	"mindteck/errors"
	"mindteck/handlers"
	"github.com/globalsign/mgo/bson"
	"path/filepath"
	"os"
	"encoding/json"
	"io"
	"fmt"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(1 << 4)
		errors.GetMultiPartFormError(err)
		file,handler,err := r.FormFile("img")
		errors.GetFileError(err)
		defer file.Close()
		ext := filepath.Ext(handler.Filename)
		var er []errors.Error
		switch ext {
			//validating type of file
			case ".jpeg", ".jpg", ".png", ".webp":
				path := "images"
				if _, err := os.Stat(path); err != nil {
					merr := os.Mkdir(path, 0700)
					errors.GetDirectoryError(merr)
				}
				fname := bson.NewObjectId().Hex()
				dst, err := os.Create(path + "/" + fname + ext)
				errors.GetError(err)
				defer dst.Close()
				if _, err := io.Copy(dst, file); err != nil {
					fmt.Println(err)
					return
				}
				var mesg []handlers.ImageResponse
				mesg = append(mesg, handlers.ImageResponse{
					Status: 200,
					ID: fname,
					Message: "success",
				})
				json.NewEncoder(w).Encode(mesg)
				return
			default:	
				//catching invalid type files			
				er = GetErrorCodeMessages(201,"Invalid Input File",w)
				json.NewEncoder(w).Encode(er)
				return
		}
	}else{
		var err []errors.Error
		err = GetErrorCodeMessages(401,"Method Not Allowed",w)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func GetErrorCodeMessages(status int,mesg string, w http.ResponseWriter) []errors.Error{
	var err []errors.Error
	err = append(err, errors.Error{
		Status: status,
		Message: mesg,
	})
	w.WriteHeader(status)
	return err
}

