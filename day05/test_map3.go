package main

func main() {
	type user struct{ name string }
	m := map[int]user{ // 当 map 因扩张⽽而重新哈希时，各键值项存储位置都会发⽣生改变。 因此，map
		1: {"user1"}, // 被设计成 not addressable。 类似 m[1].name 这种期望透过原 value
	} // 指针修改成员的行为自然会被禁⽌。
	//	m[1].name = "Tom" // Error: cannot assign to m[1].name

	u := m[1]
	u.name = "Tom"
	m[1] = u // 替换 value。

	m2 := map[int]*user{
		1: &user{"user1"},
	}
	m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。

	println(m[1].name)  // Tom
	println(m2[1].name) // Jack
}
