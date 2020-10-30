package graph

import (
	"ajiiranetservice/utils"
	"ajiiranetservice/vo"
)

// AjiiraGraph :
type AjiiraGraph struct {
	Connections      map[string][]string         `json:"connections,omitempty"`
	ConnectedDevices map[string]*vo.AJIIRADevice `json:"connectedDevices,omitempty"`
}

// AddEdge :
func (ajiiraGraph *AjiiraGraph) AddEdge(v, u string) {

	if val, ok := ajiiraGraph.Connections[v]; ok {

		if !utils.IsSliceContains(val, u) {
			val = append(val, u)
			ajiiraGraph.Connections[v] = val
		}

	} else {
		if ajiiraGraph.Connections == nil {
			ajiiraGraph.Connections = make(map[string][]string, 0)
		}
		ajiiraGraph.Connections[v] = []string{u}
	}
}

// AddDevice :
func (ajiiraGraph *AjiiraGraph) AddDevice(name, typee string, strength int) {

	ajiiraGraph.ConnectedDevices[name] = &vo.AJIIRADevice{Name: name, Type: typee, Strength: &vo.AJIIRADeviceStrength{Value: strength}}

}

// DepthFirst :
func (ajiiraGraph *AjiiraGraph) DepthFirst(visited []string, end string, paths [][]string) [][]string {
	if len(visited) == 0 {
		return paths
	}
	visitNode := visited[len(visited)-1]
	if nodes, ok := ajiiraGraph.Connections[visitNode]; ok {
		for _, node := range nodes {
			if utils.IsSliceContains(visited, node) {
				continue
			}

			if node == end {
				visited = append(visited, node)
				paths = append(paths, visited)
				// printPath(visited)
				visited = visited[:len(visited)-1]
				return paths
			}
		}

		for _, node := range nodes {
			if utils.IsSliceContains(visited, node) || node == end {
				continue

			}
			visited = append(visited, node)
			paths = ajiiraGraph.DepthFirst(visited, end, paths)
			visited = visited[:len(visited)-1]

		}

	} else {

		if len(visited) > 1 {
			visited = visited[:len(visited)-1]
		}
		for key, values := range ajiiraGraph.Connections {
			if utils.IsSliceContains(values, visited[len(visited)-1]) && !utils.IsSliceContains(visited, key) {

				visited = append(visited, key)
				for key, values := range ajiiraGraph.Connections {
					if utils.IsSliceContains(values, visited[len(visited)-1]) && !utils.IsSliceContains(visited, key) {
						visited = append(visited, key)
						paths = ajiiraGraph.DepthFirst(visited, end, paths)
						if len(paths) == 0 {
							visited = visited[:len(visited)-1]
						}
					}
				}

				paths = ajiiraGraph.DepthFirst(visited, end, paths)
				if len(paths) == 0 {
					visited = visited[:len(visited)-1]
				}

			}

		}

	}

	return paths

}

func (ajiiraGraph *AjiiraGraph) isConnected(v, u string) bool {
	if val, ok := ajiiraGraph.Connections[v]; ok {
		if utils.IsSliceContains(val, u) {
			return true
		}
	}
	return false
}
