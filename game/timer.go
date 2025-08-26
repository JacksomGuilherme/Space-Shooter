package game

// Timer representa o objeto com os ticks decorridos e ticks necessários para uma ação ser executada
type Timer struct {
	CurrentTicks int
	TargetTicks  int
}

// NewTimer cria uma nova instância de Timer
func NewTimer(targetTicks int) *Timer {
	return &Timer{
		CurrentTicks: 0,
		TargetTicks:  targetTicks,
	}
}

// Update é responsável por atualizar os ticks do Timer
func (timer *Timer) Update() {
	if timer.CurrentTicks < timer.TargetTicks {
		timer.CurrentTicks++
	}
}

// IsReady é responsável por verificar se já se passaram a quantidade necessária de ticks para que uma ação seja executada
func (timer *Timer) IsReady() bool {
	return timer.CurrentTicks >= timer.TargetTicks
}

// Reset é responsável por reiniciar o valor de ticks decorridos desde a ultima ação
func (timer *Timer) Reset() {
	timer.CurrentTicks = 0
}
