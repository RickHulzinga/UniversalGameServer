package UniversalGameServer

func findConnectionIndex(x *Client) int {
	for i, n := range connections {
		if x == n {
			return i
		}
	}
	return len(connections)
}

func removeConnection(i int) {
	connections = append(connections[:i], connections[i+1:]...)
}
