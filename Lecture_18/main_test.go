package main

import (
	"testing"
	"fmt"
)

func Benchmark100PrimesWith0MSSleep(b *testing.B){
	fmt.Println(primesAndSleep(100,0))
}

func Benchmark100PrimesWith5MSSleep(b *testing.B){
	fmt.Println(primesAndSleep(100,5))
}

func Benchmark100PrimesWith10MSSleep(b *testing.B){
	fmt.Println(primesAndSleep(100,10))
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B){
	fmt.Println(primesAndSleepGo(100,0))
}

func Benchmark100GoPrimesWith5MSSleep(b *testing.B){
	fmt.Println(primesAndSleepGo(100,5))
}

func Benchmark100GoPrimesWith10MSSleep(b *testing.B){
	fmt.Println(primesAndSleepGo(100,10))
}