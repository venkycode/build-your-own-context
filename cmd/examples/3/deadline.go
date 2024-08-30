package main

import (
	"fmt"
	"time"

	context "github.com/venkycode/build-your-own-context"
)

func deadlineUseCase(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	makeDinner(ctx)
}

func makeDinner(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
	}
	makeDaalTadka(ctx)
	makeRice(ctx)
	makeRoti(ctx)
	makeAlooGobi(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("could not finish dinner")
		return
	default:
		fmt.Println("Dinner ready")
	}
}

func makeRoti(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for roti")
		return
	default:
	}
	workFor(100)
	fmt.Println("Roti ready")
}

func makeAlooGobi(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for aloo gobi")
		return
	default:
	}
	prepareAloo(ctx)
	prepareGobi(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("could not finish aloo gobi")
		return
	default:
		fmt.Println("Aloo gobi ready")
	}
}

func prepareAloo(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for aloo")
		return
	default:
	}
	workFor(300)
	fmt.Println("Aloo prepared")
}

func prepareGobi(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for gobi")
		return
	default:
	}
	workFor(400)
	fmt.Println("Gobi prepared")
}

func makeRice(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for rice")
		return
	default:
	}
	workFor(500)
	fmt.Println("Rice cooked")
}

func makeDaalTadka(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for daal tadka")
		return
	default:
	}
	go prepareTadka(ctx)
	boilDaal(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("could not finish daal tadka")
		return
	default:
		fmt.Println("Daal tadka ready")
	}
}

func boilDaal(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for daal")
		return
	default:
	}
	workFor(1000)
	fmt.Println("Daal boiled")
}

func prepareTadka(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("No time for tadka")
		return
	default:
	}
	workFor(200)
	fmt.Println("Tadka prepared")
}
