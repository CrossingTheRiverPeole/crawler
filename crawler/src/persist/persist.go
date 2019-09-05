package persist

import "fmt"

func ItemPersist() chan interface{}  {
	itemChan := make(chan interface{})
	go func() {
		var itemCount int
		for   {
			item := <- itemChan
			fmt.Println(item)
			//
			itemCount++
		}



	}()
	return itemChan
}
