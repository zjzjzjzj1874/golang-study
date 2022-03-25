> ## golang-study
golang study

> ## doc
> some documents with plantuml

> ## src
> my golang source code
> ### base
> golang base syntax and struct
> ### extension
> extension provide some useful package and data struct
> ### helper
> helper provide some tools,and some can use in project

> ## golang中的坑

- **slice**
  - `slice与nil`比较
      - `slice == nil`的比较:本质上是与`slice.Data(uintptr)`比较,所以需要看是否被初始化
      - 未初始化的情况:与nil比较返回true
          - `var s []int`
      - 已初始化:与nil比较返回false
          - `s = make([]int,0)`
          - `s3 := []int{1, 2, 3, 4, 5, 6}   s4 := s3[:3]`
      - 特殊情况
          - `s := new([]int)`,这个需要注意,new返回的是指向`[]int`类型的指针,所以`s`是一个地址,`*s`是一个未初始化的`[]int`.
  - `slice`元素修改
    - golang是值传递,所以不管是append还是modify:一看是否能指回原来切片,二看是否会生成新切片
  - `slice`元素扩容原理
    - `old.cap < 1024` => `newcap = doublecap`
    - `old.cap >= 1024` => 先确定大致容量`newcap += newcap/4` => 再根据切片元素大小对齐内存,然后向上取整
> ## TODO list

- [ ] base:待实现:error方法传入context,追踪调用链路
- [ ] github.com/pkg/errors:学习warp error包装
- [ ] select的使用场景:超时控制,并发channel..