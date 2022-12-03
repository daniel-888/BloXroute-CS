package workerpool

// import "fmt"

type WorkerPool struct {
	numWorkers int
	chTasks    chan func()
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		chTasks:    make(chan func()),
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.numWorkers; i++ {
		go func(count int) {
			for task := range w.chTasks {
				// fmt.Printf("this is worker %d\n", count)
				task()
			}
		}(i)
	}
}

func (w *WorkerPool) Quit() {
	close(w.chTasks)
}

func (w *WorkerPool) SubmitTask(task func()) {
	w.chTasks <- task
}
