# waterfall
implement a waterfall of go

usage:

Waterfall_Go([]Task{func(args ...interface{}) {
		callback, _ := args[0].(Callback)		
        
        // go async_op1()
        // result := <- somechan

		callback(false, 1, 2)
	}, func(args ...interface{}) {
		callback, _ := args[0].(Callback)
		x, _ := args[1].(int)
		y, _ := args[2].(int)		
		callback(false, x+y)
	}, func(args ...interface{}) {		
		go func() {			
			time.Sleep(time.Second * 2)
			callback, _ := args[0].(Callback)
			x, _ := args[1].(int)
			callback(false, x)			
		}()
	}}, func(args ...interface{}) {		
		fmt.Println(args...)
	})



Waterfall_Go([]Task{func(args ...interface{}) {
		callback, _ := args[0].(Callback)		
		callback(false, 1, 2)
	}, func(args ...interface{}) {
		callback, _ := args[0].(Callback)
		x, _ := args[1].(int)
		y, _ := args[2].(int)		
		callback(false, x+y)
	}, func(args ...interface{}) {		
		go func() {			
			time.Sleep(time.Second * 2)
			callback, _ := args[0].(Callback)
			x, _ := args[1].(int)
			callback(false, x)			
		}()
	}}, func(args ...interface{}) {		
		fmt.Println(args...)
	})
