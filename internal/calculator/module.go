package calculator

import (
	"math"
	"strconv"
	"strings"
)

type ModuleType int

const (
	Module10 ModuleType = 10
	Module11 ModuleType = 11
)

type DigitCalculatable interface {
	CalculateDigit(base string) string
}

type module struct {
	numbers       []int
	multipliers   []int
	additional    bool
	module        ModuleType
	singleSum     bool
	replacements  map[int]int
	sumMultiplier int
}

func NewModule(m ModuleType, number string) *module {
	dc := &module{}
	dc.module = m
	dc.numbers = convertToInt(strings.Split(reverseNumber(number), ""))
	dc.WithMultipliersInterval(2, 9)
	dc.MultiplySumBy(1)

	return dc
}

func NewModule10(number string) *module {
	return NewModule(Module10, number)
}

func NewModule11(number string) *module {
	return NewModule(Module11, number)
}

func (d *module) AddDigit(digit int) {
	d.numbers = append([]int{digit}, d.numbers...)
}

func (d *module) Calculate() int {
	sum := 0
	position := 0
	total := 0
	multiplier := 0
	for _, digit := range d.numbers {
		multiplier = d.multipliers[position]
		total = digit * multiplier
		sum += d.calculateSingleSum(total)
		position = d.nextMultiplier(position)
	}
	sum = d.calculateSumMultiplier(sum)
	result := sum % int(d.module)
	result = d.calculateAdditionalDigit(result)
	return d.replaceDigit(result)
}

func (d *module) calculateAdditionalDigit(digit int) int {
	if d.additional {
		return int(d.module) - digit
	}
	return digit
}

func (d *module) calculateSumMultiplier(sum int) int {
	return d.sumMultiplier * sum
}

func (d *module) calculateSingleSum(total int) int {
	if d.singleSum {
		return int(math.Abs(float64((total / 10) + (total % 10))))
	}
	return total
}

// Defines the multiplier factor after calculate the sum of digits.
func (d *module) MultiplySumBy(multiplier int) {
	d.sumMultiplier = multiplier
}

func (d *module) nextMultiplier(position int) int {
	position++
	if position == len(d.multipliers) {
		position = 0
	}
	return position
}

func (d *module) replaceDigit(digit int) int {
	v, ok := d.replacements[digit]
	if ok {
		return v
	}
	return digit
}

func (d *module) ReplaceWhen(replaceTo int, replacingNumbers ...int) {
	rep := make(map[int]int, 1)
	for _, number := range replacingNumbers {
		rep[number] = replaceTo
	}
	d.replacements = rep
}

func reverseNumber(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func (d *module) SingleSum() {
	d.singleSum = true
}

func (d *module) UseComplementaryInsteadOfModule() {
	d.additional = true
}

func (d *module) WithMultipliers(multipliers ...int) {
	d.multipliers = multipliers
}

func (d *module) WithMultipliersInterval(start int, end int) {
	idx := 0
	m := make([]int, (end-start)+1)
	for multiplier := start; multiplier <= end; multiplier++ {
		m[idx] = multiplier
		idx++
	}
	d.WithMultipliers(m...)
}

func convertToInt(numbers []string) []int {
	nbrs := make([]int, len(numbers))
	for i, number := range numbers {
		nb, _ := strconv.Atoi(number)
		nbrs[i] = nb
	}
	return nbrs
}
