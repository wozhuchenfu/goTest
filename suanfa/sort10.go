package suanfa

import "fmt"

//选择排序
/*
选择排序的基本思想是对待排序的记录序列进行n-1遍的处理，第i遍处理是将L[i..n]中最小者与L[i]交换位置。这样，经过i遍处理之后，前i个记录的位置已经是正确的了。

选择排序是不稳定的。算法复杂度是O(n ^2 )。
个人总结：
选择排序，就是要又一个游标，每次在做比较的同时，纪录最大值，或者最小值的位置，遍历一遍之后，跟外层每次纪录的位置，做位置交换。为什么叫选择排序呢，估计就是这个原因，每次遍历一遍，选个最大或者最小的出来。算法因此得名。
 */
type SortInterface interface {
	sortByChoice(arry []int)
	sortByBubble(arry []int)
	sortByQuickly(arry []int)
	sortByInsert(arry []int)
}

type Sortor struct {
	name string
}

var arry = &[]int{3,4,21,2,6,9,7,5}

func SortedByChoice()  {
	learnsort := Sortor{name:"选择排序--从小到大--不稳定--n*n---"}
	learnsort.sortByChoice(*arry)
	fmt.Println(arry)
}

func (sortor *Sortor) sortByChoice(arr []int) {
	arrylength := len(arr)
	for i:=0;i<arrylength ;i++ {
		min := i
		for j:=i+1;j<arrylength ;j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		t:=arr[i]
		arr[i] = arr[min]
		arr[min] = t
	}
}

//冒泡排序 稳定的排序方法
func SortedByBubble()  {
	learnSort := Sortor{name:"冒泡排序--从小到大--不稳定--n*n---"}
	learnSort.sortByBubble(*arry)
	fmt.Println(arry)
}

func (sortor *Sortor) sortByBubble(arry []int) {
	arryLength := len(arry)
	done := true
	for i:=0;i<arryLength&&done ; i++ {
		done = false
		for j:=0;j<arryLength-i-j ;j++  {
			done = true
			if arry[j]>arry[j+1] {
				t := arry[j]
				arry[j] = arry[j+1]
				arry[j+1] = t
			}
		}
	}
}
/*
快速排序
快速排序由C. A. R. Hoare在1962年提出。它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

算法步骤：
设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据（通常选用数组的第一个数）作为关键数据，然后将所有比它小的数都放到它前面，所有比它大的数都放到它后面，这个过程称为一趟快速排序。值得注意的是，快速排序不是一种稳定的排序算法，也就是说，多个相同的值的相对位置也许会在算法结束时产生变动。
一趟快速排序的算法是：
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j–)，找到第一个小于key的值A[j]，将A[j]和A[i]互换；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]互换；
5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。

快速排序是不稳定的。最理想情况算法时间复杂度O(nlog2n)，最坏O(n ^2)。
个人总结：
快速排序其实是以数组第一数字作为中间值，开始排序，当这个值不是最佳中间值的时候，就会出现最坏的情况，当一次排序完成后，准备进入递归，递归传入的是slice，递归的退出条件是，当这个slice已经只能和自己比较了，也就是变为了中间值，slice为1。
 */

func SortedByQuick()  {

	learnsort := Sortor{name:"快速排序--从小到大--不稳定--nlog2n最坏n＊n---"}
	learnsort.sortByQuickly(*arry)
	fmt.Println(arry)
}

func (sortor *Sortor) sortByQuickly(arry []int) {
	if len(arry) <= 1 {
		return
	}
	mid := arry[0]
	i := 1//arry[0]为中间值mid，所以要从1开始比较
	head,tail := 0,len(arry)-1
	for head < tail {
		if arry[i] >mid {
			arry[i],arry[tail] = arry[tail],arry[i]
			tail--
		} else {
			arry[i],arry[head] = arry[head],arry[i]
			head++
			i++
		}
	}
	arry[head] = mid
	sortor.sortByQuickly(arry[:head])//这里的head就是中间值。左边是比它小的，右边是比他大的，开始递归
	sortor.sortByQuickly(arry[head+1:])
}

/*
插入排序
插入排序的基本思想是，经过i-1遍处理后,L[1..i-1]己排好序。
第i遍处理仅将L[i]插入L[1..i-1]的适当位置，使得L[1..i]又是排好序的序列。要达到这个目的，
我们可以用顺序比较的方法。首先比较L[i]和L[i-1]，如果L[i-1]≤ L[i]，则L[1..i]已排好序，
第i遍处理就结束了；否则交换L[i]与L[i-1]的位置，继续比较L[i-1]和L[i-2]，直到找到某一个位置j(1≤j≤i-1)，
使得L[j] ≤L[j+1]时为止。图1演示了对4个元素进行插入排序的过程，共需要(a),(b),(c)三次插入。
直接插入排序是稳定的。算法时间复杂度是O(n ^2) 。

个人总结：
算法的出发点在于一个需要排序的数值，然后依次对前面排序好的数据执行，不停的挪窝操作。这个要排序的值应该为1，因为0位置的那个数据肯定是排好序的呀
 */

func SortedByInsert()  {
	sortlearn := Sortor{name:"插入排序--从小到大--稳定--n＊n---"}
	sortlearn.sortByInsert(*arry)
	fmt.Println(arry)
}

func (sortor *Sortor) sortByInsert(arry []int) {
	arrylength := len(arry)
	for i,j:=1,0;i<arrylength ; i++ {//i从1开始，就是要插入的值。之所以从1开始。其实是要对挪窝做准备。因为接下来，我要对前面排好序的数据，依次挪窝。
		temp := arry[i]   //temp为开始排序的位置，从1开始。
		for j = i;j>0&&arry[j-1] > temp; j++ {//每次都从外层循环的计数器开始，跟一个temp变量比较，大的话，就往前挪一个窝。
			//因为前面都是排好序的，所以是依次挪窝，不会有数据丢失。
			arry[j] = arry[j-1]
		}
		arry[j] = temp//最后将挪出来的恰当位置给这个temp变量。也就是每次要插入的值。
	}
}















