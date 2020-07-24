# Go-Dom

A very simple implementation of the DOM (document object model) in which we compare/benchmark node look-up strategies.\
The implemented node look-up strategies are based on the DOM function - `document.getElementById(...)`.

The following node look-up strategies are implemented:

- Breadth first search
- Depth first search
- Parallel search (custom Go implementation)

The goal of this benchmark is to identify how efficient we can get when using parallel/concurrent searching technique as compared to traditional BFS & DFS searching strategies.

...

## Resources

The project was inspired by the following [blog](https://ieftimov.com/post/golang-datastructures-trees/)

---

- what this is
- inspiration from blog
- data structure implementations
- non-recursive strategies
- parallel concurrency based search implementation
- tests
- benchmarks
