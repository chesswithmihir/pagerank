# Design Doc: PageRank Implementation

## 0. Overview

PageRank is the algorithm developed by Larry Page and Sergey Brin at Stanford in an attempt to build their first search engine, BackRub, now known today as Google... BackRub would have been a terrible name lol.

This algorithm was made using Linear Algebra and Graph Theory.

## 1. Overview & Theory

**The Goal:** Calculate the "importance" of nodes in a graph based on incoming links.

* **Technically:** It is a **Discrete-Time Markov Chain**. I initially thought it was an MDP, but an MDP usually implies an agent making decisions to maximize a reward. Here, there is no "decision." There is only a "Random Surfer" moving probabilistically.

* **The computation:** We are solving for the **Principal Eigenvector** of the web's transition matrix. We want to find the stationary distribution (where the surfer spends the most time).

**Do we need a GPU?**

I asked cuz that would be cool! Go NVIDIA! but i don't think it's needed. Confirmed by Gemini (BackRub --> Google --> Gemini) lol. Do people say *Lemme Gemini that...* i dont think so :(

But then again. they prob didn't say, let's BackRub that, right? idk it was 1996 in some garage in Menlo Park, CA. Don't ask me.

Anyways, CPU is fine for millions of nodes. We can rewrite it to include CUDA or Metal (I have a Mac), but let's probably.

---

## 2. High-Level Architecture

We will design this as a **Stream Processing Engine**. We assume the graph might be too big to fit fully into memory as a standard adjacency matrix, so we process it in chunks (or use sparse formats).

### Components:

1. **The Ingestor:** Parses raw data (e.g., CSV, Adjacency List, or a Web Crawl) into a standardized format.
2. **The Graph Indexer:** Maps string URLs to 64-bit Integers (Node IDs). You cannot do math on strings.
3. **The Engine:** Performs the Power Iteration algorithm.
4. **The Convergence Monitor:** Checks if the ranks have stabilized (difference between iterations < ).

---

## 3. Directory Structure

Let's build it in go!

```text
/pagerank
├── /data                # Place your toy graphs here (web-Stanford.txt, etc.)
├── /docs                # Your diagrams and notes
├── /src
│   ├── /graph
│   │   ├── loader.go    # File reading logic
│   │   └── mapper.go    # URL <-> ID Bi-directional mapping
│   ├── /math
│   │   ├── matrix.go    # Sparse Matrix data structures (CSR/CSC)
│   │   └── vector.go    # Vector operations (dot product, normalization)
│   ├── /engine
│   │   ├── pagerank.go  # The core iteration logic
│   │   └── dumper.go    # Writing results to disk
│   └── /utils
│       └── config.go    # Constants (Damping factor, Epsilon)
├── /tests               # Unit and Integration tests
├── Makefilel
└── README.md

```

---

## 4. Data Structures & Algorithms

### A. Data Structures

This is where the magic happens. You cannot store a matrix for the web; it's mostly zeros.

1. **The Dictionary (Hash Map):**
* Key: `String` (URL)
* Value: `uint64` (Node ID)
* *Challenge:* This can get huge. Efficient string hashing is key.


2. **Compressed Sparse Row (CSR) Format (representing the Reverse Graph):**
* To enable lock-free parallel writing to the `next_ranks` vector, we need to know who points *to* a node (Incoming Links).
* Therefore, our CSR structure will effectively store the **Transpose** of the web graph:
    * `values[]`: The weights (usually $1.0 / \text{out\_degree}(\text{source})$).
    * `column_indices[]`: The **Source** node IDs (who links *to* the current row).
    * `row_ptr[]`: Index range for a node's **Incoming** links.
* *Why:* This allows Thread $T$ to own nodes $0..N$, iterate their incoming neighbors, and write results safely without mutexes.


3. **The Rank Vectors (Double Buffering):**
* `current_ranks[]`: Array of floats (size ).
* `next_ranks[]`: Array of floats (size ).
* *Why two?* To avoid **Race Conditions**. You read from `current` and write to `next`.



### B. Algorithms

1. **Power Iteration:**
* The core loop. . Repeat until .


2. **Teleportation Handling (Damping Factor):**
* The "Random Surfer" gets bored and jumps to a random page.
* Formula:
* Usually .


3. **Dead-End Handling:**
* Nodes with no outgoing links (Dangling nodes) absorb rank and destroy the math (sum of ranks < 1).
* *Algo fix:* Redistribute their rank evenly to all other nodes.



---

## 5. Multithreading & Concurrency

**Will we face race conditions?**

* **Yes**, if you try to update a single vector in place using multiple threads.
* **No**, if you use the **Read/Write Split** strategy.

### The Strategy: Parallel Vector Multiplication

Since matrix multiplication is a series of dot products, it is "embarrassingly parallel."

**Pseudocode Logic for Threading:**

1. Divide the `row_ptr` array into chunks based on the number of threads (e.g., if , Thread 1 does rows 0-25, Thread 2 does 26-50...).
2. Spawning threads:
* Input: (Shared pointer to Graph, Shared pointer to `current_ranks`, Mutable pointer to slice of `next_ranks`).


3. **Barrier:** Wait for all threads to finish calculating `next_ranks`.
4. **Swap:** Pointer swap `current_ranks` and `next_ranks`.
5. Repeat.

*Note: You do not need Mutexes/Locks for the calculation phase if threads write to disjoint slices of memory!*

---

## 6. Implementation & Testing Guide

Since you want to write the code, here are the function signatures you need to implement.

### File: `src/graph/mapper`

```text
class NodeMapper:
    function get_id(url: string) -> int
    function get_url(id: int) -> string
    function size() -> int

```

### File: `src/engine/pagerank`

```text
function calculate_dangling_weight(graph, current_ranks) -> float
// Sums up rank from pages with no outbound links to redistribute

function step(graph, current_ranks, damping_factor) -> next_ranks
// The heavy lifter.
// 1. Launch threads.
// 2. Each thread calculates incoming rank for its chunk of nodes.
// 3. Add (1-d) teleportation constant.
// 4. Add redistributed dangling weight.

```

### Testing (Crucial!)

How do you know it works?

1. **The Sum Test:** The sum of all values in the PageRank vector **must always equal 1.0** (or the number of nodes, depending on normalization). If it drops to 0.999... after 10 iterations, you have a leak (likely dangling nodes).
2. **The Tiny Web Test:** Hand-draw a 3-node graph. Calculate the ranks on paper. Run your code. Do they match?
3. **Convergence Test:** Ensure the error (delta) decreases with every iteration.

## 7. Next Steps

1. **Download a dataset:** Search for "Stanford SNAP Datasets" and grab the `web-Google` or `web-Stanford` text file.
2. **Draft your struct:** Define exactly how you will hold the graph in memory.
3. **Start Coding:** Implement the `Loader` first.

Do you want me to expand on the **CSR (Compressed Sparse Row)** logic? It is the trickiest part of the data structure design but essential for performance.
