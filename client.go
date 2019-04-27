package main

type CounterClient struct {
	ID    int
	Index int
}

type CounterClient2 struct{}

func (c *CounterClient) GetCounter() int {
	// send http request
	return 233
}

func (c *CounterClient2) GetCounter(id int, index int) int {
	// send http request
	return 233
}

func CallGetCounter() {
	client := &CounterClient{
		ID:    1,
		Index: 233,
	}

	result := client.GetCounter()
}

func CallGetCounter2() {
	client := &CounterClient2{}
	result := client.GetCounter()
}
