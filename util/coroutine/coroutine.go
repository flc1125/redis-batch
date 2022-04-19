package coroutine

type Coroutine struct {
	chanel chan bool
}

func New(limit int) *Coroutine {
	return &Coroutine{
		chanel: make(chan bool, limit),
	}
}

func (this *Coroutine) Run(fc func(co *Coroutine)) {
	this.chanel <- true

	go func() {
		fc(this)

		<-this.chanel
	}()
}
