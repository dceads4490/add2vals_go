/*********************************************************************
**
** Package declaration
**
**/
package main


/*********************************************************************
**      IMPORTS
*/
import (
	"testing" 
)


/*********************************************************************
**      Test Functions
*/

func Test_add2val(t *testing.T) {
    total := add2val(5, 5)
    if total != 10 {
       t.Errorf("Output was incorrect, got: %d, want: %d.", total, 10)
    }
}

