/*********************************************************************
**
** Package declaration
**
**/
package main

/*********************************************************************
**
**     FILE:       add2vals.go
**     SYSTEM:     Home 
**     RELEASE:    1
**     PURPOSE:    Simple demo program to build using jenkins
**     Change History:
**
**     CHG        DATE   WHO   COMMENTS
**     ---     --------  ---   --------
**      0      05/07/20  dce   Initial code
*/

/*********************************************************************
**      IMPORTS
*/
import (
	"fmt" 
	"os"
	"strconv"
)

/*********************************************************************
**      GLOBAL DEFINITIONS
*/
const SUCCESS =0
const FAILURE =1
var    	local_rc int = SUCCESS



/**PROC+**********************************************************************/
/*                                                                           */
/* Name:        main                                                         */
/* Purpose:     Does the processing for add2vals                             */
/* Returns:     SUCCESS or FAILURE                                           */
/* Params:      Command Line Options                                         */
/*               See usage section                                           */
/*                                                                           */
/**PROC-**********************************************************************/
func main() {


    /* ============================================================
     *  Check for correct usage and parse values
     */
     if (len(os.Args) == 3){
	     // convert val to int
	     val1, err :=strconv.Atoi(os.Args[1])
	     if err != nil {
      		fmt.Printf("Error converting [%s] to int!\n",os.Args[1])
	        // Print usage
		usage()
		os.Exit(FAILURE)
   	     }
	     // convert val to int
	     val2, err :=strconv.Atoi(os.Args[2])
	     if err != nil {
      		fmt.Printf("Error converting [%s] to int!\n",os.Args[2])
	        // Print usage
		usage()
		os.Exit(FAILURE)
   	     }
	     fmt.Printf("%s: The result is %d.\n",os.Args[0],add2val(val1, val2))
     } else {
	     fmt.Println("Incorrect number of arguments!")
	     // Print usage
	     usage()
     }

    
    os.Exit(local_rc);
}
        

/**PROC+**********************************************************************/
/*                                                                           */
/* Name:        usage                                                        */
/* Purpose:     display command options                                      */
/* Returns:     None                                                         */
/* Params:      None                                                         */
/*                                                                           */
/**PROC-**********************************************************************/
func usage() {
	fmt.Printf("\nUsage:\n\tadd2vals val1 val2\n")
    	local_rc=FAILURE
}

/**FUNC+**********************************************************************/
/*                                                                           */
/* Name:        add2val                                                      */
/* Purpose:     adds 2 values and returns result                             */
/* Returns:     result                                                       */
/* Params:      int val1, int val2                                           */
/*                                                                           */
/**FUNC-**********************************************************************/
func add2val(val1 int, val2 int) int {
	return val1 + val2
}
