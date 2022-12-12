package main

type Stack []string

func (s *Stack) push(str string){
    *s = append(*s, str)
}

func (s *Stack) pop() (string, bool) {
    var top int = len(*s) - 1

    if top == -1 {
        return "", false
    }

    var element string = (*s)[top]
    *s = (*s)[:top]
    return element, true
}

func (s *Stack) peek() (string, bool) {
    var top int = len(*s) - 1

    if top == -1 {
        return "", false
    }

    return (*s)[top], true
}
