package main

import (
	"sort"
	"time"
	"math/rand"
	"fmt"
)

 func main(){
	for i :=0;i<5;i++{
		fmt.Println(getRecord())
	}
 }


 func getRecord() ([]int,[]int){
	reds := createBall(33)
	blue := createBall(16)
	reds = shuffle(reds)[0:6]
	sort.Ints(reds)
	blue = shuffle(blue)[0:1]
	return reds,blue
 }

 func shuffle(balls []int) []int{
	rand.Seed(time.Now().UnixNano()) //设置种子
    rand.Shuffle(len(balls), func(i, j int) { //调用算法
        balls[i], balls[j] = balls[j], balls[i]
	})
	return balls
 }

 func createBall(n int) []int{
	balls := make([]int,0)
	for i := 1 ; i<= n;i++{
		balls = append(balls,i)
	}
	return balls
 }