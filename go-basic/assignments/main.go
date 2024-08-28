package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func main() {
	// fmt.Println("hello world")

	// fmt.Println(`arraySign {2, 1} | expected = 1, result =`, arraySign([]int{2, 1}))
	// fmt.Println(`arraySign {-2, 1} | expected = -1, result =`, arraySign([]int{-2, 1}))
	// fmt.Println(`arraySign {-1, -2, -3, -4, 3, 2, 1} | expected = 1, resuit =`, arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))

	// fmt.Println(`==========================`)

	// fmt.Println(`isAnagram("anak", "kana") | expected = true, result =`, isAnagram("anak", "kana"))
	// fmt.Println(`isAnagram("anak", "mana") | expected = false, result =`, isAnagram("anak", "mana"))
	// fmt.Println(`isAnagram("anagram", "managra") | expected = true, result =`, isAnagram("anagram", "managra"))

	// fmt.Println(`==========================`)

	// fmt.Println(`findTheDifference("abcd", "abcde") | expected = e, result =`, string(byte(findTheDifference("abcd", "abcde"))))
	// fmt.Println(`findTheDifference("abcd", "abced") | expected = e, result =`, string(byte(findTheDifference("abcd", "abced"))))
	// fmt.Println(`findTheDifference("", "y") | expected = y, result =`, string(byte(findTheDifference("", "y"))))

	// fmt.Println(`==========================`)

	// fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3})) // true; 1, 3, 5 adalah baris aritmatik +2
	// fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9})) // true; 9, 5, 1 adalah baris aritmatik -4
	// fmt.Println(canMakeArithmeticProgression([]int{1, 2, 3, 8})) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	// fmt.Println(`==========================`)

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	result := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			fmt.Println(result)
			result++
		}
	}
	if result%2 != 0 {
		fmt.Println(result % 2)
		return -1
	}
	return 1
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// panjang len yang berbeda sudah pasti bukan anagram
	if len(s) != len(t) {
		return false
	}

	sourceArray := []byte(s)
	sort.Slice(sourceArray, func(i, j int) bool {
		// fmt.Println("source i ", sourceArray[i])
		// fmt.Println("source j ", sourceArray[j])
		// fmt.Println("compare source i<j", sourceArray[i] < sourceArray[j])
		// fmt.Println("==============")
		return sourceArray[i] < sourceArray[j]
	})

	// fmt.Println("################")
	targetArray := []byte(t)
	sort.Slice(targetArray, func(i, j int) bool {
		// fmt.Println("target i ", targetArray[i])
		// fmt.Println("target j ", targetArray[j])
		// fmt.Println("compare target i<j", targetArray[i] < targetArray[j])
		// fmt.Println("==============")
		return targetArray[i] < targetArray[j]
	})

	for i := 0; i < len(sourceArray); i++ {
		// fmt.Println("len sourceArray")
		// fmt.Println(len(sourceArray))
		// fmt.Println(sourceArray[i], targetArray[i])
		if sourceArray[i] != targetArray[i] {
			return false
		}
	}
	return true

}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	source := 0
	target := 0

	for _, v := range s {
		source += int(v)
		// fmt.Println(source)
	}
	for _, v := range t {
		target += int(v)
		// fmt.Println(target)
	}
	diff := int(math.Abs(float64(source) - float64(target)))
	// fmt.Println(diff)
	return byte(diff)
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	// fmt.Println(arr)              // sort ascending
	difference := arr[1] - arr[0] // starting difference
	// fmt.Println(arr[1], arr[0])
	// fmt.Println(difference)

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != difference { // sliding window
			// fmt.Println(arr[i+1], arr[i])
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

func (d *Deck) New() {
	d.cards = make([]Card, 0, 52)

	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}

}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]

}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[len(d.cards)-n:]

}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	for i := range d.cards {
		j := rand.Intn(len(d.cards) - i)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	d.cards = append(d.cards[n:], d.cards[:n]...)

}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	// top5Cards := deck.PeekTop(3)
	// for _, c := range top5Cards {
	// 	fmt.Println(c.ToString())
	// }
	fmt.Println("---\n")

	// fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	// fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	// fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	// fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	// top5Cards = deck.PeekTop(10)
	// for _, c := range top5Cards {
	// 	fmt.Println(c.ToString())
	// }

	fmt.Println("---\n")
	deck.New()
	// deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
