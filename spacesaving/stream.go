package spacesaving

type StreamManager interface {
	Offer(item string, increment int)
	Top() []Result
	Name() string
}
