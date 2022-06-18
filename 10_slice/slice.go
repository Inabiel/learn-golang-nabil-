// slice
// Reference dari array, jika slice diganti maka array juga akan diganti,
// ada 3 buah hal yang perlu diperhatikan, yaitu pointer awal, length, dan capacitynya.
// ptr = index pertama slice, length = panjang slice, cap = panjang dari awal slice ke akhir array
// Jika capacity dari slice penuh, akan membuat array baru, jika tidak maka akan mengganti array awal.
// cara membuat adalah dengan namaArray[low:high]

package main

import "fmt"

func main() {
	var dynamicArr = []string{"a", "b", "c", "d", "e"}
	var slicedArr = dynamicArr[0:2]
	slicedArr[1] = "232"
	fmt.Println(dynamicArr, slicedArr, len(slicedArr), cap(slicedArr))

	var itsArray = [...]int{1,2,3} //Array, while its dynamic, cannot be appended
	var itsSlice = []int{1,2,3} //Slice is appended
	itsSlice2 := append(itsSlice, 5)
	fmt.Println(itsArray, len(itsArray), cap(itsArray), " ", itsSlice, len(itsSlice), cap(itsSlice), itsSlice2, len(itsSlice2), cap(itsSlice2))

	var itsArr = [4]string{"Januari", "Feb", "Mar", "Apr"}
	slicedArrFromMonth := itsArr[0:2]
	fmt.Println(itsArr)
	appendedSlice := append(slicedArrFromMonth, "May") // it still has capacity! so using old array/slice
	appendedSlice[0] = "Some Month"
	fmt.Println(slicedArrFromMonth, len(slicedArrFromMonth), cap(slicedArrFromMonth))
	fmt.Println(appendedSlice)

	newSlicedArrFromMonth := itsArr[3:] //only 1 cap/len, lets see appending on this one
	appendingNewSlice := append(newSlicedArrFromMonth, "Jun")
	appendingNewSlice[0] = "New Val" // lets see if it changes the old arr
	//Spoiler: it doesnt!
	fmt.Println(newSlicedArrFromMonth)
	fmt.Println(appendingNewSlice)
}