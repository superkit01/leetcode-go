package middle

type SnapshotArray struct {
	list          map[int][]*SnapshotNode
	CurrentSnapId int
}

type SnapshotNode struct {
	SnapId int
	Val    int
}

func Constructor1146(length int) SnapshotArray {
	list := make(map[int][]*SnapshotNode, length)
	for i := 0; i < length; i++ {
		list[i] = []*SnapshotNode{
			&SnapshotNode{
				SnapId: 0,
				Val:    0,
			},
		}
	}

	return SnapshotArray{
		list:          list,
		CurrentSnapId: 0,
	}

}

func (this *SnapshotArray) Set(index int, val int) {
	nodes := this.list[index]

	node := nodes[len(nodes)-1]
	if node.SnapId == this.CurrentSnapId {
		node.Val = val
	} else {
		this.list[index] = append(this.list[index],
			&SnapshotNode{
				SnapId: this.CurrentSnapId,
				Val:    val})
	}
}

func (this *SnapshotArray) Snap() int {
	ans := this.CurrentSnapId
	this.CurrentSnapId++
	return ans
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	nodes := this.list[index]

	if len(nodes) == 1 {
		return nodes[0].Val
	}

	//二分  <=target 的最大值
	l := 0
	r := len(nodes) 
	for l < r {
		mid := (l + r) / 2
		if nodes[mid].SnapId <= snap_id {
			l = mid + 1
		}
		if nodes[mid].SnapId > snap_id {
			r = mid
		}
	}
	return nodes[l-1].Val

}
