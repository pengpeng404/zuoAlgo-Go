package main

func main() {

}

/*
如何使用大根堆排序
1、把整个数组搞成大根堆 [0 - n-1]   heapSize == N
2、swap(0, n-1) heapSize--
3、heapify(0) swap(0, n-1) heapSize--
4、loop 直到 heapSize == 0
*/
