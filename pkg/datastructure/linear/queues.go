package linear


type Queue []*Order

type Order struct {
	priority int
	quantity int
	product string
	customerName string
}
// New method
func (order *Order) New(priority,quantity int,product,customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}
// Add method
func (queue *Queue) Add(order *Order) {
	if len(*queue)==0{
		*queue = append(*queue, order)
	}else {
		appended := false
		for i,addedOrder := range *queue{
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue,order)
		}
	}
}
