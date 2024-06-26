package main

func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2
	// 채널을 닫는다
	close(ch)

	// 채널 수신
	println(<-ch)
	println(<-ch)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신
	if _, success := <-ch; !success {
		println("더이상 데이타 없음.")
	}

	// 방법2
	// 위 표현과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}
