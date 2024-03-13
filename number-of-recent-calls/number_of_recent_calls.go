package numberofrecentcalls

type RecentCounter struct {
	reqs []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.reqs = append(r.reqs, t)
	r.removeOutdated(t)
	return len(r.reqs)
}

func (r *RecentCounter) removeOutdated(t int) {
	if len(r.reqs) == 0 || r.reqs[0] >= t-3000 {
		return
	}
	r.reqs = r.reqs[1:]
	r.removeOutdated(t)
}
