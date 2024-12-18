package day09

func Run(data []string) [2]interface{} {
	nums := make([]int, len(data[0]))
	for i, l := range data[0] {
		nums[i] = int(l - '0')
	}

	s := Solver{}

	return [2]interface{}{
		s.CalculateSingle(nums),
		0,
	}
}

type Solver struct {
	Files, Free, Nums []int
}

func (s *Solver) Init(nums []int) {
	s.Files = make([]int, 0)
	s.Free = make([]int, 0)
	s.Nums = nums
	for i := range nums {
		if i%2 == 0 {
			s.Files = append(s.Files, nums[i])
		} else {
			s.Free = append(s.Free, nums[i])
		}
	}
}

func (s *Solver) TakeFromBack() int {
	for i := len(s.Nums) - 1; i >= 0; i -= 2 {
		if s.Nums[i] == 0 {
			continue
		}

		s.Nums[i]--
		return i / 2
	}

	return -1
}

func (s *Solver) CalculateSingle(nums []int) int {
	s.Init(nums)
	var total int
	var pos int

	for i := range s.Nums {
		if i%2 == 0 {
			for j := 0; j < nums[i]; j++ {
				total += pos * (i / 2)
				pos++
			}
			s.Nums[i] = 0
			continue
		}

		for j := 0; j < nums[i]; j++ {
			n := s.TakeFromBack()

			if n < 0 {
				return total
			}

			total += pos * (n)
			pos++
		}
	}

	return total
}
