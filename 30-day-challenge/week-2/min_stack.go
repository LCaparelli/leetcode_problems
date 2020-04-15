/*
 * Min Stack
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
 *
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 *
 *
 *
 * Example:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 */

package challenge

type element struct {
	value int
	min   int
}

type MinStack struct {
	values []element
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]element{}}
}

func (m *MinStack) Push(x int) {
	if len(m.values) == 0 {
		m.values = append(m.values, element{x, x})
	} else {
		curMin := m.values[len(m.values)-1].min
		m.values = append(m.values, element{x, min(x, curMin)})
	}
}

func (m *MinStack) Pop() {
	if len(m.values) > 0 {
		m.values = m.values[:len(m.values)-1]
	}
}

func (m *MinStack) Top() int {
	return m.values[len(m.values)-1].value
}

func (m *MinStack) GetMin() int {
	return m.values[len(m.values)-1].min
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
