// 线段树

int N = (int)1e9; // N for MAX
// ls for left_son
// rs for right_son
// add for value need to pushdown
// val for value of this node
int ls = 0, rs = 1, add = 2, val = 3, cnt = 1, cur_max = 0;
// 准备好 120010个结点，具体数量看题目
// 使用结点前需要 全部置 0
int tr[120010][4];

// 更新 tr[u][val]
void pushup(int u) {
	// 从两个子结点得出u结点的值
	// 具体计算方式need
	tr[u][val] = max(tr[tr[u][ls]][val], tr[tr[u][rs]][val]);
}

// 把 tr[u] 的 add 对两个子树进行下传
void pushdown(int u, int len) {
	int cur_add = tr[u][add];
	if(cur_add == 0) return;
	int l_son = tr[u][ls], r_son = tr[u][rs];
	// 具体计算方法need
	tr[l_son][add] += cur_add;
	tr[r_son][add] += cur_add;
	tr[l_son][val] += cur_add;
	tr[r_son][val] += cur_add;

	tr[u][add] = 0;
}

// 在代表 [lc, rc) 范围的结点 u 更新 [l, r] v
void update(int u, int lc, int rc, int l, int r, int v) {
	if(u == 0)
		cout<<"error"<<endl;
	if(l <= lc && rc <= r) {
		// 具体计算方式need
		tr[u][val] += v;
		tr[u][add] += v;
		return;
	}
	lazyCreate(u);
	// 第二个参数，具体计算方式need
	pushdown(u, rc-lc+1);
	int mid = (lc + rc) >> 1;
	if(l <= mid) update(tr[u][ls], lc, mid, l, r, v);
	if (r > mid) update(tr[u][rs], mid + 1, rc, l, r, v);
	pushup(u);
}

int query(int u, int lc, int rc, int l, int r) {
	// 如果此结点代表的范围在 [l, r] 中，直接返回val
	if (l <= lc && rc <= r) return tr[u][val];
	// 不在的话就造两个子结点
	lazyCreate(u);
	// 把当前结点下传
	pushdown(u, rc - lc + 1);
	int mid = (lc + rc) >> 1, ans = 0;
	// 具体计算方式need
    if (l <= mid)
    	ans = query(tr[u][ls], lc, mid, l, r);
    if (r > mid) ans = max(query(tr[u][rs], mid + 1, rc, l, r), ans);
    return ans;
}

// 确保 tr[u]左右子树都有所指向
void lazyCreate(int u) {
	if(tr[u][ls] == 0) {
		tr[u][ls] = ++cnt;
	}
	if(tr[u][rs] == 0) {
		tr[u][rs] = ++cnt;
	}
}