package deck

//go:generate enumer -type=Suit
type Suit int

const (
	SuitNone Suit = iota
	SuitHearts
	SuitSpades
	SuitDiamonds
	SuiteClubs
)

//go:generate enumer -type=Rank
type Rank int

const (
	RankNone Rank = iota
	RankAce
	RankTwo
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankJoker
)

type Card struct {
	suit Suit
	rank Rank
	str  string
}

func (c *Card) String() string {
	return c.str
}

var kAllCards = []Card{
	kAceOfSpades,
	kTwoOfSpades,
	kThreeOfSpades,
	kFourOfSpades,
	kFiveOfSpades,
	kSixOfSpades,
	kSevenOfSpades,
	kEightOfSpades,
	kNineOfSpades,
	kTenOfSpades,
	kJackOfSpades,
	kQueenOfSpades,
	kKingOfSpades,
	kAceOfHearts,
	kTwoOfHearts,
	kThreeOfHearts,
	kFourOfHearts,
	kFiveOfHearts,
	kSixOfHearts,
	kSevenOfHearts,
	kEightOfHearts,
	kNineOfHearts,
	kTenOfHearts,
	kJackOfHearts,
	kQueenOfHearts,
	kKingOfHearts,
	kAceOfDiamonds,
	kTwoOfDiamonds,
	kThreeOfDiamonds,
	kFourOfDiamonds,
	kFiveOfDiamonds,
	kSixOfDiamonds,
	kSevenOfDiamonds,
	kEightOfDiamonds,
	kNineOfDiamonds,
	kTenOfDiamonds,
	kJackOfDiamonds,
	kQueenOfDiamonds,
	kKingOfDiamonds,
	kAceOfClubs,
	kTwoOfClubs,
	kThreeOfClubs,
	kFourOfClubs,
	kFiveOfClubs,
	kSixOfClubs,
	kSevenOfClubs,
	kEightOfClubs,
	kNineOfClubs,
	kTenOfClubs,
	kJackOfClubs,
	kQueenOfClubs,
	kKingOfClubs,
	kJoker,
}

func allCardsOfSuit(suit Suit) []Card {
	cards := make([]Card, 0, 13)
	for _, c := range kAllCards {
		if c.suit == suit {
			cards = append(cards, c)
		}
	}
	return cards
}

var kAceOfSpades = Card{
	rank: RankAce,
	suit: SuitSpades,
	str:  "🂡",
}

var kAceOfHearts = Card{
	rank: RankAce,
	suit: SuitHearts,
	str:  "🂱",
}

var kAceOfDiamonds = Card{
	rank: RankAce,
	suit: SuitDiamonds,
	str:  "🃁",
}

var kAceOfClubs = Card{
	rank: RankAce,
	suit: SuiteClubs,
	str:  "🃑",
}

var kTwoOfSpades = Card{
	rank: RankTwo,
	suit: SuitSpades,
	str:  "🂢",
}

var kTwoOfHearts = Card{
	rank: RankTwo,
	suit: SuitHearts,
	str:  "🂲",
}

var kTwoOfDiamonds = Card{
	rank: RankTwo,
	suit: SuitDiamonds,
	str:  "🃂",
}

var kTwoOfClubs = Card{
	rank: RankTwo,
	suit: SuiteClubs,
	str:  "🃒",
}

var kThreeOfSpades = Card{
	rank: RankThree,
	suit: SuitSpades,
	str:  "🂣",
}

var kThreeOfHearts = Card{
	rank: RankThree,
	suit: SuitHearts,
	str:  "🂳",
}

var kThreeOfDiamonds = Card{
	rank: RankThree,
	suit: SuitDiamonds,
	str:  "🃃",
}

var kThreeOfClubs = Card{
	rank: RankThree,
	suit: SuiteClubs,
	str:  "🃓",
}

var kFourOfSpades = Card{
	rank: RankFour,
	suit: SuitSpades,
	str:  "🂤",
}

var kFourOfHearts = Card{
	rank: RankFour,
	suit: SuitHearts,
	str:  "🂴",
}

var kFourOfDiamonds = Card{
	rank: RankFour,
	suit: SuitDiamonds,
	str:  "🃄",
}

var kFourOfClubs = Card{
	rank: RankFour,
	suit: SuiteClubs,
	str:  "🃔",
}

var kFiveOfSpades = Card{
	rank: RankFive,
	suit: SuitSpades,
	str:  "🂥",
}

