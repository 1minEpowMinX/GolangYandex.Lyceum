package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	swapMap := make(map[string]string)
	for k, v := range m {
		swapMap[v] = k
	}

	return swapMap
}
