package main
import "fmt"
func  buildHeap(arr []int){
	length :=len(arr)
	for i:=length/2-1;i>=0;i--{
		toHeap(i,length-1,arr)
	}

}
func toHeap(i int,n int,arr []int){
	max :=i
	left:=2*i+1
	right:=2*i+2
	if left<=n && arr[max]<arr[left]{
		max=left
	}
	if right<=n && arr[max]<arr[right]{
		max=right
	}
	if max !=i{
		arr[i],arr[max]=arr[max],arr[i]
		toHeap(max,n,arr)
	}
}
func heapSort(arr []int){
	count :=len(arr)
	buildHeap(arr)
	for count>1{
		arr[0],arr[count-1]=arr[count-1],arr[0]
		count--
	}
}
func main(){
	arr := []int{1,3,2,4,5,8,6,7}
	heapSort(arr)
	fmt.Println(arr)


}