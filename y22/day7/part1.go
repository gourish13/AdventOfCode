package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type File struct {
    name string
    size int
}

type Dir struct {
    name string
    size int
    parent *Dir
    subdirs map[string]*Dir
    files []File
}

var buffer string
var pwd *Dir
var dirs []*Dir

func newdir(name string, parent *Dir) *Dir {
    return &Dir{
        name: name,
        parent: parent,
        subdirs: make(map[string]*Dir),
    }
}


func readline(scanner *bufio.Scanner) (string, bool) {
    if buffer != "" {
        str := buffer
        buffer = ""
        return str, true
    }
    if scanner.Scan() {
        return scanner.Text(), true
    }
    return "", false
}

func unreadline(line string) {
    buffer = line
}

// ----------------------------------------------------------------------------

func updateDirTree(scanner *bufio.Scanner) bool {
    for {
        line, ok := readline(scanner)
        if !ok { // EOF while reading dir/file in fs
            return false
        }

        file := strings.Fields(line)
        if file[0] == "$" { // File/Dir list finished, next cmd read
            unreadline(line)
            return true
        }

        if file[0] == "dir" {
            pwd.subdirs[file[1]] = newdir(file[1], pwd)
            dirs = append(dirs, pwd.subdirs[file[1]])

        } else {
            newfilesize, _ := strconv.Atoi(file[0])
            pwd.files = append(pwd.files, File{name: file[1], size: newfilesize})
        }
    }
}

func processCommand(scanner *bufio.Scanner) {
    for {
        line, ok := readline(scanner)
        if !ok { // EOF while reading cmd
            return
        }

        cmd := strings.Fields(line)
        if cmd[0] == "$" && cmd[1] == "cd" {
            if cmd[2] == "/" {
                pwd = dirs[0]
            } else if cmd[2] == ".." {
                pwd = pwd.parent
            } else {
                pwd = pwd.subdirs[cmd[2]]
            }
        } else if cmd[0] == "$" && cmd[1] == "ls" {
            if !updateDirTree(scanner) { // EOF while updating fs tree
                return
            }
        }
    }
}

func calculateSizes(dir *Dir) {
    for _, subdir := range dir.subdirs {
        calculateSizes(subdir)
        dir.size += subdir.size
    }
    for _, file := range dir.files {
        dir.size += file.size
    }
}

// ----------------------------------------------------------------------------

func dirsizeSum(atmostSize int) int {
    var sum int
    for _, dir := range dirs {
        if (dir.size <= atmostSize) {
            sum += dir.size
        }
    }
    return sum
}

func main() {
    var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

    buffer = ""
    dirs = append(dirs, newdir("/", nil))

    processCommand(scanner)
    calculateSizes(dirs[0])
    var sumSize = dirsizeSum(100000)

    fmt.Println(sumSize)
}
