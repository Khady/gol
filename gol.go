package gol

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func get_rules(rules string) ([]int, []int) {
	r := strings.Split(rules, "/")
	born := make([]int, len(r[0][1:]))
	alive := make([]int, len(r[1][1:]))
	for i, e := range r[0][1:] {
		born[i] = int(e - 48)
	}
	for i, e := range r[1][1:] {
		alive[i] = int(e - 48)
	}
	return born, alive
}

func side_cells(C, L int) [][8]int {
	sides := make([][8]int, C*L)
	for i := 0; i < L; i++ {
		for j := 0; j < C; j++ {
			sides[i*C+j][0] = ((((i-2)%L + 1 + L) * C) + (((j-2+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][1] = (((i-2)%L+1+L)*C + j) % (C * L)
			sides[i*C+j][2] = ((((i-2)%L + 1 + L) * C) + (((j+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][3] = (i*C + ((j-2+C)%C+1)%C) % (C * L)
			sides[i*C+j][4] = (i*C + ((j+C)%C+1)%C) % (C * L)
			sides[i*C+j][5] = (((i%L + 1 + L) * C) + (((j-2+C)%C + 1) % C)) % (C * L)
			sides[i*C+j][6] = ((i%L+1+L)*C + j) % (C * L)
			sides[i*C+j][7] = (((i%L + 1 + L) * C) + (((j+C)%C + 1) % C)) % (C * L)
		}
	}
	return sides
}

func count_alive(world []byte, sides [8]int) int {
	alive := 0
	for i := 0; i < 8; i++ {
		if world[sides[i]] == 1 {
			alive++
		}
	}
	return alive
}

func is_in(value int, data []int) bool {
	in := false
	for i := range data {
		if data[i] == value {
			in = true
			break
		}
	}
	return in
}

func cycle(world []byte, new_world []byte, sides [][8]int, born []int, stay_alive []int) []byte {
	for pos := range world {
		alive := count_alive(world, sides[pos])
		if world[pos] == 1 {
			if is_in(alive, stay_alive) {
				new_world[pos] = 1
			} else {
				new_world[pos] = 0
			}
		} else {
			if is_in(alive, born) {
				new_world[pos] = 1
			} else {
				new_world[pos] = 0
			}
		}
	}
	return new_world
}

func dump(world []byte, C, L int) {
	fmt.Println("--------------------------------------------------------------------------------")
	for i := 0; i < (C * L); i += C {
		fmt.Println(world[i : i+C])
	}
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func import_map(m string, c, l int, world []byte) []byte {
	if len(m) != 0 && len(m) <= ((c+1)*l) + 1 {
		for i, e := range m {
			world[i] = byte(e) - 48
		}
	}
	return world
}

func import_file(f string, world []byte) []byte {
	if len(f) != 0 {
		b, _ := ioutil.ReadFile(f)
		pos := 0
		for _, c := range string(b) {
			if c == '1' || c == '0' {
				world[pos] = byte(c) - 48
				pos++
			}
		}
	}
	return world
}

func Init(rules, m, f string, c, l, cycle int) (born []int, alive []int, world []byte, new_world []byte, sides [][8]int) {
	born, alive = get_rules(rules)
	world = make([]byte, c*l)
	new_world = make([]byte, c*l)
	sides = side_cells(c, l)
	world = import_map(m, c, l, world)
	world = import_file(f, world)
	return born, alive, world, new_world, sides
}

func Launch(columns, lines, cycles int, born []int, alive []int, world []byte, new_world []byte, sides [][8]int) {
	dump(world, columns, lines)
	for i := 0; i < cycles; i++ {
		new_world = cycle(world, new_world, sides, born, alive)
		copy(world, new_world)
		dump(world, columns, lines)
	}
}
