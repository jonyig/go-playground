package simulateHashmap

import (
	"log"
	"testing"
)

func TestCore(t *testing.T) {
	myMap := CreateHashMap()
	myMap.AddKeyValue("001", "1")
	log.Print(myMap.GetValueForKey("001"))
	myMap.AddKeyValue("001", "99")
	log.Print(myMap.GetValueForKey("001"))
	//myMap.AddKeyValue("002", "2")
	//myMap.AddKeyValue("003", "3")
	//myMap.AddKeyValue("004", "4")
	//myMap.AddKeyValue("005", "5")
	//myMap.AddKeyValue("006", "6")
	//myMap.AddKeyValue("007", "7")
	//myMap.AddKeyValue("008", "8")
	//myMap.AddKeyValue("009", "9")
	//myMap.AddKeyValue("010", "10")
	//myMap.AddKeyValue("011", "11")
	//myMap.AddKeyValue("012", "12")
	//myMap.AddKeyValue("013", "13")
	//myMap.AddKeyValue("012", "14")
	//myMap.AddKeyValue("015", "15")

}
