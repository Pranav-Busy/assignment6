package main

import (
	"fmt"
	"reflect"
)


func Remove_last_element (received_arr interface{},received_boolean bool)(interface{}){

	unfold:= reflect.ValueOf(received_arr)

	var updatedinterface interface{}

	switch unfold.Kind(){

	case reflect.Slice:

		nestedslice := unfold.Interface().([]interface{})

		if(received_boolean){
           nestedslice= nestedslice[1:]
		}else{
			nestedslice=nestedslice[:unfold.Len()-1]
		}

		return nestedslice


	default:
		return updatedinterface
	}
    

}

func main (){

	// []interface{}{1, 2, 3, "pranav"}
	var Original_interface interface{}= []interface{}{1, 2, 3, "pranav"}

      var received_boolean=false




updated_interface:=Remove_last_element(Original_interface,received_boolean)

fmt.Println(updated_interface)



	


}


