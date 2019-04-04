package cache

/*
 * Thsi file cotains the definitions for the high performant cache
 */

//Type holds the request type of operation like read and write
type Type uint

const (
	//READ from the cache
	//if the key doesn't exist in the cache it will return nil
	READ Type = iota
	//WRITE into the cache
	WRITE
	//DELETE keys from the cache
	DELETE
)

//Request facilitates the read and write to the cache go routine
type Request struct {
	//Type is the type of the request
	Type Type
	//Payload to be written inot the cache the operation is write
	Payload interface{}
	//Key is the key with the payload to be accessed or written into
	Key string
	//Out chan is for getting the response in case of read operation
	Out chan Request
}

//NewRequest will generate a request with necessary initializations based on the type of the request
//Example, if the type is read will return a request with out channel initialized
func NewRequest(t Type) Request {
	switch t {
	case READ:
		return Request{Type: t, Out: make(chan Request)}
	default:
		return Request{Type: t}
	}
}

//RequestChannel through which the requests can be made
var RequestChannel = make(chan Request)

func init() {
	go Cache(RequestChannel)
}

//Cache is the in-memory cache go routine
func Cache(ch chan Request) {
	cache := LoadPersistant()
	for {
		req := <-ch
		switch req.Type {
		case READ:
			//Read operation
			req.Payload, _ = cache[req.Key]
			go SendToChannel(req.Out, req)
		case WRITE:
			//Write operation
			cache[req.Key] = req.Payload
		case DELETE:
			//Delete operation
			delete(cache, req.Key)
		}
	}
}

//SendToChannel will send a rquest to the channel. It is intended for use along with routines
//If you want regulate the spawning of the go routines, you can use a go routine pool design
func SendToChannel(ch chan Request, req Request) {
	ch <- req
}

//LoadPersistant loads the information to be loaded from persistant storage when cache goes live
func LoadPersistant() map[string]interface{} {
	//User can implement their own LoadPersistant
	//For experiment purposes I am loading foo-bar
	cache := map[string]interface{}{
		"Foo": "Bar",
	}

	return cache
}
