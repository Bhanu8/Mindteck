package testing

import (
  "bytes"
  "mime/multipart"
  "os"
  "path/filepath"
  "io"
  "net/http"
  "io/ioutil"
  "testing"
  "mindteck/errors"
)

func TestUploadEndpoint(t *testing.T) {

  url := "192.168.1.42:9195/upload"
  method := "POST"
  path := "/home/bhanuteja/Pictures/Screenshot from 2020-06-03 12-00-00.png"
  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, er := os.Open(path)
  defer file.Close()
  part1,err := writer.CreateFormFile("img",filepath.Base(path))
  _, er = io.Copy(part1, file)
  errors.GetError(er)
  e := writer.Close()
  errors.GetError(e)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
  errors.GetError(err)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  errors.GetError(err)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  errors.GetError(err)
  t.Log(string(body))
}