package ContactList
import "fmt"
type ContactList struct {
    Name string
    phoneNumber string
}

func (c ContactList) GetContact(singleName, singleNumber string) string{
	c.Name = singleName
	c.phoneNumber = singleNumber
	return c.Name + c.phoneNumber	  
}

func main() {
	c := new(ContactList)
	res := c.GetContact("John", "123")
	fmt.Println(res)
}
