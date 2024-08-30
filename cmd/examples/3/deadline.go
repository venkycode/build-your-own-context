package main

import (
	"fmt"
	"time"
)

func deadlineUseCase() {
	t := time.NewTimer(1 * time.Second)
	defer t.Stop()
	done := make(chan struct{})
	go func() {
		<-t.C
		fmt.Println("Times up!")
		close(done)
	}()

	makeDinner(done)
}

func makeDinner(done chan struct{}) {
	select {
	case <-done:
		return
	default:
	}
	makeDaalTadka(done)
	makeRice(done)
	makeRoti(done)
	makeAlooGobi(done)

	select {
	case <-done:
		fmt.Println("could not finish dinner")
		return
	default:
		fmt.Println("Dinner ready")
	}
}

func makeRoti(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for roti")
		return
	default:
	}
	workFor(100)
	fmt.Println("Roti ready")
}

func makeAlooGobi(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for aloo gobi")
		return
	default:
	}
	prepareAloo(done)
	prepareGobi(done)

	select {
	case <-done:
		fmt.Println("could not finish aloo gobi")
		return
	default:
		fmt.Println("Aloo gobi ready")
	}
}

func prepareAloo(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for aloo")
		return
	default:
	}
	workFor(300)
	fmt.Println("Aloo prepared")
}

func prepareGobi(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for gobi")
		return
	default:
	}
	workFor(400)
	fmt.Println("Gobi prepared")
}

func makeRice(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for rice")
		return
	default:
	}
	workFor(500)
	fmt.Println("Rice cooked")
}

func makeDaalTadka(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for daal tadka")
		return
	default:
	}
	go prepareTadka(done)
	boilDaal(done)

	select {
	case <-done:
		fmt.Println("could not finish daal tadka")
		return
	default:
		fmt.Println("Daal tadka ready")
	}
}

func boilDaal(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for daal")
		return
	default:
	}
	workFor(1000)
	fmt.Println("Daal boiled")
}

func prepareTadka(done chan struct{}) {
	select {
	case <-done:
		fmt.Println("No time for tadka")
		return
	default:
	}
	workFor(200)
	fmt.Println("Tadka prepared")
}
