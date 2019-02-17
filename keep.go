
package main

// The keep type
type Keep struct {
    ID       string   `json:"id,omitempty"`
    Title    string   `json:"title,omitempty"`
    Message  string   `json:"message,omitempty"`
}

/**
 * Remove keep from keep list
 * @param id of keep
 * @param list of keep
 * @return list of keep
 */
func removeKeepFromList(id string, keeps []Keep) []Keep {
   for index, item := range keeps {
      if item.ID == id {
         keeps = append(keeps[:index], keeps[index+1:]...)
      }
    }

    return  keeps
}



