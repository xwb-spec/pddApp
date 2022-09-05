package client

//type ImgUploadRequestParam struct {
//	File []byte
//	CommonRequestParam
//}
//type ImgUploadResponse struct {
//	GoodsImgUploadResponse GoodsImgUploadResponse `json:"goods_img_upload_response"`
//	ErrorResultResponse
//}
//
//type GoodsImgUploadResponse struct {
//	url string
//}
//
//func UploadImg(filePath string) (img ImgUploadResponse, err error) {
//	fileBytes, err := ioutil.ReadFile(filePath)
//	reqParam, err := json.Marshal(&ImgUploadRequestParam{
//		File: fileBytes,
//	})
//	resp, err := http.Post("https://gw-api.pinduoduo.com/api/router", "multipart/form-data", bytes.NewReader(reqParam))
//	if err != nil {
//		fmt.Println("err=", err)
//	}
//	defer resp.Body.Close()
//	bodyBytes, _ := ioutil.ReadAll(resp.Body)
//	if (resp.StatusCode == 200) || (resp.StatusCode == 201) {
//		_ = json.Unmarshal(bodyBytes, &img)
//		return img, nil
//	}
//	return ImgUploadResponse{}, err
//}
