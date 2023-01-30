package main

/*{Examples: The program should print “Found!” for the following example entered strings, 
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. 
The program should print “Not Found!” for the following strings, 
“ihhhhhn”, “ina”, “xian”. }*/

import "testing"
 
func TestConvert(t *testing.T) {
    const positives = [...] string {"ian", "Ian", "iuiygaygn", "I d skd a efju N"}
    const negatives = [...] string {"ihhhhhn", "ina", "xian", ""}

    for _, item := range positives {
        if (!checkPattern(item)) {
            t.Log("The string '" + item + "' should be positive")
            t.Fail()
        }   
    }

    for _, item := range negatives {
        if (checkPattern(item)) {
            t.Log("The string '" + item + "' should be negative")
            t.Fail()
        }   
    }

}