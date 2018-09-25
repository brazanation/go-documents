package brdocument

import (
	"math"
	"strconv"
	"strings"
)

const Module10 = 10
const Module11 = 11

type DigitCalculator struct {
	numbers       []int
	multipliers   []int
	additional    bool
	module        int
	singleSum     bool
	replacements  map[int]int
	sumMultiplier int
}

func NewDigitCalculator(number string) *DigitCalculator {
	dc := &DigitCalculator{}

	dc.numbers = convertToInt(strings.Split(reverseNumber(number), ""))
	dc.WithMultipliersInterval(2, 9)
	dc.WithModule(Module11)
	dc.MultiplySumBy(1)

	return dc
}

func (dc *DigitCalculator) AddDigit(digit int) {
	dc.numbers = append([]int{digit}, dc.numbers...)
}

func (dc *DigitCalculator) Calculate() int {
	sum := 0
	position := 0
	total := 0
	multiplier := 0
	for _, digit := range dc.numbers {
		multiplier = dc.multipliers[position]
		total = digit * multiplier
		sum += dc.calculateSingleSum(total)
		position = dc.nextMultiplier(position)
	}
	sum = dc.calculateSumMultiplier(sum)
	result := sum % dc.module
	result = dc.calculateAdditionalDigit(result)
	return dc.replaceDigit(result)
}

func (dc *DigitCalculator) calculateAdditionalDigit(digit int) int {
	if dc.additional {
		return dc.module - digit
	}
	return digit
}

func (dc *DigitCalculator) calculateSumMultiplier(sum int) int {
	return dc.sumMultiplier * sum
}

func (dc *DigitCalculator) calculateSingleSum(total int) int {
	if dc.singleSum {
		return int(math.Abs(float64((total / 10) + (total % 10))))
	}
	return total
}

// Defines the multiplier factor after calculate the sum of digits.
func (dc *DigitCalculator) MultiplySumBy(multiplier int) {
	dc.sumMultiplier = multiplier
}

func (dc *DigitCalculator) nextMultiplier(position int) int {
	position++
	if position == len(dc.multipliers) {
		position = 0
	}
	return position
}

func (dc *DigitCalculator) replaceDigit(digit int) int {
	v, ok := dc.replacements[digit]
	if ok {
		return v
	}
	return digit
}

func (dc *DigitCalculator) ReplaceWhen(replaceTo int, replacingNumbers []int) {
	rep := make(map[int]int, 1)
	for _, number := range replacingNumbers {
		rep[number] = replaceTo
	}
	dc.replacements = rep
}

func reverseNumber(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func (dc *DigitCalculator) SingleSum() {
	dc.singleSum = true
}

func (dc *DigitCalculator) UseComplementaryInsteadOfModule() {
	dc.additional = true
}

func (dc *DigitCalculator) WithModule(module int) {
	dc.module = module
}

func (dc *DigitCalculator) WithMultipliers(multipliers []int) {
	dc.multipliers = multipliers
}

func (dc *DigitCalculator) WithMultipliersInterval(start int, end int) {
	idx := 0
	m := make([]int, (end-start)+1)
	for multiplier := start; multiplier <= end; multiplier++ {
		m[idx] = multiplier
		idx++
	}
	dc.WithMultipliers(m)
}

func convertToInt(numbers []string) []int {
	nbrs := make([]int, len(numbers))
	for i, number := range numbers {
		nb, _ := strconv.Atoi(number)
		nbrs[i] = nb
	}
	return nbrs
}
