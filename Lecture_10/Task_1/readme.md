Given the type:

type ConcurrentPrinter struct {
    sync.WaitGroup
    sync.Mutex
    /* Feel free to add your own state here */
}

func (cp *ConcurrentPrinter) printFoo(times int) { /* provide your implementation */ }

func (cp *ConcurrentPrinter) printBar(times int) { /* provide your implementation */ }


Extend the type and implement the two methods, such that:

times := 10
cp = &ConcurrentPrinter{}
cp.PrintFoo(times)
cp.PrintBar(times)
cp.Wait()

Will result in the program printing the following:

foobarfoobarfoobarfoobarfoobar

Both printFoo and printBar need to internally spawn a function that
runs inside a different goroutine. You can use everything provided by the
ConcurrentPrinter instance (including your own fields), to synchronize
control across the separate goroutines.

HINT: It is important to make goroutines sleep between iterations, in order
to make the scheduler switch to the other ones. You can use the following
block of code for sleep:

time.Sleep(10 * time.Millisecond)