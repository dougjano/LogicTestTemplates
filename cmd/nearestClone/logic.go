package main

func CreateShorterArray(ids []int, colorId int) []int {

	result := []int{}

	for i, v := range ids {
		if colorId == v {
			result = append(result, i+1)
		}
	}

	return result
}

type Metadata struct {
	source   int
	distance int
}

func RecursiveExpansion(queue []int, visited map[int]Metadata, adj map[int][]int) int {
	if len(queue) == 0 {
		return -1
	}

	nextQueue := []int{}

	for _, u := range queue {
		for _, v := range adj[u] {
			metaU := visited[u]

			if metaV, exists := visited[v]; exists {
				if metaV.source != metaU.source {
					return metaU.distance + metaV.distance + 1
				}
				continue
			}

			visited[v] = Metadata{source: metaU.source, distance: metaU.distance + 1}
			nextQueue = append(nextQueue, v)
		}
	}

	return RecursiveExpansion(nextQueue, visited, adj)
}

func FindNeighbours(colorNode int, graphFrom []int, graphTo []int, ids []int) int {
	adj := make(map[int][]int)
	for i := range graphFrom {
		u, v := graphFrom[i], graphTo[i]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	startNodes := []int{}
	visited := make(map[int]Metadata)
	for i, color := range ids {
		if color == colorNode {
			nodeID := i + 1
			startNodes = append(startNodes, nodeID)
			visited[nodeID] = Metadata{source: nodeID, distance: 0}
		}
	}

	if len(startNodes) < 2 {
		return -1
	}

	return RecursiveExpansion(startNodes, visited, adj)
}
