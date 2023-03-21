// how to write switch

// regular switch
// 每个都是 == 的时候用
switch size := len(word); size {
case 1:
	// do something
case 2, 3, 4:
	// do something
default:
	// do something
}


// blank switch
// 需要不等式的时候用
switch size := len(word);  {
case size < 4:
	// do something
case size < 6 && size >= 4:
	// do something
default:
	// do something
}

// 某case后加 fallthrough 后可以接着判断

// if 与 switch 的选择
// 当有一侧的变量都一样时，偏向于用 switch
// 当总是不一样时，偏向于用 if