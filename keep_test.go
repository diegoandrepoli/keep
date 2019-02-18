
package main

import (
   "testing"
)

/**
 * Test on remove keep from list
 */
func TestRemoveKeepFromList(t *testing.T) {

    //generate keep list
    var keeps []Keep
    keeps = append(keeps, Keep{ID: "1", Title: "My keep", Message: "Keep display message"})
    keeps = append(keeps, Keep{ID: "2", Title: "My second keep", Message: "Keep display on second keep message"})

    //remove item from keep list
    keeps = removeKeepFromList("2", keeps)

    //check result
    if(len(keeps) != 1){
       t.Fatalf("expected array size 1, but got size %v", len(keeps)) 
    }
}
