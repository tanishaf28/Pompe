package main

import (
	"fmt"
	"sort"
	"time"
	"math/rand"
)

// Replica represents a node in the network
type Replica struct {
	ID     int
	Honest bool // Determines if the replica is Byzantine or honest
}

// Transaction represents a client request
type Transaction struct {
	ClientID string
	Command  string
}

// TimestampProposal represents a timestamp proposal
type TimestampProposal struct {
	ReplicaID int
	Timestamp int64
}

// Byzantine Ordered Consensus Algorithm
func ByzantineOrderedConsensus(replicas []Replica, tx Transaction) int64 {
	fmt.Printf("\n 🔹New Transaction from %s: %s\n", tx.ClientID, tx.Command)
	proposals := []TimestampProposal{}

	// Step 1: Each replica proposes a timestamp
	fmt.Println("Step 1: Replicas propose timestamps...")
	for _, replica := range replicas {
		var ts int64
		if replica.Honest {
			// Honest replicas use correct timestamps
			ts = time.Now().UnixNano()
			fmt.Printf("Replica %d (Honest) proposes timestamp: %d\n", replica.ID, ts)
		} else {
			// Byzantine replicas provide random timestamps
			ts = rand.Int63n(time.Now().UnixNano())
			fmt.Printf("Replica %d (Byzantine) proposes timestamp: %d\n", replica.ID, ts)
		}

		proposals = append(proposals, TimestampProposal{ReplicaID: replica.ID, Timestamp: ts})
	}

	// Step 2: Sort timestamps
	fmt.Println("\nStep 2: Sorting timestamps...")
	sort.Slice(proposals, func(i, j int) bool {
		return proposals[i].Timestamp < proposals[j].Timestamp
	})
	for _, p := range proposals {
		fmt.Printf("   - Replica %d proposed timestamp: %d\n", p.ReplicaID, p.Timestamp)
	}

	// Step 3: Select the median timestamp
	medianIndex := len(proposals) / 2
	orderedTimestamp := proposals[medianIndex].Timestamp
	fmt.Printf("\nStep 3: Median timestamp selected: %d\n", orderedTimestamp)

	fmt.Printf("🔸Final Order for '%s' -> %d\n", tx.Command, orderedTimestamp)
	return orderedTimestamp
}

func main() {
	// Simulating a 3f+1 system with f = 1 (total 4 replicas, 3 honest, 1 Byzantine)
	replicas := []Replica{
		{ID: 1, Honest: true},
		{ID: 2, Honest: true},
		{ID: 3, Honest: true},
		{ID: 4, Honest: false}, // Byzantine replica
	}

	// Sample transactions from Alice and Bob
	aliceTx := Transaction{ClientID: "Alice", Command: "Transfer 10 BTC"}
	bobTx := Transaction{ClientID: "Bob", Command: "Transfer 5 BTC"}

	// Print Initial Transactions
	fmt.Println("Initial Transactions:")
	fmt.Printf("🔺Alice: %s\n", aliceTx.Command)
	fmt.Printf("🔺Bob: %s\n", bobTx.Command)

	// Run Byzantine Ordered Consensus for both transactions
	aliceOrder := ByzantineOrderedConsensus(replicas, aliceTx)
	bobOrder := ByzantineOrderedConsensus(replicas, bobTx)

	// Determine who goes first
	fmt.Println("\n🔹Comparing Orders:")
	if aliceOrder < bobOrder {
		fmt.Println("Alice's transaction is processed before Bob's.")
	} else {
		fmt.Println("Bob's transaction is processed before Alice's.")
	}
}

