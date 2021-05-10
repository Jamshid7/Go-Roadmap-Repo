package ContactList
import (
	"testing"
)

func TestContactList(t *testing.T){
	c := new(ContactList)
	res := c.GetContact("John", "123")
	res1 := c.GetContact("Mike", "02173")
	res2 := c.GetContact("Joey", "98765")	
	if(res != "John123") {
		t.Error("Expected John123 but got error")
	}
	if(res1 != "Mike02173") {
		t.Error("Expected Mike02173 but got error")
	}
	if(res2 != "Joey9876") {
		t.Error("Expected Mike02173 but got error")
	}
}
