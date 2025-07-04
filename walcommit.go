package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Command struct {
	Key   string
	Value string
}

type WAL struct {
	mu   sync.Mutex
	file *os.File
}

func NewWAL(filename string) *WAL {
	f, _ := os.Create(filename)
	return &WAL{file: f}
}

func (w *WAL) Write(cmd Command) {
	w.mu.Lock()
	defer w.mu.Unlock()
	fmt.Fprintf(w.file, "%s=%s\n", cmd.Key, cmd.Value)
	w.file.Sync()
}

type Node struct {
	id        int
	log       []Command
	committed []Command
	mu        sync.Mutex
}

func NewNode(id int) *Node {
	return &Node{id: id}
}

func (n *Node) Replicate(cmd Command, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	n.mu.Lock()
	n.log = append(n.log, cmd)
	n.mu.Unlock()
}

func main() {
	nodes := []*Node{NewNode(1), NewNode(2), NewNode(3)}
	wal := NewWAL("wal.log")
	cmd := Command{Key: "x", Value: "42"}
	start := time.Now()
	wal.Write(cmd)
	fmt.Println("Client ACK at:", time.Since(start).Milliseconds(), "ms")
	var wg sync.WaitGroup
	for _, n := range nodes {
		wg.Add(1)
		go n.Replicate(cmd, 20*time.Millisecond, &wg)
	}
	wg.Wait()
	for _, n := range nodes {
		n.mu.Lock()
		n.committed = append(n.committed, cmd)
		n.mu.Unlock()
	}
	fmt.Println("Committed at:", time.Since(start).Milliseconds(), "ms")
}
