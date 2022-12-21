package encapsulation

type Student struct {
	Name  string
	score int // 工厂模式 小写,隐私属性,不能随便操作,只能通过特定的方法
}

// 同一个包
// 封装一个new方法,用于实例化结构体
// 使用它创建的结构体的可以打印,但小写资源需要通过特定的方法
func NewStudent(name string, score int) *Student {
	stu := &Student{
		// 不一定要全部成员
		Name:  name,
		score: score, //同一个包下,所以可以使用score:
	}
	return stu
}

//获取分数
func (s *Student) GetScore() int {
	return s.score
}

//设置分数
func (s *Student) SetScore(score int) {
	s.score = score
}
