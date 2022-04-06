package menu

 
 type LinkTableNode struct {
	data interface{}
	pNext *LinkTableNode
 }
 
 
 type LinkTable struct{
	pHead *LinkTableNode
	pTail *LinkTableNode
	SumOfNode int			
	mutex sync.RWMutex
 }
 
 func NewLinkTableNode(value interface{})*LinkTableNode{
	return &LinkTableNode{
		data:value,
		pNext:nil,
	}
}

 func  CreateLinkTable() *LinkTable{
	return  &LinkTable{
		root.pHead:  nil,
		root.pTail:  nil,
		root.SumOfNode:0,
	}
	
 }

 func IsEmpty(pLinkTable*LinkTable) bool{
	return pLinkTable.SumOfNode==0
 }

 func DeleteLinkTable(pLinkTable*LinkTable) int{
	if pLinkTable == nil {
		return false
	}
	for !IsEmpty(pLinkTable) {
		var p *LinkTableNode = pLinkTable.pHead
		pLinkTable.pHead = pLinkTable.pHead.pNext
		p.pNext = nil
		p.data = nil
		pLinkTable.SumOfNode--
	}
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	return true
 }

 func AddLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) int{
	if pLinkTable == nil || pNode == nil {
		return false
	}
	if IsEmpty(pLinkTable){
		pLinkTable.pHead=pNode
		pLinkTable.pTail=pNode
		pLinkTable.SumOfNode++
	}else{
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
		pLinkTable.SumOfNode++
	}


 }
 
 func DelLinkTableNode(pLinkTable *LinkTable , pNode *LinkTableNode) int{
	 if IsEmpty(pLinkTable) return -1
	 p:= pLinkTable.pHead
	 if p.data==pNode.data{
		pLinkTable.pHead = pLinkTable.pHead.pNext
		if pLinkTable.pHead == nil {
			pLinkTable.pTail =pLinkTable.pHead
		}
	 }
	 
	 for ;p!=nil;p=p.pNext{
		if p.data == pNode.data{
			 
			if pLinkTable.pTail==p{
				pLinkTable.pTail=pre
			}
			pre.pNext=p.pNext
			pLinkTable.SumOfNode--
			return 1
		
		}
		pre:=p
	}
	return 0

 }

 func  GetLinkTableHead(pLinkTable *LinkTable)* LinkTableNode{
	if pLinkTable == nil {
		return nil
	}
	 return pLinkTable.pHead
 }

 func GetNextLinkTableNode(pLinkTable *LinkTable, pNode *LinkTableNode) *LinkTableNode{
	if pLinkTable == nil || pNode == nil {
		return nil
	}
	
	for p:= pLinkTable.pHead;p!=nil;p=p.pNext{
		if p == pNode{
			return p.pNext
		}
	}

 }
 

 

