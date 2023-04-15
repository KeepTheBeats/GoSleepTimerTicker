# GoSleepTimerTicker
Compare the performance of time.Sleep, time.Ticker, time.Timer

I set the time interval of the periodic task as _15 Minutes_, because I want to compare the performance while ***sleeping***, not ***executing***. 


I compare:
- sleepprogram/sleep.go
- tickerrangeprogram/ticker_range.go
- tickerselectprogram/ticker_select.go
- timerprogram/timer.go
- Shell sleep


### Result:

At first, I use Linux `top` to see the _CPU_ and _Memory_ consumption:

Linux `top` command sorted by **CPU**:
>almost no difference, all 0, including _Shell_ `sleep` (pid 3093)\
all lower than _Shell_ `sleep` (pid 3093)\
![top command sorted by CPU](results/top_sort_by_CPU.png "top command sorted by CPU")

Linux `top` command sorted by **Memory**:
>almost no difference, all 0, including _Shell_ `sleep` (pid 3093)\
all higher than _Shell_ `sleep` (pid 3093)\
![top command sorted by Memory](results/top_sort_by_memory.png "top command sorted by Memory")


Then, I use Linux the tool `pprof` of Go language to see the _CPU_ and _Memory_ consumption. The results are the same: 
>Any one of time.Sleep, time.Ticker (select), time.Ticker (range), or time.Timer **_does not consume any_** CPU or Memory while sleeping.
> 
> ![pprof_sleep_cpu](results/sleep_cpu.png "pprof_sleep_cpu")
> ![pprof_sleep_memory](results/sleep_memory.png "pprof_sleep_memory")
> ![pprof_ticker_range_cpu](results/ticker_range_cpu.png "pprof_ticker_range_cpu")
> ![pprof_ticker_range_memory](results/ticker_range_memory.png "pprof_ticker_range_memory")
> ![pprof_ticker_select_cpu](results/ticker_select_cpu.png "pprof_ticker_select_cpu")
> ![pprof_ticker_select_memory](results/ticker_select_memory.png "pprof_ticker_select_memory")
> ![pprof_timer_cpu](results/timer_cpu.png "pprof_timer_cpu")
> ![pprof_ticker_select_memory](results/timer_memory.png "pprof_timer_memory")


