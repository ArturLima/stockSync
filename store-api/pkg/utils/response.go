package utils

type Response struct {
	Data interface{} `json:"data"`
}

func AsResult(obj interface{}) Response {
	return Response{
		Data: obj,
	}
}
