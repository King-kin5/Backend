package models

type Response struct{
	Data interface{} `json:"data" bson:"data"`
	Error string `json:"error" bson:"error"`
	OK bool `json:"ok"   bson:"ok"`
}
func NewResponse(data interface{}, err string, OK bool) *Response {
    return &Response{
        Data:  data,
        Error: err,
        OK:    OK,
    }
}