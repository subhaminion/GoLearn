import "fmt"

var URLStore = make(map[string]string)

func main() {

}

func Set(key, value string) error{
	URLStore[key] = value
	return nil
}

func Get(key string) ()