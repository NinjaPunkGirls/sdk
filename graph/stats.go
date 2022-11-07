package graph

// keeps track of total predicate content counts
type PredicateStat struct {
	Key   string `firestore: "key,omitempty"`
	Value int    `firestore: "value"`
}
