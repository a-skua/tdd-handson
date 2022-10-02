package lifegame

// positions

type X int

func (x X) Int() int {
	return int(x)
}

type Y int

func (y Y) Int() int {
	return int(y)
}

// Cell's State
type State bool

const (
	Live State = true
	Die  State = false
)

func (s State) IsLive() bool {
	return s == Live
}

func (s State) IsDie() bool {
	return s == Die
}

type Cell struct {
	x      X
	y      Y
	states []State
}

func NewCell(x X, y Y, seed []State) *Cell {
	return &Cell{x, y, seed}
}

func (c *Cell) inrangeX(x X) bool {
	return x >= 0 && x < c.x
}

func (c *Cell) outrangeX(x X) bool {
	return !c.inrangeX(x)
}

func (c *Cell) inrangeY(y Y) bool {
	return y >= 0 && y < c.y
}

func (c *Cell) outrangeY(y Y) bool {
	return !c.inrangeY(y)
}

// current states
func (c *Cell) current(x X, y Y) State {
	if c.outrangeX(x) || c.outrangeY(y) {
		return Die
	}

	return c.states[c.x.Int()*y.Int()+x.Int()]
}

// future state
func (c *Cell) future(x X, y Y) State {
	countAroundLives := func(x X, y Y) int {
		count := 0
		for _, state := range []State{
			c.current(x-1, y-1),
			c.current(x, y-1),
			c.current(x+1, y-1),
			c.current(x+1, y),
			c.current(x+1, y+1),
			c.current(x, y+1),
			c.current(x-1, y+1),
			c.current(x-1, y),
		} {
			if state {
				count += 1
			}
		}
		return count
	}(x, y)

	state := c.current(x, y)
	switch {
	case state.IsLive() && (countAroundLives == 2 || countAroundLives == 3),
		state.IsDie() && countAroundLives == 3:
		return Live
	default:
		return Die
	}
}

// Resizing
func (c *Cell) resize(x X, y Y) *Cell {
	states := make([]State, 0, x.Int()*y.Int())

	for cy := Y(0); cy < y; cy++ {
		for cx := X(0); cx < x; cx++ {
			states = append(states, c.current(cx, cy))
		}
	}

	return NewCell(x, y, states)
}

// management states
type Lifegame struct {
	cell *Cell
}

func New(cell *Cell) *Lifegame {
	return &Lifegame{cell}
}

func (lg *Lifegame) Next() {
	states := make([]State, 0, len(lg.cell.states))
	for y := Y(0); y < lg.cell.y; y++ {
		for x := X(0); x < lg.cell.x; x++ {
			states = append(states, lg.cell.future(x, y))
		}
	}
	lg.cell = NewCell(lg.cell.x, lg.cell.y, states)
}

func (lg *Lifegame) Table() [][]State {
	table := make([][]State, 0, lg.cell.y.Int())

	start, end := 0, lg.cell.x.Int()
	for y := Y(0); y < lg.cell.y; y++ {
		table = append(table, lg.cell.states[start:end])
		start += lg.cell.x.Int()
		end += lg.cell.x.Int()
	}
	return table
}
