package Words

import (
	internals2 "xueyigou_demo/tools/illegal_word/internals"
)

type WordsSearch struct {
	_first    []*internals2.TrieNode2
	_keywords []string
}

func NewWordsSearch() *WordsSearch {
	return &WordsSearch{
		_first: make([]*internals2.TrieNode2, 0),
	}
}

func (this *WordsSearch) SetKeywords(keywords []string) {
	this._keywords = keywords

	root := internals2.NewTrieNode()
	allNodeLayers := make(map[int]*internals2.TrieNodes, 0)

	for r, p := range keywords {
		length := len(p)
		if length > 0 {
			var nd *internals2.TrieNode
			nd = root
			for i, ch := range p {
				nd = nd.Add(int(ch))
				if nd.Layer == 0 {
					nd.Layer = i + 1
					if trieNodes, ok := allNodeLayers[nd.Layer]; ok {
						trieNodes.Items = append(trieNodes.Items, nd)
					} else {
						trieNodes = internals2.NewTrieNodes()
						allNodeLayers[nd.Layer] = trieNodes
						trieNodes.Items = append(trieNodes.Items, nd)
					}
				}
			}
			nd.SetResults(r)
		}
	}

	allNode := make([]*internals2.TrieNode, 1)
	allNode[0] = root
	for i := 0; i < len(allNodeLayers); i++ {
		nodes := allNodeLayers[i+1].Items
		for j := 0; j < len(nodes); j++ {
			allNode = append(allNode, nodes[j])
		}
	}

	for i := 1; i < len(allNode); i++ {
		var nd *internals2.TrieNode
		var r *internals2.TrieNode

		nd = allNode[i]
		nd.Index = i
		r = nd.Parent.Failure
		c := nd.Char
		for {
			if r != nil {
				if _, ok := r.M_values[c]; ok {
					break
				} else {
					r = r.Failure
				}
			} else {
				break
			}
		}

		if r == nil {
			nd.Failure = root
		} else {
			nd.Failure = r.M_values[c]
			for j := 0; j < len(nd.Failure.Results); j++ {
				nd.SetResults(nd.Failure.Results[j])
			}
		}
	}
	root.Failure = root

	allNode2 := make([]*internals2.TrieNode2, len(allNode))
	for i := 0; i < len(allNode); i++ {
		allNode2[i] = internals2.NewTrieNode2()
	}
	for i := 0; i < len(allNode); i++ {
		oldNode := allNode[i]
		newNode := allNode2[i]

		for key, val := range oldNode.M_values {
			var index = val.Index
			newNode.Add(int32(key), allNode2[index])
		}
		for j := 0; j < len(oldNode.Results); j++ {
			newNode.SetResults(oldNode.Results[j])
		}

		oldNode = oldNode.Failure
		for oldNode != root {
			for key, val := range oldNode.M_values {
				if newNode.HasKey(int32(key)) == false {
					var index = val.Index
					newNode.Add(int32(key), allNode2[index])
				}
			}
			for j := 0; j < len(oldNode.Results); j++ {
				newNode.SetResults(oldNode.Failure.Results[j])
			}
			oldNode = oldNode.Failure
		}
	}

	first := make([]*internals2.TrieNode2, 0x10000)
	for key, val := range allNode2[0].M_values {
		first[key] = val
	}
	this._first = first
}

func (this *WordsSearch) FindFirst(text string) *WordsSearchResult {
	var ptr *internals2.TrieNode2
	var i int
	for _, t := range text {
		var tn *internals2.TrieNode2
		if ptr == nil {
			tn = this._first[t]
		} else {
			var b bool
			b, tn = ptr.TryGetValue(t)
			if b == false {
				tn = this._first[t]
			}
		}
		if tn != nil {
			if tn.End == true {
				k := this._keywords[tn.Results[0]]
				maxLength := len(k)
				return NewWordsSearchResult(k, i+1-maxLength, i, tn.Results[0])
			}
		}
		ptr = tn
		i++
	}
	return nil
}

func (this *WordsSearch) FindAll(text string) []*WordsSearchResult {
	list := make([]*WordsSearchResult, 0)
	var ptr *internals2.TrieNode2
	var i int
	for _, t := range text {
		var tn *internals2.TrieNode2
		if ptr == nil {
			tn = this._first[t]
		} else {
			var b bool
			b, tn = ptr.TryGetValue(t)
			if b == false {
				tn = this._first[t]
			}
		}
		if tn != nil {
			if tn.End == true {
				for i := 0; i < len(tn.Results); i++ {
					k := this._keywords[tn.Results[i]]
					maxLength := len(k)
					r := NewWordsSearchResult(k, i+1-maxLength, i, tn.Results[i])
					list = append(list, r)
				}
			}
		}
		ptr = tn
		i++
	}
	return list
}

func (this *WordsSearch) ContainsAny(text string) bool {
	var ptr *internals2.TrieNode2
	for _, t := range text {
		var tn *internals2.TrieNode2
		if ptr == nil {
			tn = this._first[t]
		} else {
			var b bool
			b, tn = ptr.TryGetValue(t)
			if b == false {
				tn = this._first[t]
			}
		}
		if tn != nil {
			if tn.End == true {
				return true
			}
		}
		ptr = tn
	}
	return false
}

func (this *WordsSearch) Replace(text string, replaceChar int32) string {
	result := []rune(text)
	var ptr *internals2.TrieNode2
	var i int
	for _, t := range text {
		var tn *internals2.TrieNode2
		if ptr == nil {
			tn = this._first[t]
		} else {
			var b bool
			b, tn = ptr.TryGetValue(t)
			if b == false {
				tn = this._first[t]
			}
		}
		if tn != nil {
			if tn.End == true {
				maxLength := len([]rune(this._keywords[tn.Results[0]]))
				start := i + 1 - maxLength
				for j := start; j <= i; j++ {
					result[j] = replaceChar
				}
			}
		}
		ptr = tn
		i++
	}
	return string(result)
}
