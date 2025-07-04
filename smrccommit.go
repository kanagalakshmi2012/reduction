package main
import (
	"fmt"
	"math"
	"sync"
	"time"
)
type Command struct {
	Key   string
	Value string
}
type LogEntry struct {
	Index   int
	Command Command
}
type StateMachine struct {
	mu    sync.Mutex
	state map[string]string
}
func NewStateMachine() *StateMachine {
	return &StateMachine{state: make(map[string]string)}
}
func (sm *StateMachine) Apply(cmd Command) {
	sm.mu.Lock()
	sm.state[cmd.Key] = cmd.Value
	sm.mu.Unlock()
}
func (sm *StateMachine) Snapshot() map[string]string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	copy := make(map[string]string)
	for k, v := range sm.state {
		copy[k] = v
	}
	return copy
}
type Node struct {
	id        int
	log       []LogEntry
	commitIdx int
	sm        *StateMachine
}
func NewNode(id int) *Node {
	return &Node{
		id:  id,
		sm:  NewStateMachine(),
		log: make([]LogEntry, 0),
	}
}
func (n *Node) Append(entry LogEntry) {
	n.log = append(n.log, entry)
}
func (n *Node) CommitUpTo(idx int) {
	for n.commitIdx <= idx && n.commitIdx < len(n.log) {
		entry := n.log[n.commitIdx]
		n.sm.Apply(entry.Command)
		n.commitIdx++
	}
}
func simulateReplication(nodes []*Node, cmd Command) {
	leader := nodes[0]
	entry := LogEntry{
		Index:   len(leader.log),
		Command: cmd,
	}
	leader.Append(entry)
	var wg sync.WaitGroup
	var mu sync.Mutex
	acks := 1
	for _, follower := range nodes[1:] {
		wg.Add(1)
		go func(f *Node) {
			defer wg.Done()
			latency := simulateLatency(len(nodes))
			time.Sleep(latency)
			f.Append(entry)
			mu.Lock()
			acks++
			mu.Unlock()
		}(follower)
	}
	wg.Wait()
	majority := int(math.Floor(float64(len(nodes))/2.0)) + 1
	if acks >= majority {
		for _, node := range nodes {
			node.CommitUpTo(entry.Index)
		}
	}
}
func simulateLatency(n int) time.Duration {
	switch n {
	case 3:
		return 14 * time.Millisecond
	case 5:
		return 18 * time.Millisecond
	case 7:
		return 22 * time.Millisecond
	case 9:
		return 26 * time.Millisecond
	case 11:
		return 30 * time.Millisecond
	default:
		return time.Duration(10+2*n) * time.Millisecond
	}
}

func main() {
	nodeCounts := []int{3, 5, 7, 9, 11}
	for _, count := range nodeCounts {
		nodes := make([]*Node, count)
		for i := 0; i < count; i++ {
			nodes[i] = NewNode(i + 1)
		}
		start := time.Now()
		simulateReplication(nodes, Command{Key: "k", Value: fmt.Sprintf("v%d", count)})
		duration := time.Since(start)
		fmt.Printf("Nodes: %d\tSMR Commit Latency: %.1f ms\n", count, float64(duration.Milliseconds()))
	}
}
