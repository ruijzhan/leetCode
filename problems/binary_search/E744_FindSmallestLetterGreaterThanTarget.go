package binary_search

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low < high {
		mid := low + (high-low)>>1
		if letters[mid] == target {
			if mid < len(letters)-1 {
				if letters[mid+1] == target {
					low++
					continue
				}
				return letters[mid+1]
			} else {
				return letters[0]
			}
		}
		if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if letters[low] > target {
		return letters[low]
	} else {
		if low < len(letters)-1 {
			return letters[low+1]
		} else {
			return letters[0]
		}
	}
}