var kFiveOfHearts = Card{
	rank: RankFive,
	suit: SuitHearts,
	str:  "🂵",
}

var kFiveOfDiamonds = Card{
	rank: RankFive,
	suit: SuitDiamonds,
	str:  "🃅",
}

var kFiveOfClubs = Card{
	rank: RankFive,
	suit: SuiteClubs,
	str:  "🃕",
}

var kSixOfSpades = Card{
	rank: RankSix,
	suit: SuitSpades,
	str:  "🂦",
}

var kSixOfHearts = Card{
	rank: RankSix,
	suit: SuitHearts,
	str:  "🂶",
}

var kSixOfDiamonds = Card{
	rank: RankSix,
	suit: SuitDiamonds,
	str:  "🃆",
}

var kSixOfClubs = Card{
	rank: RankSix,
	suit: SuiteClubs,
	str:  "🃖",
}

var kSevenOfSpades = Card{
	rank: RankSeven,
	suit: SuitSpades,
	str:  "🂧",
}

var kSevenOfHearts = Card{
	rank: RankSeven,
	suit: SuitHearts,
	str:  "🂷",
}

var kSevenOfDiamonds = Card{
	rank: RankSeven,
	suit: SuitDiamonds,
	str:  "🃇",
}

var kSevenOfClubs = Card{
	rank: RankSeven,
	suit: SuiteClubs,
	str:  "🃗",
}

var kEightOfSpades = Card{
	rank: RankEight,
	suit: SuitSpades,
	str:  "🂨",
}

var kEightOfHearts = Card{
	rank: RankEight,
	suit: SuitHearts,
	str:  "🂸",
}

var kEightOfDiamonds = Card{
	rank: RankEight,
	suit: SuitDiamonds,
	str:  "🃈",
}

var kEightOfClubs = Card{
	rank: RankEight,
	suit: SuiteClubs,
	str:  "🃘",
}

var kNineOfSpades = Card{
	rank: RankNine,
	suit: SuitSpades,
	str:  "🂩",
}

var kNineOfHearts = Card{
	rank: RankNine,
	suit: SuitHearts,
	str:  "🂹",
}

var kNineOfDiamonds = Card{
	rank: RankNine,
	suit: SuitDiamonds,
	str:  "🃉",
}

var kNineOfClubs = Card{
	rank: RankNine,
	suit: SuiteClubs,
	str:  "🃙",
}

var kTenOfSpades = Card{
	rank: RankTen,
	suit: SuitSpades,
	str:  "🂪",
}

var kTenOfHearts = Card{
	rank: RankTen,
	suit: SuitHearts,
	str:  "🂺",
}

var kTenOfDiamonds = Card{
	rank: RankTen,
	suit: SuitDiamonds,
	str:  "🃊",
}

var kTenOfClubs = Card{
	rank: RankTen,
	suit: SuiteClubs,
	str:  "🃚",
}

var kJackOfSpades = Card{
	rank: RankJack,
	suit: SuitSpades,
	str:  "🂫",
}

var kJackOfHearts = Card{
	rank: RankJack,
	suit: SuitHearts,
	str:  "🂻",
}

var kJackOfDiamonds = Card{
	rank: RankJack,
	suit: SuitDiamonds,
	str:  "🃋",
}

var kJackOfClubs = Card{
	rank: RankJack,
	suit: SuiteClubs,
	str:  "🃛",
}

var kQueenOfSpades = Card{
	rank: RankQueen,
	suit: SuitSpades,
	str:  "🂭",
}

var kQueenOfHearts = Card{
	rank: RankQueen,
	suit: SuitHearts,
	str:  "🂽",
}

var kQueenOfDiamonds = Card{
	rank: RankQueen,
	suit: SuitDiamonds,
	str:  "🃍",
}

var kQueenOfClubs = Card{
	rank: RankQueen,
	suit: SuiteClubs,
	str:  "🃝",
}

var kKingOfSpades = Card{
	rank: RankKing,
	suit: SuitSpades,
	str:  "🂮",
}

var kKingOfHearts = Card{
	rank: RankKing,
	suit: SuitHearts,
	str:  "🂾",
}

var kKingOfDiamonds = Card{
	rank: RankKing,
	suit: SuitDiamonds,
	str:  "🃎",
}

var kKingOfClubs = Card{
	rank: RankKing,
	suit: SuiteClubs,
	str:  "🃞",
}

var kJoker = Card{
	rank: RankJoker,
	suit: SuitNone,
	str:  "🃟",
}
