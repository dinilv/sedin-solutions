/* array = [1, 2];
snap = Snapshot.new(array);
array[0] = 3;
array = snap.restore();
print(array.join(',')); #It should log "1,2"
array.push(4);
array = snap.restore();
print(array.join(',')); #It should log "1,2"
*/

// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

//snapshot repo
type SnapShot struct {
	OriginalArray []string
}

//initialize SnapShot Saver
func NewSnapShot(inputArray []string) SnapShot {
	snap := SnapShot{}
	snap.OriginalArray = duplicate(inputArray)
	return snap
}

//supply new copy from store
func (snap SnapShot) Restore() []string {
	return duplicate(snap.OriginalArray)
}

//copy the received array
func duplicate(inputArray []string) []string {
	copyArray := []string{}
	for _, elem := range inputArray {
		copyArray = append(copyArray, elem)
	}
	return copyArray
}

func main() {
	displayArray := []string{"1", "2"}
	snap := NewSnapShot(displayArray)
	displayArray[0] = "3"
	displayArray = snap.Restore()
	fmt.Println(strings.Join(displayArray, ","))
	displayArray = append(displayArray, "4")
	displayArray = snap.Restore()
	fmt.Println(strings.Join(displayArray, ","))
}
