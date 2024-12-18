package dhtmlform

// TODO: formDataStore records expiration
var formDataStore map[string]*FormData

func init() {
	formDataStore = make(map[string]*FormData)
}
